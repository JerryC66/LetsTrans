package config

type Server struct {
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mongo Mongo `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Qiniu Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	System System `mapstructure:"system" json:"system" yaml:"system"`
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"` // 跨域配置
}
