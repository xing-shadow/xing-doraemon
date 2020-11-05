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
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	clienApi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/rules"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"
	"xing-doraemon/interval/service/AlertService"
)

type Manager struct {
	Config  Config
	Prom    Prom
	Options *rules.ManagerOptions
	Manager *rules.Manager
	Rules   Rules

	logger log.Logger
}

func NewManager(ctx context.Context, logger log.Logger, prom Prom, cfg Config) (*Manager, error) {
	localStorage, err := NewMockStorage()
	if err != nil {
		return nil, err
	}
	options := &rules.ManagerOptions{
		QueryFunc: HTTPQueryFunc(
			log.With(logger, "component", "http query func"),
			prom.URL,
		),
		NotifyFunc: HTTPNotifyFunc(
			log.With(logger, "component", "http notify func"),
			cfg.NotifyRetries,
		),
		Context:         ctx,
		Appendable:      localStorage,
		TSDB:            localStorage,
		ExternalURL:     &url.URL{},
		Logger:          log.With(logger, "component", "rule manager"),
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
	level.Info(m.logger).Log("msg", "start rule manager", "prom_id", m.Prom.ID)
	m.Manager.Run()
}

// Stop ...
func (m *Manager) Stop() {
	level.Info(m.logger).Log("msg", "stop rule manager", "prom_id", m.Prom.ID)
	m.Manager.Stop()
}

func (m *Manager) Update(rules Rules) error {
	m.Rules = rules
	filePath := filepath.Join(os.TempDir(), fmt.Sprintf("rule.%d.yml", m.Prom.ID))
	content, err := rules.Content()
	if err != nil {
		level.Error(m.logger).Log("msg", "get prom rule error", "error", err, "prom_id", m.Prom.ID)
		return err
	}
	level.Info(m.logger).Log("path", filePath)
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		level.Error(m.logger).Log("msg", "write file error", "error", err, "prom_id", m.Prom.ID)
		return err
	}
	return m.Manager.Update(time.Duration(m.Config.EvaluationInterval), []string{filePath})
}

// HTTPNotifyFunc
func HTTPNotifyFunc(logger log.Logger, retries int) rules.NotifyFunc {
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
			level.Error(logger).Log("msg", "encode json error", "error", err, "alerts", alerts)
			return
		}
		for i := 0; i < retries; i++ {
			err := AlertService.GetAlertHandle().HandleAlert(data)
			if err != nil {
				level.Error(logger).Log("msg", "notify error", "error", err, "retries", i)
				time.Sleep(time.Second)
			} else {
				level.Debug(logger).Log("msg", "notify success")
				return
			}
		}
	}
}

// HTTPQueryFunc
// TODO: use http keep-alive
func HTTPQueryFunc(logger log.Logger, url string) rules.QueryFunc {
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
			level.Debug(logger).Log(
				"msg", "query vector seccess",
				"query", q,
				"vector", vector,
			)
			return vector, nil
		default:
			// TODO: other type: "matrix" | "vector" | "scalar" | "string",
			return vector, fmt.Errorf("unknown result type [%s] query=[%s]", value.Type().String(), q)
		}
	}
}
