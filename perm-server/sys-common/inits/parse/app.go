package parse

import (
	"strconv"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

var (
	// conf strut
	C iris.Configuration

	// 解析app.yml中的Other项
	O Other
	// app.conf配置项key定义
	Port       string = "Port"
	ignoreURLs string = "IgnoreURLs"
	jwtTimeout string = "JWTTimeout"
	logLevel   string = "LogLevel"
	secret     string = "Secret"
)

type (
	Other struct {
		Port       string
		IgnoreURLs []string
		JWTTimeout int64
		LogLevel   string
		Secret     string
	}
)

func AppOtherParse() {
	golog.Info("@@@ Init app conf")
	c := iris.YAML("./perm-server/sys-common/conf/app.yml")

	//appData, err := conf.Asset("conf/app.yml")
	//if err != nil {
	//  golog.Fatalf("Error. %s", err)
	//}
	//c := iris.DefaultConfiguration()
	golog.Println(c)
	//if err = yaml.Unmarshal(appData, &c); err != nil {
	//  golog.Fatalf("Error. %s", err)
	//}
	//C = c

	// -------------- 解析每个Other配置项 ---------------

	O.Port = strconv.Itoa(c.GetOther()[Port].(int))

	// 解析other的key
	iURLs := c.GetOther()[ignoreURLs].([]interface{})
	for _, v := range iURLs {
		O.IgnoreURLs = append(O.IgnoreURLs, v.(string))
	}

	jTimeout := c.GetOther()[jwtTimeout].(int)
	O.JWTTimeout = int64(jTimeout)

	O.LogLevel = c.GetOther()[logLevel].(string)

	O.Secret = c.GetOther()[secret].(string)

}
