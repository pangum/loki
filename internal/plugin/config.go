package plugin

import (
	"github.com/pangum/loki/internal/config"
	"github.com/pangum/pangu"
)

func LoadConfig(config *pangu.Config) (logging *config.Logging, err error) {
	wrapper := new(Wrapper)
	if le := config.Build().Get(wrapper); nil != le {
		err = le
	} else {
		logging = wrapper.Logging
	}

	return
}
