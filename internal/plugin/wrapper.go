package plugin

import (
	"github.com/pangum/loki/internal/config"
)

type Wrapper struct {
	// 日志配置
	Logging *config.Logging `json:"logging,omitempty" validate:"required"`
}
