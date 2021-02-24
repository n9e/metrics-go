package statsdlib

import (
	"fmt"
	"strings"
	"time"
)

const (
	maxTagkLen   = 100 // tag key   最大长度
	maxTagvLen   = 100 // tag value 最大长度
	maxTagCnt    = 8   // tag       最大个数
	maxMetricLen = 100 // metric    最大长度

	DefaultRpcVersion = iota
	EnhanceRpcVersion
)

/***************************************************************************
 **********************     业务使用接口区        **************************
 **********************  下述的接口提供给业务使用 **************************
 **************************************************************************/
/**
 * @note
 * rpc统计接口
 * @param string   $metric  指标名
 * @param string   $caller  主调服务标识
 * @param string   $callee  被掉服务标识
 * @param duration $latency 调用耗时
 * @param string   $code    调用结果,  取值 "ok" "0" "200" "201" "203"为成功、其他均为失败
 * @param map      $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func RpcMetric(metric string, caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) error {
	rb := RpcMetricBuilder(metric, caller, callee, latency, code, DefaultRpcVersion, tags...)
	return rb.Push()
}

/**
 * @note
 * counter统计接口
 * @param string $metric  计数指标名称
 * @param map    $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func Counter(metric string, tags ...map[string]string) error {
	return CounterN(metric, 1, tags...)
}
func CounterN(metric string, cnt int, tags ...map[string]string) error {
	cb := CounterNBuilder(metric, cnt, tags...)
	return cb.Push()
}

/**
 * @note
 * gauge统计接口
 * @param string $metric  计数指标名称
 * @param map    $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func Gauge(metric string, value float64, tags ...map[string]string) error {
	gb := GaugeBuilder(metric, value, tags...)
	return gb.Push()
}

/**
 * @note
 * ratio统计接口(统计各类code占比)
 * @param string $metric 计数指标名称
 * @param map	 $tags	 可选的tag, 最多只能有4个
 *
 * @return error
 */
func Ratio(metric string, code string) error {
	rb := RatioBuilder(metric, code)
	return rb.Push()
}

/**
 * @note
 * ratio统计接口升级版, 允许设置计数值
 * @param string $metric 计数指标名称
 * @param map	 $tags	 可选的tag, 最多只能有4个
 *
 * @return error
 */
func RatioN(metric string, code string, cnt int) error {
	rb := RatioBuilder(metric, code, cnt)
	return rb.Push()
}

/**
 * @note
 * percentile统计接口(分位值, 类比Prometheus的Summary)
 * @param string    $metric		 计数指标名称
 * @param float64   $value		 计数值
 * @param []string  $percentiles 分位值
 * @param map		$tags		 可选的tag, 最多只能有4个
 *
 * @return error
 */
func Percentile(metric string, value float64, percentiles []string, tags ...map[string]string) error {
	if len(percentiles) == 0 {
		return fmt.Errorf("percentile not defined")
	}
	pb := PercentileBuilder(metric, value, percentiles, tags...)
	return pb.Push()
}

/***************************************************************************
 **********************     业务定制接口区	      **************************
 **********************  下述接口不常用,酌情使用  **************************
 **************************************************************************/
/**
 * @note
 * counterE统计接口, counter接口增强版, 支持秒级max/min/avg
 * @param string $metric  计数指标名称
 * @param map    $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func CounterE(metric string, tags ...map[string]string) error {
	return CounterNE(metric, 1, tags...)
}
func CounterNE(metric string, cnt int, tags ...map[string]string) error {
	cb := CounterNEBuilder(metric, cnt, tags...)
	return cb.Push()
}

/**
 * @note
 * rpcE统计接口, rpc接口增强版, 支持统计各code的比例
 * @param string   $metric  指标名
 * @param string   $caller  主调服务标识
 * @param string   $callee  被掉服务标识
 * @param duration $latency 调用耗时
 * @param string   $code    调用结果,  取值 "ok" "0" "200" "201" "203"为成功、其他均为失败
 * @param map      $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func RpcMetricE(metric string, caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) error {
	rb := RpcMetricBuilder(metric, caller, callee, latency, code, EnhanceRpcVersion, tags...)
	return rb.Push()
}

/***************************************************************************
 **********************      即将被废弃的接口   **************************
 **************************************************************************/
/**
 * @note
 * rpc统计接口
 * @param string   $caller  主调服务标识
 * @param string   $callee  被掉服务标识
 * @param duration $latency 调用耗时
 * @param string   $code    调用结果,  取值 "ok" "0" "200" "201" "203"为成功、其他均为失败
 * @param map      $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func Rpc(caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) error {
	rb := RpcBuilder(caller, callee, latency, code, tags...)
	return rb.Push()
}

/**
 * @note
 * rpcE统计接口, rpc接口增强版, 支持统计各code的比例
 * @param string   $caller  主调服务标识
 * @param string   $callee  被掉服务标识
 * @param duration $latency 调用耗时
 * @param string   $code    调用结果,  取值 "ok" "0" "200" "201" "203"为成功、其他均为失败
 * @param map      $tags    可选的tag, 最多只能有4个
 *
 * @return error
 */
func RpcE(caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) error {
	rb := RpcEBuilder(caller, callee, latency, code, tags...)
	return rb.Push()
}

/***************************************************************************
 **********************        metric底层接口区   **************************
 *******    下述的接口不建议直接在业务中使用，需要包装   *******************
 **************************************************************************/

/**
* @note
* 设置default nid, 只能在初始化时调用(如非必要,请勿调用!)
* @param string $nid
*
* @return void
 */
func SetDefaultNs(ns string) {
	_serviceMeta.Ns = ns
}

func RpcBuilder(caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) *metricBuilder {
	return RpcMetricBuilder("rpc", caller, callee, latency, code, DefaultRpcVersion, tags...)
}

