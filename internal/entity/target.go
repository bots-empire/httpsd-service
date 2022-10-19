package entity

import (
	"github.com/jackc/pgtype"
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

type TargetDbDTO struct {
	IpAddress pgtype.Name  `json:"ip_address,omitempty"`
	Labels    pgtype.JSONB `json:"labels,omitempty"`
}
