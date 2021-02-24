## 概述
golang的业务程序，可以通过statsdlib，接入metrics统计服务。

## 使用
业务代码引入statsdlib包，在代码片段中调用statsdlib的API。比如，用户想要统计rpc请求数据，可以调用`RpcMetric`方法。

```go
package xxx

import (
	"time"
	statsd "github.com/n9e/metrics-go/statsdlib"
)

func foo(rpcSrv string) {
	startT := time.Now()

	result, err := doRpc(rpcSrv)

	latency := time.Now().Sub(startT)
	caller := "foo"
	callee := rpcSrv

	if ( err!=nil ) {
		statsd.RpcMetric("rpc", caller, callee, latency, "rpcFunc.error")
		return		
	} else {
		statsd.RpcMetric("rpc", caller, callee, latency, "ok")		
	}

	....
}

```

## API
几个常用接口，如下。

|接口名称|例子|使用场景|
|:----|:----|:---|
|Counter(name string)|`// 统计调用次数加1`<br/>`Counter("api.checkhealth") `|Counter输出一个统计周期内的计数累加和|
|RpcMetric(metric,caller,callee string, latency time.Duration, code string)|`// 统计接口调用质量`<br/>`RpcMetric("rpc","caller", "callee", time.Second*1, "ok") `|RpcMetric输出一个统计周期内 (1)不同code的统计计数、错误率、访问质量 (2)所有code的统计计数、错误率、访问质量。code取值"ok", "0", "200", "201", "203"代表成功，其余均代表失败。Rpc产生的监控指标项，包括`rpc.counter, rpc.error.ratio, rpc.latency`。|