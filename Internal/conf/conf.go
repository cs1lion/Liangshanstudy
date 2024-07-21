package conf
import(
    "/common/net/chttp"
)
type Conf struct{
    Server *chttp.Config 'yaml:"server"'
}