package testing
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/go-redis/redis"
//	"net"
//)

//
//import (
//	"fmt"
//	"github.com/go-redis/redis"
//	"testing"
//)
//
//
//func TestRedisConnect(t *testing.T) {
//
//	rdb := redis.NewFailoverClient(&redis.FailoverOptions{MasterName: "mymaster", SentinelAddrs: []string{"192.168.230.106:26379", "192.168.230.109:26379", "192.168.230.105:26379"}})
//	fmt.Println(rdb)
//	rdb.
//	fmt.Println(rdb.Options())
//	fmt.Println(rdb.Info())
//
//	//fmt.Printf("%v", rdb.Info())
//}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/go-redis/redis"
//	"net"
//)
//
//// TCP 服务端
//func process(conn net.Conn) {
//	// 函数执行完之后关闭连接
//	defer conn.Close()
//	// 输出主函数传递的conn可以发现属于*TCPConn类型, *TCPConn类型那么就可以调用*TCPConn相关类型的方法, 其中可以调用read()方法读取tcp连接中的数据
//	fmt.Printf("服务端: %T\n", conn)
//	for i:=0;i<2;i++ {
//		var buf [128]byte
//		// 将tcp连接读取到的数据读取到byte数组中, 返回读取到的byte的数目
//		n, err := conn.Read(buf[:])
//		if err != nil {
//			// 从客户端读取数据的过程中发生错误
//			fmt.Println("read from client failed, err:", err)
//			break
//		}
//		recvStr := string(buf[:n])
//		fmt.Println("服务端收到客户端发来的数据：", recvStr)
//		// 由于是tcp连接所以双方都可以发送数据, 下面接收服务端发送的数据这样客户端也可以收到对应的数据
//		tt := getRedis()
//		fmt.Printf("%+v", tt)
//		jsonBytes, err := json.Marshal(tt)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Printf("%+v", jsonBytes)
//		// 向当前建立的tcp连接发送数据, 客户端就可以收到服务端发送的数据
//		conn.Write(jsonBytes)
//		fmt.Println("success---")
//	}
//}
//
//func getRedis() (r *redis.Client) {
//	rdb := redis.NewFailoverClient(&redis.FailoverOptions{MasterName: "mymaster", SentinelAddrs: []string{"192.168.230.109:26379", "192.168.230.105:26379", "192.168.230.106:26379"}})
//	return rdb
//}
//
//func main() {
//	// 监听当前的tcp连接
//	listen, err := net.Listen("tcp", "127.0.0.1:9001")
//	fmt.Printf("服务端: %T=====\n", listen)
//	if err != nil {
//		fmt.Println("listen failed, err:", err)
//		return
//	}
//	for {
//		conn, err := listen.Accept() // 建立连接
//		fmt.Println("当前建立了tcp连接")
//		if err != nil {
//			fmt.Println("accept failed, err:", err)
//			continue
//		}
//		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
//		go process(conn)
//	}
//}