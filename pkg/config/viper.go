package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Viper struct {
	viper  *viper.Viper
	Notify chan bool
}

func NewViper() *Viper {
	return &Viper{
		viper:  viper.New(),
		Notify: make(chan bool, 1),
	}
}

func (v *Viper) Load(path string) error {

	if path != "" {
		v.viper.SetConfigFile(path) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		v.viper.AddConfigPath("config") // 如果没有指定配置文件，则解析默认的配置文件
		v.viper.SetConfigName("config")
	}
	v.viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	if err := v.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		}
		return err
	}

	return nil
}

func (v *Viper) Parse(conf interface{}) (interface{}, error) {

	err := v.viper.Unmarshal(conf)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return conf, nil
}

func (v *Viper) WatchConfig() {
	go func() {
		v.viper.WatchConfig()
		v.viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Config file changed: %s", e.Name)
			//通知配置文件更新
			v.Notify <- true
		})
	}()
}
