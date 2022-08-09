package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//)

// fetch html head data

//import (
//	"fmt"
//	"net/http"
//)
//
//var urls = []string{
//	"http://www.baidu.com",
//	"http://sougou.com",
//}
//
//func main(){
//	for _, url := range urls {
//		resp, err := http.Head(url)
//		if err != nil {
//			fmt.Println("the error is ", url, err)
//		}
//		fmt.Println(url, resp.Status)
//	}
//}


// 获取网页
//
//func main() {
//	rep, err := http.Get("http://www.baidu.com")
//	CheckError(err)
//	data, err := ioutil.ReadAll(rep.Body)
//	CheckError(err)
//	fmt.Printf("Got data %q", data)
//}
//
//
//func CheckError(err error) {
//	if err != nil {
//		log.Fatalf("get url err: %v", err)
//	}
//}



// 使用xml

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status Status
}

func main() {
	// 发起请求查询推特 Goodland 用户的状态
	response, _ := http.Get("http://twitter.com/users/Googland.xml")
	// 初始化 XML 返回值的结构
	user := User{xml.Name{"", "user"}, Status{""}}
	// 将 XML 解析为我们的结构
	xml.Unmarshal(response.Body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}



