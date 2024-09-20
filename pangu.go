package loki

import (
	"github.com/pangum/loki/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Puts(
		ctor.LoadConfig,
		ctor.New,
	).Build().Apply()
}
