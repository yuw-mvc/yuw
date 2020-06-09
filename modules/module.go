package modules

import (
	"github.com/spf13/cast"
	"github.com/spf13/pflag"
	C "github.com/yuw-mvc/yuw/configs"
	E "github.com/yuw-mvc/yuw/exceptions"
	"time"
)

type (
	YuwInitialize struct {
		Table bool
		Redis bool
		I18nT bool
		Email bool
	}

	module struct {
		util *Utils
	}
)

func init() {
	m := New()
	m.cfg()

	var YuwInitialized *YuwInitialize
	E.ErrPanic(m.util.MapToStruct(
		I.Get("YuwInitialize", []interface{}{}),
		&YuwInitialized,
	))

	if YuwInitialized.Table {
		strTimeLocation := cast.ToString(I.Get("Yuw.TimeLocation", C.LocationAsiaShanghai))
		TimeLocation, err := m.util.SetTimeLocation(strTimeLocation)
		E.ErrPanic(err)

		m.db(TimeLocation)
	}

	if YuwInitialized.Redis {
		m.rd()
	}
}

func New() *module {
	return &module {
		util: NewUtils(),
	}
}

func (module *module) cfg() {
	if I != nil {
		return
	}

	pflag.String("env", "", "environment configure")
	pflag.Parse()

	init := NewInitialize()
	E.ErrPanic(init.Env.BindPFlags(pflag.CommandLine))

	I = init.LoadInitializedFromYaml()
	E.ErrArray(&E.ErrType{"yuw^m_a":I == nil})
}

func (module *module) db(timeLocation *time.Location) {
	sysTimeLocation = timeLocation

	var configs *dbConfigs

	cfg := I.Get("DBClusters.Configure", map[string]interface{}{}).(map[string]interface{})
	if len(cfg) == 0 {
		configs = &dbConfigs {
			DriverName: "mysql",
			MaxOpen: 1000,
			MaxIdle: 500,
			ShowedSQL: false,
			CachedSQL: false,
		}
	} else {
		E.ErrPanic(module.util.MapToStruct(cfg, &configs))
	}

	env := I.Get("DBClusters.Databases", map[string]interface{}{}).(map[string]interface{})
	E.ErrArray(&E.ErrType{"yuw^m_c":len(env) == 0})

	masterDB = &dbs{}
	slaverDB = &dbs{}

	for table, db := range env {
		for method, databases := range db.(map[string]interface{}) {
			var dbEngines []*database = make([]*database, len(databases.([]interface{})))
			for key, database := range databases.([]interface{}) {
				var cluster *dbCluster

				toMap := module.util.InterfaceToStringInMap(database.(map[interface{}]interface{}))
				toMap["Table"] = table

				E.ErrPanic(module.util.MapToStruct(toMap, &cluster))

				dbEngine := NewDatabase()
				dbEngine.dbCluster = cluster
				dbEngine.dbConfigs = configs

				dbEngines[key] = dbEngine.instance()
			}

			switch method {
			case "master":
				(*masterDB)[table] = dbEngines
			case "slaver":
				(*slaverDB)[table] = dbEngines
			default:
				continue
			}
		}
	}
}

func (module *module) rd() {
	rdNetwork := I.Get("Redis.Network", "").(string)
	if ok, _ := module.util.StrContains(rdNetwork, rdNetworkType ...); ok == false {
		rdNetwork = defaultRdNetworkType
	}

	rdReader := I.Get("Redis.Reader", []interface{}{}).([]interface{})
	rdWriter := I.Get("Redis.Writer", []interface{}{}).([]interface{})
	E.ErrArray(&E.ErrType{
		"yuw^m_init_e":len(rdReader) == 0 || len(rdWriter) == 0,
	})

	rdsRW := map[string][]interface{} {
		"reader": rdReader,
		"writer": rdWriter,
	}

	rdsReader = &rds{}
	rdsWriter = &rds{}

	for k, rdsVal := range rdsRW {
		for _, rdVal := range rdsVal {
			var res *RedisPoT

			err := module.util.InterfaceToStruct(rdVal.(map[interface{}]interface{}), &res)
			if err != nil {
				continue
			}

			r := NewRedis()
			r.rPoT = res

			var dbKey string = prefixRdDB + cast.ToString(r.rPoT.DB)

			switch k {
			case "reader":
				(*rdsReader)[dbKey] = append((*rdsReader)[dbKey], r.instance())
				continue

			case "writer":
				(*rdsWriter)[dbKey] = append((*rdsWriter)[dbKey], r.instance())
				continue

			default:
				continue
			}
		}
	}
}
