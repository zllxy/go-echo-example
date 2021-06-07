package conf

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-echo-example/pkg/db"
	"go-echo-example/pkg/file"
	logger "go-echo-example/pkg/log"
	"log"
	"time"
)

const DefaultPath = "./configs/config.yaml"

type Conf struct {
	App    AppConf
	Logger logger.Conf
	Db     db.Conf
}

type AppConf struct {
	Name    string
	Version string
	Debug   bool
	Addr    string
	Timeout time.Duration
}

type LoadConf struct {
	*viper.Viper
	Notify chan bool
}

func NewLoadConf(notify chan bool) *LoadConf {
	return &LoadConf{
		Viper:  viper.New(),
		Notify: notify,
	}
}

func (l *LoadConf) Load(path string) error {
	if path == "" {
		path = DefaultPath
	}
	l.SetConfigFile(path)
	l.SetConfigType(file.GetExt(path)) // 设置配置文件格式

	if err := l.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("conf file not found")
			return errors.New("conf file not found")
		}
		return err
	}

	return nil
}

func (l *LoadConf) Parse(c *Conf) (*Conf, error) {
	err := l.Unmarshal(c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return c, nil
}

func (l *LoadConf) Watch() {
	go func() {
		l.WatchConfig()
		l.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Conf file changed: %s", e.Name)
			//通知配置文件更新
			l.Notify <- true
		})
	}()
}
