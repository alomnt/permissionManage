package conf

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	Conf = New()
)

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	config, err := toml.LoadFile("./perm-server/sys-common/conf/config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}
