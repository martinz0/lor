package etc

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var config atomic.Value

func init() {
	checkErr(viper.AddRemoteProvider("consul", "localhost:8500", "/etc/default.toml"))
	viper.SetConfigType("toml")
	checkErr(viper.ReadRemoteConfig())
	config.Store(viper.AllSettings())
	go watchConfig()
}

func watchConfig() {
	for {
		time.Sleep(1 * time.Second)
		if err := watch(); err != nil {
			log.Println(err)
			continue
		}
	}
}

func watch() error {
	if err := viper.WatchRemoteConfig(); err != nil {
		return err
	}
	config.Store(viper.AllSettings())
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
