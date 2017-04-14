package etc

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var mu sync.RWMutex

func init() {
	checkErr(viper.AddRemoteProvider("consul", "localhost:8500", "/etc/default.toml"))
	viper.SetConfigType("toml")
	checkErr(viper.ReadRemoteConfig())
	go watchConfig()
}

func watchConfig() {
	for {
		time.Sleep(10 * time.Second)
		if err := watch(); err != nil {
			log.Println(err)
			continue
		}
	}
}

func watch() error {
	mu.Lock()
	defer mu.Unlock()
	return viper.WatchRemoteConfig()
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Bool(key string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return viper.GetBool(key)
}
