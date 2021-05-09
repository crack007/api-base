package config

import "cracker.com/base/common/constant"

type CommonConfig struct {
	env         string `desc:"运行环境"`
	httpTimeout int    `desc:"http请求超时时间(单位：秒)"`
}

func (a *CommonConfig) Env() string {
	return a.env
}

func (a *CommonConfig) SetEnv(env string) {
	a.env = env
}

func (a *CommonConfig) HttpTimeout() int {
	return a.httpTimeout
}

func (a *CommonConfig) SetHttpTimeout(httpTimeout int) {
	a.httpTimeout = httpTimeout
}

var commonConfig = &CommonConfig{}

func GetCommonConfig() *CommonConfig {
	return commonConfig
}
func init() {
}

func (a *CommonConfig) IsProd() bool {
	return a.Env() == constant.ENV_PROD
}

func (a *CommonConfig) IsTest() bool {
	return a.Env() == constant.ENV_TEST
}

func (a CommonConfig) IsDev() bool {
	return a.Env() == constant.ENV_DEV
}
