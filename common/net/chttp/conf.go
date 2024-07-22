package chttp
import "time"
type config struct{
    Network string `yaml:network`
    Address string `yaml:address`
    ReadTimeout time.Duration `yaml:readTimeout`
    WriteTimeout time.Duration `yaml:writeTimeout`
}
