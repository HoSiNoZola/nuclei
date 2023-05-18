package redis

import (
	original_redis "github.com/projectdiscovery/nuclei/v2/pkg/js/libs/redis"

	"github.com/dop251/goja"
	"github.com/projectdiscovery/nuclei/v2/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("nuclei/libredis")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"GetServerInfo":   original_redis.GetServerInfo,
			"IsAuthenticated": original_redis.IsAuthenticated,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
