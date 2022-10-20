package entity

import (
	"github.com/prometheus/prometheus/model/labels"
)

type TargetPrometheus struct {
	Targets []string      `json:"targets"`
	Labels  labels.Labels `json:"labels"`
}

type TargetDb struct {
	IpAddress string            `json:"ip_address"`
	Labels    map[string]string `json:"labels"`
}
