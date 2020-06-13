package subscribe

import (
	"github.com/go-redis/redis"
	M "github.com/yuw-mvc/yuw/modules"
)

type (
	Provider interface {
		Provided(channel string, content interface{})
	}

	provider struct {
		util *M.Utils
	}
)

func NewProvider() *provider {
	return &provider {
		util: M.NewUtils(),
	}
}

func (srv *provider) Do(msg *redis.Message) {
	if ok, _ := srv.util.StrContains(msg.Channel, *keys ...); ok == false {
		return
	}

	x, ok := (*channels)[msg.Channel]
	if ok {
		for _, v := range x {
			v.Provided(msg.Channel, msg.Payload)
		}
	}
}



