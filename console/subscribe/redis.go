package subscribe

import (
	E "github.com/yuw-mvc/yuw/exceptions"
	M "github.com/yuw-mvc/yuw/modules"
)

type rd struct {}

func (subscribed *rd) Subscribe(channels ...string) {
	r, errRd := M.InstanceRd(&M.RdPoT{Selector:true})
	E.ErrPanic(errRd)

	rSubscribe := r.Engine().R.Subscribe(channels ...)
	_, err := rSubscribe.Receive()
	if err != nil {
		return
	}

	provider := NewProvider()
	for msg := range rSubscribe.Channel() {
		provider.Do(msg)
	}
}