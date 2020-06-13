package subscribe

import (
	"github.com/spf13/cast"
	M "github.com/yuw-mvc/yuw/modules"
)

func Do(subscribePoT *PoT) {
	boolRd := M.I.Get("YuwInitialize.Redis", false).(bool)
	boolSubscribed := M.I.Get("YuwConsole.Subscribed", false).(bool)

	if boolRd && boolSubscribed {
		keys = subscribePoT.Keys
		channels = subscribePoT.Channels

		go new(subscribe).do()
	}
}

type (
	PoTKeys []interface{}
	PoTChannels map[string][]Provider

	Subscribe interface {
		Subscribe(channels ...string)
	}

	PoT struct {
		Keys *PoTKeys
		Channels *PoTChannels
	}

	subscribe struct {

	}
)

var (
	_ Subscribe = new(rd)

	keys *PoTKeys
	channels *PoTChannels
)

func (subscribed *subscribe) do() {
	new(rd).Subscribe(cast.ToStringSlice(keys) ...)
}
