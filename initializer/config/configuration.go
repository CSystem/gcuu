package config

type Configuration struct {
	JWT             JWTConfiguration
	Server          ServerConfiguration
	DB              DatabaseConfiguration `mapstructure:"database"`
	Redis           RedisConfiguration    `mapstructure:"redis"`
	DumpRepResp     bool                  `mapstructure:"dump_request_response"`
	DebugSQL        bool                  `mapstructure:"debug_sql"`
	NewRelicLicense string                `mapstructure:"new_relic_license"`
	VerifySign      string                `mapstructure:"verify_sign"`
	Submail         SubmailConfiguration  `mapstructure:"submail"`
	WXPay           WXPayConfiguration    `mapstructure:"wxpay"`
	AppInfo			AppInfoConfiguration  `mapstructure:"appinfo"`
}

type DatabaseConfiguration struct {
	Adapter  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

type RedisConfiguration struct {
	Host     string
	Port     int
	Password string
}

type ServerConfiguration struct {
	Port int
	Mode string
}

type JWTConfiguration struct {
	Realm string
	Key   string
}

type SubmailConfiguration struct {
	ID  string
	Key string
}

type WXPayConfiguration struct {
	AppID      string `mapstructure:"app_id"`
	MchID      string `mapstructure:"mch_id"`
	APIKey     string `mapstructure:"api_key"`
	CertFile   string `mapstructure:"cert_file"`
	KeyFile    string `mapstructure:"key_file"`
	Pkcs12File string `mapstructure:"pkcs12_file"`
}

type AppInfoConfiguration struct {
	AppId	uint `mapstructure:"app_id"`
}