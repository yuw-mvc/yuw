package modules

import (
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	E "github.com/yuw-mvc/yuw/exceptions"
	"sync"
)

const (
	prefixRdDB string = "db_"
	defaultRdNetworkType string = "tcp"
)

func InstanceRd(rPoT *RdPoT) (rdInstance *rd, err error) {
	var (
		intRand int = 0
		dbKey string = prefixRdDB + cast.ToString(rPoT.DB)
		rdsReaderToDB []*rd = (*rdsReader)[dbKey]
		rdsWriterToDB []*rd = (*rdsWriter)[dbKey]
	)

	if rPoT.Selector {
		if len(rdsWriterToDB) == 0 {
			err = E.Err("yuw^m_rd_b", E.ErrPosition())
			return
		}

		intRand = NewUtils().IntRand(0, len(rdsWriterToDB))
		rdInstance = rdsWriterToDB[intRand]
	} else {
		if len(rdsReaderToDB) == 0 {
			err = E.Err("yuw^m_rd_c", E.ErrPosition())
			return
		}

		intRand = NewUtils().IntRand(0, len(rdsReaderToDB))
		rdInstance = rdsReaderToDB[intRand]
	}

	return
}

type (
	RdPoT struct {
		DB int
		Selector bool
	}

	RedisPoT struct {
		Addr string
		Password string
		DB int
	}

	rd struct {
		R *redis.Client
		rNetworkPoT string
		rPoT *RedisPoT
		mx sync.Mutex
	}

	rds map[string][]*rd
)

var (
	rdsWriter *rds
	rdsReader *rds

	rdNetworkType []interface{} = []interface{}{"tcp", "unix"}
)

func NewRedis() *rd {
	return &rd {}
}

func (rd *rd) Engine() *rd {
	E.ErrArray(&E.ErrType{"yuw^m_rd_a":rd.R == nil})
	return rd
}

func (rd *rd) instance() *rd {
	rd.mx.Lock()
	defer rd.mx.Unlock()

	if rd.R != nil {
		return rd
	}

	clientOptions := &redis.Options{}
	clientOptions.Network = rd.rNetworkPoT
	clientOptions.Addr = rd.rPoT.Addr
	clientOptions.Password = rd.rPoT.Password
	clientOptions.DB = rd.rPoT.DB

	client := redis.NewClient(clientOptions)

	_, err := client.Ping().Result()
	E.ErrPanic(err)

	rd.R = client
	return rd
}


