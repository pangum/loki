package config

type Loki struct {
	// 推送地址
	Url string `json:"url,omitempty"validate:"required,url"`
	// 标签
	Labels map[string]string `json:"labels,omitempty"`
	// 批量操作
	Batch *Batch `json:"batch,omitempty"`
	// 用户名
	Username string `json:"username,omitempty" validate:"required_with=Password"`
	// 密码
	Password string `json:"password,omitempty" validate:"required_with=Username"`
	// 租户
	Tenant string `json:"tenant,omitempty"`
}
