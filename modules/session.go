package modules

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	E "github.com/yuw-mvc/yuw/exceptions"
	"time"
)

func Sd(ctx *gin.Context) *sd {
	return newSd(ctx)
}

func SdInstance() gin.HandlerFunc {
	E.ErrArray(&E.ErrType{"yuw^ad_a": I == nil})
	return sessions.Sessions(sessionPoT.Sid, newSession().instance().rStore)
}

func newSd(ctx *gin.Context) *sd {
	return &sd{d: sessions.Default(ctx)}
}

func (s *sd) Default() sessions.Session {
	return s.d
}

func (s *sd) Is(key interface{}) (ok bool, val interface{}) {
	val = s.d.Get(key)

	if val != nil {
		ok = true
		s.d.AddFlash(key)
		s.d.Save()
		return
	}

	return
}

type (
	sd struct {
		d sessions.Session
	}

	session struct {
		rStore redis.Store
	}

	sPoT struct {
		Sid string
		Prefix string
		Pool int
		Network string
		Addr string
		Password string
		KeyPairs string
		StoreMaxAge time.Duration
		StorePath string
	}
)

var sessionPoT *sPoT = &sPoT{}

func newSession() (s *session) {
	E.ErrArray(&E.ErrType{"yuw^m_admin_a":sessionPoT == nil})
	return &session{}
}

func (session *session) instance() *session {
	store, err := redis.NewStore(
		sessionPoT.Pool, sessionPoT.Network, sessionPoT.Addr, sessionPoT.Password,
		[]byte(sessionPoT.KeyPairs),
	)

	E.ErrPanic(err)

	store.Options(sessions.Options{
		Path:     sessionPoT.StorePath,
		Domain:   "",
		MaxAge:   int(sessionPoT.StoreMaxAge * time.Minute),
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})

	if sessionPoT.Prefix != "" {
		redis.SetKeyPrefix(store, sessionPoT.Prefix)
	}

	session.rStore = store
	return session
}




