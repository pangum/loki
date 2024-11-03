package loki

import (
	"github.com/pangum/loki/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Put(
		plugin.LoadConfig,
	).Build().Put(
		plugin.New,
	).Name("logger.pangum.loki").Build().Build().Apply()
}
