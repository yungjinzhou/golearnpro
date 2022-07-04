package testing


import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)


func TestRedisConnect(t *testing.T) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{MasterName: "redisMaster", SentinelAddrs: []string{"192.168.231.35:26379","192.168.231.39:26379","192.168.231.44:26379"}})
	fmt.Println(rdb)
	t.Log(rdb)
	my, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接出错：", err)
		return
	}
	fmt.Printf("%s", my)
}

