package plugin

import (
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/loki"
	"github.com/pangum/loki/internal/config"
	"github.com/pangum/loki/internal/core"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *config.Logging, http *http.Client) (logger log.Logger, err error) {
	builder := log.New().Level(log.ParseLevel(config.Level))
	if nil != config.Stacktrace {
		builder.Stacktrace(*config.Stacktrace)
	}

	switch config.Type {
	case core.TypeLoki:
		self := config.Loki
		factory := loki.New().Url(self.Url).Http(http)
		if "" != self.Username || "" != self.Password {
			factory.Username(self.Username)
			factory.Password(self.Password)
		}
		if nil != self.Batch {
			factory.Batch(self.Batch.Size, self.Batch.Wait)
		}
		if 0 != len(self.Labels) {
			factory.Labels(self.Labels)
		}
		if "" != self.Tenant {
			factory.Tenant(self.Tenant)
		}
		logger, err = builder.Factory(factory.Build()).Build()
	default:
		err = exception.New().Message("不支持的日志收集器").Field(field.New("type", config.Type)).Build()
	}

	return
}

func (c *Constructor) LoadConfig(config *pangu.Config) (logging *config.Logging, err error) {
	wrapper := new(Wrapper)
	if le := config.Build().Get(wrapper); nil != le {
		err = le
	} else {
		logging = wrapper.Logging
	}

	return
}
