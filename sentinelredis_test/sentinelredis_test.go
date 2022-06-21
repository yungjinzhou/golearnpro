package testing


import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)


func TestRedisConnect(t *testing.T) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{MasterName: "mymaster", SentinelAddrs: []string{"192.168.230.106:26003", "192.168.230.106:26004", "192.168.230.106:26005"}})
	fmt.Println(rdb)
	t.Log(rdb)
	my, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接出错：", err)
		return
	}
	fmt.Printf("%s", my)
}

