package config

import (
	"time"
)

type Batch struct {
	// 最大条数
	// 达到数量后会主动触发推送
	Size int `json:"size" yaml:"size" xml:"size" toml:"size"`
	// 最大等待时间
	// 超过后会主动触发推送
	Wait time.Duration `json:"wait" yaml:"wait" xml:"wait" toml:"wait"`
}
