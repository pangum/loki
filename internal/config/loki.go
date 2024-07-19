package config

type Loki struct {
	// 推送地址
	Url string `json:"url" yaml:"url" xml:"url" toml:"url" validate:"required,url"`
	// 标签
	Labels map[string]string `json:"labels" yaml:"labels" xml:"labels" toml:"labels"`
	// 批量操作
	Batch *Batch `json:"batch" yaml:"batch" xml:"batch" toml:"batch"`
	// 用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username" validate:"required_with=Password"`
	// 密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password" validate:"required_with=Username"`
	// 租户
	Tenant string `json:"tenant,omitempty" yaml:"tenant" xml:"tenant" toml:"tenant"`
}
