package RuleEngine

import (
	"encoding/json"
	"github.com/prometheus/prometheus/rules"
	"math"
)

// Alert
type Alert rules.Alert

func (a *Alert) MarshalJSON() ([]byte, error) {
	for idx, i := range a.Labels {
		if i.Name == "alertname" {
			a.Labels = append(a.Labels[:idx], a.Labels[idx+1:]...)
		}
	}
	return json.Marshal(map[string]interface{}{
		"state":        a.State,
		"labels":       a.Labels,
		"annotations":  a.Annotations,
		"value":        math.Round(a.Value*100) / 100,
		"active_at":    a.ActiveAt,
		"fired_at":     a.FiredAt,
		"resolved_at":  a.ResolvedAt,
		"last_sent_at": a.LastSentAt,
		"valid_until":  a.ValidUntil,
	})
}