func RpcEBuilder(caller string, callee string, latency time.Duration, code interface{}, tags ...map[string]string) *metricBuilder {
	return RpcMetricBuilder("rpc", caller, callee, latency, code, EnhanceRpcVersion, tags...)
}

func RpcMetricBuilder(metric string, caller string, callee string, latency time.Duration, code interface{}, version int, tags ...map[string]string) *metricBuilder {
	pos := strings.Index(caller, "?")
	if pos > 0 {
		caller = caller[:pos]
	}
	if len(caller) > maxTagkLen {
		caller = caller[:maxTagkLen]
	}

	pos = strings.Index(callee, "?")
	if pos > 0 {
		callee = callee[:pos]
	}
	if len(callee) > maxTagkLen {
		callee = callee[:maxTagkLen]
	}

	aggr := "rpc"
	if version == EnhanceRpcVersion {
		aggr = "rpce"
	}
	rb := metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Tag("caller", caller).Tag("callee", callee).Agg(
		fmt.Sprintf("%v,%v", latency.Nanoseconds()/1000000, code), aggr)
	if len(tags) > 0 {
		rb.AddTags(tags[0])
	}

	return rb
}

func CounterNBuilder(metric string, cnt int, tags ...map[string]string) *metricBuilder {
	cb := metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(fmt.Sprintf("%d", cnt), "c")
	if len(tags) > 0 {
		cb.AddTags(tags[0])
	}
	return cb
}

func CounterNEBuilder(metric string, cnt int, tags ...map[string]string) *metricBuilder {
	cb := metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(fmt.Sprintf("%d", cnt), "ce")
	if len(tags) > 0 {
		cb.AddTags(tags[0])
	}
	return cb
}

func GaugeBuilder(metric string, value float64, tags ...map[string]string) *metricBuilder {
	gb := metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(fmt.Sprintf("%f", value), "g")
	if len(tags) > 0 {
		gb.AddTags(tags[0])
	}
	return gb
}

func RatioBuilder(metric string, code string, cnt ...int) *metricBuilder {
	if len(cnt) == 1 {
		return metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(fmt.Sprintf("%d,%s", cnt[0], code), "rt")
	}

	return metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(code, "rt")
}

func PercentileBuilder(metric string, value float64, percentiles []string, tags ...map[string]string) *metricBuilder {
	aggr := strings.Join(percentiles, ",")
	pb := metricBuilder{}.Name(metric).Ns(_serviceMeta.Ns).Agg(fmt.Sprintf("%f", value), aggr)
	if len(tags) > 0 {
		pb.AddTags(tags[0])
	}
	return pb
}

// builder
type metricBuilder struct {
	Namespace   string
	Metric      string
	NsAndMetric string
	Aggregator  string
	Tags        map[string]string
	Value       string
}

func (this metricBuilder) Name(metric string) *metricBuilder {
	return &metricBuilder{
		Metric: metric,
	}
}

func (self *metricBuilder) Ns(namespace string) *metricBuilder {
	self.Namespace = namespace
	return self
}

func (self *metricBuilder) Tag(tagK string, tagV string) *metricBuilder {
	if self.Tags == nil {
		self.Tags = map[string]string{}
	}
	self.Tags[tagK] = tagV
	return self
}

func (self *metricBuilder) AddTags(tags map[string]string) *metricBuilder {
	if self.Tags == nil {
		self.Tags = map[string]string{}
	}
	for k, v := range tags {
		self.Tags[k] = v
	}
	return self
}

func (self *metricBuilder) Agg(value string, aggregator string) *metricBuilder {
	self.Value = value
	self.Aggregator = aggregator
	return self
}

func (self *metricBuilder) Check() error {

	// check metric
	metricLen := len(self.Metric)
	if metricLen == 0 {
		return fmt.Errorf("empty metric")
	}
	if metricLen > maxMetricLen {
		return fmt.Errorf("metric too long: %s", self.Metric)
	}

	// check tags
	tagCnt := len(self.Tags)
	if tagCnt > maxTagCnt {
		return fmt.Errorf("too many tags")
	}
	if tagCnt != 0 {
		for k, v := range self.Tags {
			ksize := len(k)
			if ksize == 0 {
				return fmt.Errorf("empty tagk")
			}
			if ksize > maxTagkLen {
				return fmt.Errorf("tagk too long: %s", k)
			}

			vsize := len(v)
			if vsize == 0 {
				return fmt.Errorf("empty tagv")
			}
			if vsize > maxTagvLen {
				return fmt.Errorf("tagv too long: %s", v)
			}
		}
	}

	return nil
}

func (self *metricBuilder) Build() string {
	lines := []string{self.Value}                                        // value
	self.NsAndMetric = fmt.Sprintf("%s/%s", self.Namespace, self.Metric) // ns/metric
	lines = append(lines, self.NsAndMetric)
	for k, v := range self.Tags {
		lines = append(lines, fmt.Sprintf("%v=%v", k, v)) // tags
	}
	lines = append(lines, self.Aggregator) // aggregator
	return strings.Join(lines, "\n")
}

func (self *metricBuilder) Push() error {
	defer func() {
		if r := recover(); r != nil {
			logger{}.Erro("metrics push panic: %v\n", r)
		}
	}()
	// check
	err := self.Check()
	if err != nil {
		return err
	}

	// build
	body := self.Build()

	// send
	server := _statsdServer.UdpAddr
	client := _statsdClient.UdpConn
	if server == nil {
		return fmt.Errorf("server not init")
	}
	if client == nil {
		return fmt.Errorf("client not init")
	}
	_, _, err = client.WriteMsgUDP([]byte(body), nil, server)
	if nil != err {
		return err
	}

	return nil
}
