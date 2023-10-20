package plugin

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/simaqian"
	"github.com/pangum/loki/internal/core"
	"github.com/pangum/pangu"
)

type Creator struct {
	// 纯方法封装
}

func (c *Creator) New(config *pangu.Config, http *http.Client) (logger simaqian.Logger, err error) {
	wrapper := new(Wrapper)
	if le := config.Build().Get(wrapper); nil != le {
		err = le
	} else {
		logger, err = c.new(&wrapper.Logging, http)
	}

	return
}

func (c *Creator) new(config *Config, http *http.Client) (logger simaqian.Logger, err error) {
	builder := simaqian.New().Level(simaqian.ParseLevel(config.Level))
	if nil != config.Stacktrace {
		builder.Stacktrace(*config.Stacktrace)
	}

	switch config.Type {
	case core.TypeLoki:
		self := config.Loki
		loki := builder.Loki().Url(self.Url).Http(http)
		if "" != self.Username || "" != self.Password {
			loki.Username(self.Username)
			loki.Password(self.Password)
		}
		if nil != self.Batch {
			loki.Batch(self.Batch.Size, self.Batch.Wait)
		}
		if 0 != len(self.Labels) {
			loki.Labels(self.Labels)
		}
		logger, err = loki.Build()
	default:
		err = exc.NewField("不支持的日志收集器", field.New("type", config.Type))
	}

	return
}
