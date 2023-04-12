package main
import (
	"fmt"
    "time"

	"training/localcache"
)

func main() {
	fmt.Println(time.Minute*1)

	cache := localcache.NewLocalCache()

	cache.Set("key1", "value1")

	if val, _ := cache.Get("key1"); val != nil {
		fmt.Println(val)
	}

	cache.Set("key1", "value2")

	if val, _ := cache.Get("key1"); val != nil {
		fmt.Println(val)
	}

	time.Sleep(time.Second * 31)

	if _, error := cache.Get("key1"); error != nil {
		fmt.Println(error)
	}

	if val, _ := cache.Get("key2"); val != nil {
		fmt.Println(val)
	}
}