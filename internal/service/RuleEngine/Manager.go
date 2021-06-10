/*
 * @Time : 2020/10/23 14:49
 * @Author : wangyl
 * @File : Manager.go
 * @Software: GoLand
 */
package RuleEngine

import (
	"context"
	"encoding/json"
	"fmt"
	clienApi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/rules"
	"go.uber.org/zap"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"
	"xing-doraemon/internal/service/AlertService"
)

type Manager struct {
	Config  Config
	Prom    Prom
	Options *rules.ManagerOptions
	Manager *rules.Manager
	Rules   Rules

	logger *zap.Logger
}

func NewManager(ctx context.Context, logger *zap.Logger, prom Prom, cfg Config) (*Manager, error) {
	localStorage, err := NewMockStorage()
	if err != nil {
		return nil, err
	}
	options := &rules.ManagerOptions{
		QueryFunc: HTTPQueryFunc(
			logger,
			prom.URL,
		),
		NotifyFunc: HTTPNotifyFunc(
			logger,
			cfg.NotifyRetries,
		),
		Context:         ctx,
		Appendable:      localStorage,
		TSDB:            localStorage,
		ExternalURL:     &url.URL{},
		OutageTolerance: time.Hour,
		ForGracePeriod:  10 * time.Minute,
		ResendDelay:     time.Minute,
	}

	manager := rules.NewManager(options)

	return &Manager{
		Config:  cfg,
		Prom:    prom,
		Options: options,
		Manager: manager,
		logger:  logger,
	}, nil
}

// Run ...
func (m *Manager) Run() {
	m.logger.Info("start rule manager", zap.Int64("prom_id", m.Prom.ID))
	m.Manager.Run()
}

// Stop ...
func (m *Manager) Stop() {
	m.logger.Warn("stop rule manager", zap.Int64("prom_id", m.Prom.ID))
	m.Manager.Stop()
}

func (m *Manager) Update(rules Rules) error {
	m.Rules = rules
	filePath := filepath.Join(os.TempDir(), fmt.Sprintf("rule.%d.yml", m.Prom.ID))
	content, err := rules.Content()
	if err != nil {
		m.logger.Error("get prom rule fail:"+err.Error(), zap.Int64("prom_id", m.Prom.ID))
		return err
	}
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		m.logger.Error("write file fail:"+err.Error(), zap.Int64("prom_id", m.Prom.ID))
		return err
	}
	return m.Manager.Update(time.Duration(m.Config.EvaluationInterval), []string{filePath}, nil)
}

// HTTPNotifyFunc
func HTTPNotifyFunc(logger *zap.Logger, retries int) rules.NotifyFunc {
	return func(ctx context.Context, expr string, alerts ...*rules.Alert) {
		if len(alerts) == 0 {
			return
		}
		var newAlerts []*Alert
		for _, alert := range alerts {
			newAlerts = append(newAlerts, (*Alert)(alert))
		}
		data, err := json.Marshal(newAlerts)
		if err != nil {
			logger.Error("encode json fail:" + err.Error())
			return
		}
		for i := 0; i < retries; i++ {
			err := AlertService.PushNotify(data)
			if err != nil {
				logger.Error("notify error"+err.Error(), zap.Int("retries", i))
				time.Sleep(time.Second)
			} else {
				logger.Info("notify success")
				return
			}
		}
	}
}

// HTTPQueryFunc
// TODO: use http keep-alive
func HTTPQueryFunc(logger *zap.Logger, url string) rules.QueryFunc {
	c, _ := clienApi.NewClient(clienApi.Config{
		Address: url,
	})
	api := v1.NewAPI(c)
	return func(ctx context.Context, q string, t time.Time) (promql.Vector, error) {
		vector := promql.Vector{}
		value, _, err := api.Query(ctx, q, t)
		if err != nil {
			return vector, err
		}
		switch value.Type() {
		case model.ValVector:
			for _, i := range value.(model.Vector) {
				l := labels.Labels{}
				for k, v := range i.Metric {
					l = append(l, labels.Label{
						Name:  string(k),
						Value: string(v),
					})
				}
				vector = append(vector, promql.Sample{
					Point: promql.Point{
						T: int64(i.Timestamp),
						V: float64(i.Value),
					},
					Metric: l,
				})
			}
			logger.Info("query vector success", zap.String("query", q), zap.Any("vector", vector))
			return vector, nil
		default:
			// TODO: other type: "matrix" | "scalar" | "string",
			return vector, fmt.Errorf("unknown result type [%s] query=[%s]", value.Type().String(), q)
		}
	}
}
