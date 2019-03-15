package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Base BaseConf `mapstructure:"base"`
}

type BaseConf struct {
	Pidfile string `mapstructure:"pidfile"`
	Version string `mapstructure:"version"`
}

var (
	Conf     *ApplicationConfig
	confPath string
)

func init() {
	//log configuration
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	//load configuration
	flag.StringVar(&confPath, "d", "../configs", "set logic config file path")

	log.WithField("confPath", confPath).Debugln("read configuration from confPath")

	Conf = &ApplicationConfig{
		Base: BaseConf{
			Pidfile: "/tmp/a.pid",
		},
	}

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(confPath)

	if err := viper.ReadInConfig(); err != nil {

		log.WithField("confPath", confPath).WithError(err).Panicln("unable find configPath")
	}

	for k, v := range viper.AllKeys() {
		fmt.Println(k, v)
	}

	log.Infoln(viper.GetString("base.pidfile"))

	if err := viper.Unmarshal(&Conf); err != nil {
		log.WithError(err).Panicln("unable decode into struct")
	}
}

func main() {

	log.Debugln("start")

	log.WithField("pidfile", Conf.Base.Pidfile).WithField("version", Conf.Base.Version).Infoln("configuration")
}
