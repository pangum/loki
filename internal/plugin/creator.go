package plugin

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/loki"
	"github.com/pangum/loki/internal/core"
	"github.com/pangum/pangu"
)

type Creator struct {
	// 纯方法封装
}

func (c *Creator) New(config *pangu.Config, http *http.Client) (logger log.Logger, err error) {
	wrapper := new(Wrapper)
	if le := config.Build().Get(wrapper); nil != le {
		err = le
	} else {
		logger, err = c.new(&wrapper.Logging, http)
	}

	return
}

func (c *Creator) new(config *Config, http *http.Client) (logger log.Logger, err error) {
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
		logger, err = builder.Factory(factory.Build()).Build()
	default:
		err = exc.NewField("不支持的日志收集器", field.New("type", config.Type))
	}

	return
}
