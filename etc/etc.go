package etc

import (
	"strings"
)

func Get(key string) interface{} {
	fields := strings.Split(key, ".")
	cfg := config.Load().(map[string]interface{})
	num := len(fields)
	for i := 0; i < num; i++ {
		val := cfg[fields[i]]
		if i < num-1 {
			c, ok := val.(map[string]interface{})
			if !ok {
				panic("key not found: " + key)
			}
			cfg = c
		} else {
			return val
		}
	}
	return nil
}

func Dev() bool {
	return Bool("system.dev")
}

func Bool(key string) bool {
	return Get(key).(bool)
}

func String(key string) string {
	return Get(key).(string)
}

func Int(key string) int {
	return Get(key).(int)
}
