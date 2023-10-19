package loki

import (
	"github.com/pangum/loki/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependencies().Build().Provide(new(plugin.Creator).New)
}
