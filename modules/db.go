package modules

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	E "github.com/yuw-mvc/yuw/exceptions"
	"strings"
	"sync"
	"time"
)

func InstanceDB(dbTable string, selector ...int) (dbInstance *database, err error) {
	var (
		intRand int = 0
		doTable string = strings.ToLower(dbTable)
		masterLen int = len((*masterDB)[doTable])
		slaverLen int = len((*slaverDB)[doTable])
	)

	if len(selector) == 0 {
		if masterLen == 0 {
			err = E.Err("yuw^m_db_b", E.ErrPosition())
			return
		}

		intRand = NewUtils().IntRand(0, masterLen)
		dbInstance = (*masterDB)[doTable][intRand]
	} else {
		if slaverLen == 0 {
			err = E.Err("yuw^m_db_c", E.ErrPosition())
			return
		}

		intRand = NewUtils().IntRand(0, slaverLen)
		dbInstance = (*slaverDB)[doTable][intRand]
	}

	return
}

type (
	dbCluster struct {
		Host      	string
		Port      	int
		Table     	string
		Username  	string
		Password  	string
	}

	dbConfigs struct {
		DriverName  string
		MaxOpen   	int
		MaxIdle   	int
		ShowedSQL 	bool
		CachedSQL 	bool
	}

	database struct {
		engine *xorm.Engine
		mx sync.Mutex
		dbCluster *dbCluster
		dbConfigs *dbConfigs
	}

	dbs map[string][]*database
)

var (
	sysTimeLocation *time.Location

	masterDB *dbs
	slaverDB *dbs
)

func NewDatabase() *database {
	return &database {}
}

func (db *database) Engine() *xorm.Engine {
	if db.engine == nil {
		E.LogErr("yuw^m_db_a", E.ErrPosition())
		return nil
	}

	return db.engine
}

func (db *database) instance() *database {
	db.mx.Lock()
	defer db.mx.Unlock()

	if db.engine != nil {
		return db
	}

	if db.dbConfigs == nil || db.dbCluster == nil {
		E.ErrPanic(E.Err("yuw^m_db_a", E.ErrPosition()))
	}

	driverFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	driverSource := fmt.Sprintf(
		driverFormat,
		db.dbCluster.Username,
		db.dbCluster.Password,
		db.dbCluster.Host,
		db.dbCluster.Port,
		db.dbCluster.Table,
	)

	engine, err := xorm.NewEngine(db.dbConfigs.DriverName, driverSource)
	if err != nil {
		if engine != nil {
			engine.Close()
		}

		E.ErrPanic(err)
	}

	engine.SetMaxOpenConns(db.dbConfigs.MaxOpen)
	engine.SetMaxIdleConns(db.dbConfigs.MaxIdle)
	engine.SetConnMaxLifetime(time.Second)

	engine.ShowSQL(db.dbConfigs.ShowedSQL)
	engine.SetTZDatabase(sysTimeLocation)

	if db.dbConfigs.CachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	db.engine = engine
	return db
}