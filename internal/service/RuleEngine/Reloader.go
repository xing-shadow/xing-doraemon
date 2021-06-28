package RuleEngine

import (
	"context"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strconv"
	"time"
	"xing-doraemon/global"
	"xing-doraemon/internal/model/db"
)

type Config struct {
	NotifyRetries      int
	EvaluationInterval time.Duration
	ReloadInterval     time.Duration
}

type Reloader struct {
	config   Config
	managers []*Manager
	context  context.Context
	cancel   context.CancelFunc
	running  bool

	logger *zap.Logger
}

func NewReloader(cfg Config) *Reloader {
	ctx, cancel := context.WithCancel(context.Background())
	reloader := &Reloader{
		config:  cfg,
		context: ctx,
		cancel:  cancel,
		running: false,
		logger:  global.Log,
	}

	return reloader
}

func (r *Reloader) Run() {
	r.running = true
	for _, manager := range r.managers {
		manager.Run()
	}
}

// Stop rule manager
func (r *Reloader) Stop() {
	r.running = false
	r.cancel()
	for _, i := range r.managers {
		i.Stop()
	}
}

// Loop for checking the rules
func (r *Reloader) Start() {
	go func() {
		r.Run()
		for r.running {
			r.Upload()
			select {
			case <-r.context.Done():
				r.Stop()
			case <-time.After(time.Duration(r.config.ReloadInterval)):
			}
		}
	}()
}

// download the rules and update rule manager
func (r *Reloader) Upload() error {
	r.logger.Info("start update rule")
	promRules, err := r.getPromRules()
	if err != nil {
		return err
	}

	// stop invalid manager
	for idx, m := range r.managers {
		del := true
		for _, p := range promRules {
			if m.Prom.ID == p.Prom.ID && m.Prom.URL == p.Prom.URL && p.Prom.URL != "" {
				del = false
			}
		}
		if del {
			r.logger.Warn("prom not exist, delete manager", zap.Int64("prom_id", m.Prom.ID), zap.String("prom_url", m.Prom.URL))
			m.Stop()
			r.managers = append(r.managers[:idx], r.managers[idx+1:]...)
		}
	}
	// update rules
	for _, p := range promRules {
		if p.Prom.URL == "" {
			r.logger.Info("prom url is null", zap.Int64("prom_id", p.Prom.ID), zap.String("prom_url", p.Prom.URL))
			continue
		}
		var manager *Manager
		for _, m := range r.managers {
			if m.Prom.ID == p.Prom.ID && m.Prom.URL == p.Prom.URL && p.Prom.URL != "" {
				manager = m
			}
		}
		if manager == nil {
			m, err := NewManager(r.context, r.logger, p.Prom, r.config)
			if err != nil {
				r.logger.Error("create manager fail:"+err.Error(), zap.Int64("prom_id", p.Prom.ID), zap.String("prom_url", p.Prom.URL))
				return err
			}
			m.Run()
			manager = m
			r.managers = append(r.managers, manager)
		}
		if manager != nil {
			err := manager.Update(p.Rules)
			if err != nil {
				r.logger.Error("update rule error", zap.Int64("prom_id", manager.Prom.ID), zap.String("prom_url", manager.Prom.URL))
			} else {
				r.logger.Info("update rule success", zap.Int("len", len(p.Rules)), zap.Int64("prom_id", manager.Prom.ID), zap.String("prom_url", manager.Prom.URL))
			}
		}
	}
	r.logger.Debug("end update rule")
	return nil
}

func (r *Reloader) getPromRules() ([]PromRules, error) {
	var proms []db.Prom
	var result []PromRules
	err := opt.DB.Find(&proms).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		r.logger.Error("get proms fail:" + err.Error())
		return nil, err
	}
	for _, prom := range proms {
		rulesDB := []db.Rule{}
		err := opt.DB.Where("prom_id=?", prom.ID).Find(&rulesDB).Error
		if err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				r.logger.Error("get rules fail"+err.Error(), zap.Uint("prom_id", prom.ID))
				continue
			}
		}
		var rules Rules
		for _, rule := range rulesDB {
			rules = append(rules, Rule{
				ID:          int64(rule.ID),
				PromID:      int64(rule.PromID),
				Expr:        rule.Expr,
				Op:          rule.Op,
				Value:       rule.Value,
				For:         strconv.Itoa(rule.Duration),
				Labels:      nil,
				Summary:     rule.Summary,
				Description: rule.Description,
			})
		}
		result = append(result, PromRules{
			Prom: Prom{
				ID:  int64(prom.ID),
				URL: prom.Url,
			},
			Rules: rules,
		})
	}
	return result, nil
}
