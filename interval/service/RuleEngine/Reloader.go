/*
 * @Time : 2020/10/23 14:49
 * @Author : wangyl
 * @File : Reloader.go
 * @Software: GoLand
 */
package RuleEngine

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/promlog"
	"strconv"
	"time"
	"xing-doraemon/interval/model/db"
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

	logger log.Logger
}

func NewReloader(cfg Config) *Reloader {
	ctx, cancel := context.WithCancel(context.Background())
	level := promlog.AllowedLevel{}
	level.Set("debug")
	formart := promlog.AllowedFormat{}
	formart.Set("json")
	logger := promlog.New(&promlog.Config{
		Level:  &level,
		Format: &formart,
	})
	reloader := &Reloader{
		config:  cfg,
		context: ctx,
		cancel:  cancel,
		running: false,
		logger:  logger,
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
			case <-time.After(time.Duration(r.config.ReloadInterval)):
			}
		}
	}()
}

// download the rules and update rule manager
func (r *Reloader) Upload() error {
	level.Debug(r.logger).Log("msg", "start update rule")
	promrules, err := r.getPromRules()
	if err != nil {
		return err
	}
	// stop invalid manager
	for idx, m := range r.managers {
		del := true
		for _, p := range promrules {
			if m.Prom.ID == p.Prom.ID && m.Prom.URL == p.Prom.URL && p.Prom.URL != "" {
				del = false
			}
		}
		if del {
			level.Info(r.logger).Log("msg", "prom not exist, delete manager", "prom_id", m.Prom.ID, "prom_url", m.Prom.URL)
			m.Stop()
			r.managers = append(r.managers[:idx], r.managers[idx+1:]...)
		}
	}
	// update rules
	for _, p := range promrules {
		if p.Prom.URL == "" {
			level.Error(r.logger).Log("msg", "prom url is null", "prom_id", p.Prom.ID, "prom_url", p.Prom.URL)
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
				level.Error(r.logger).Log("msg", "create manager error", "error", err, "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
				return err
			}
			m.Run()
			manager = m
			r.managers = append(r.managers, manager)
		}
		if manager != nil {
			err := manager.Update(p.Rules)
			if err != nil {
				level.Error(r.logger).Log("msg", "update rule error", "error", err, "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
			} else {
				level.Info(r.logger).Log("msg", "update rule success", "len", len(p.Rules), "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
			}
		}
	}
	level.Debug(r.logger).Log("msg", "end update rule")
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
		level.Error(r.logger).Log("msg", "get proms fail", "error", err)
		return nil, err
	}
	for _, prom := range proms {
		rulesDB := []db.Rule{}
		err := opt.DB.Where("prom_id=?", prom.ID).Find(&rulesDB).Error
		if err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				level.Error(r.logger).Log("msg", "get rules fail", "prom_id", prom.ID, "error", err)
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
				For:         strconv.Itoa(rule.For),
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
