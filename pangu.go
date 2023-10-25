package loki

import (
	"github.com/pangum/loki/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Creator)
	pangu.New().Get().Dependency().Put(
		creator.New,
	).Build().Build().Apply()
}
