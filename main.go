package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	statsd "github.com/n9e/metrics-go/statsdlib"
)

func init() {
	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()+os.Getppid()))

	// 全局设置树节点ID，一般是模块节点的ID
	statsd.SetDefaultNs("18")
}

func rpcWrapper(rpcSrv string) {
	// 业务逻辑处理之前，获取一下当前时间戳
	startT := time.Now()

	// 开始处理业务逻辑
	code := doRPC(rpcSrv)

	// 计算时间差，即耗时数据
	latency := time.Now().Sub(startT)

	// caller 姑且全部写成all，不用变
	caller := "all"

	// callee 写成rpc方法名，如果是http的话可以写成url path，比如/api/v1/user/login
	callee := rpcSrv

	// 通过udp上报监控数据
	fmt.Printf(">>> rpc.user_service caller: %s, callee: %s, code: %s, latency: %d\n", caller, callee, code, latency.Microseconds())

	// 这里假设我这个微服务的名字是user_service
	statsd.RpcMetric("rpc.user_service", caller, callee, latency, code)
}

// 返回的string表示code， "ok" "0" "200" "201" "203" 都表示成功，其他都表示失败
func doRPC(rpcSrv string) string {
	// 这里是具体业务逻辑，用一个sleep语句来模仿
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	return "200"
}

func main() {
	for i := 0; i < 10000; i++ {
		// 这里是在模仿user_login这个rpc接口的访问情况
		rpcWrapper("user_login")
	}
}
