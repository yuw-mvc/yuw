package modules

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	E "github.com/yuw-mvc/yuw/exceptions"
)

const defaultType = "yaml"

var (
	I *initialize = nil
	defaultEnvironment []interface{} = []interface{}{"dev", "stg", "prd"}
)

type initialize struct {
	fs *File
	Env *viper.Viper
	util *Utils
}

func NewInitialize() *initialize {
	return &initialize {
		Env: viper.New(),
		util: NewUtils(),
	}
}

func (init *initialize) LoadInitializedFromYaml() *initialize {
	init.Env.AddConfigPath(".")
	init.Env.SetConfigType(defaultType)

	str := init.Env.GetString("env")
	E.ErrArray(&E.ErrType{"yuw^m_init_a":str == ""})

	dir := ".env." + str
	env := "./" + dir + "." + defaultType

	okEnv, _ := init.util.StrContains(str, defaultEnvironment ...)
	E.ErrArray(&E.ErrType{"yuw^m_init_b":okEnv == false})

	okDir, _ := init.fs.IsExists(env)
	if okDir == false {
		E.ErrPanic(init.fs.Create(env))
	}

	init.Env.SetConfigName(dir)
	E.ErrPanic(init.Env.ReadInConfig())

	init.Env.WatchConfig()
	init.Env.OnConfigChange(func (e fsnotify.Event){

	})

	return init
}

func (init *initialize) Get(key string, val interface{}) (res interface{}) {
	res = val
	if init.Env.IsSet(key) {
		res = init.Env.Get(key)
	}

	return
}