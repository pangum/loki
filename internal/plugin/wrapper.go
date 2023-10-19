package plugin

type Wrapper struct {
	// 日志配置
	Logging Config `json:"logging" yaml:"logging" xml:"logging" toml:"logging" validate:"required"`
}
