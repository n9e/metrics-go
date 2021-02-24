package statsdlib

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	metric := "niean.test"
	cnt := 2
	cb := CounterNBuilder(metric, cnt)

	if cb.Metric != metric {
		t.Errorf("bad metric: %s", cb.Metric)
	}
	if cb.Value != fmt.Sprintf("%d", cnt) {
		t.Errorf("bad value: %s", cb.Value)
	}
	if cb.Aggregator != "c" {
		t.Errorf("bad aggregator: %s", cb.Aggregator)
	}

	// ns
	namepace := "1"
	if cb.Namespace != "" {
		t.Errorf("bad default ns: %s", cb.Namespace)
	}
	cb.Ns(namepace)
	if cb.Namespace != namepace {
		t.Errorf("bad ns: %s", cb.Namespace)
	}

	// tags
	tags := map[string]string{"k1": "v1", "k2": "v2"}
	if len(cb.Tags) != 0 {
		t.Errorf("bad default tags: %v", cb.Tags)
	}
	cb.AddTags(tags)
	if !reflect.DeepEqual(cb.Tags, tags) {
		t.Errorf("bad tags: %v", cb.Tags)
	}
}

func TestCounterN(t *testing.T) {
	metric := "hujter.test"
	cnt := 2
	cb := CounterNEBuilder(metric, cnt)

	if cb.Metric != metric {
		t.Errorf("bad metric: %s", cb.Metric)
	}
	if cb.Value != fmt.Sprintf("%d", cnt) {
		t.Errorf("bad value: %s", cb.Value)
	}
	if cb.Aggregator != "ce" {
		t.Errorf("bad aggregator: %s", cb.Aggregator)
	}

	// ns
	namepace := "1"
	if cb.Namespace != "" {
		t.Errorf("bad default ns: %s", cb.Namespace)
	}
	cb.Ns(namepace)
	if cb.Namespace != namepace {
		t.Errorf("bad ns: %s", cb.Namespace)
	}

	// tags
	tags := map[string]string{"k1": "v1", "k2": "v2"}
	if len(cb.Tags) != 0 {
		t.Errorf("bad default tags: %v", cb.Tags)
	}
	cb.AddTags(tags)
	if !reflect.DeepEqual(cb.Tags, tags) {
		t.Errorf("bad tags: %v", cb.Tags)
	}
}

func TestGauge(t *testing.T) {
	metric := "gauge.test"
	val := 94.6
	gauge := GaugeBuilder(metric, val)

	if gauge.Metric != metric {
		t.Errorf("bad metric: %s", gauge.Metric)
	}
	if gauge.Value != fmt.Sprintf("%f", val) {
		t.Errorf("bad value: %s", gauge.Value)
	}
	if gauge.Aggregator != "g" {
		t.Errorf("bad aggregator: %s", gauge.Aggregator)
	}

	// ns
	namepace := "1"
	if gauge.Namespace != "" {
		t.Errorf("bad default ns: %s", gauge.Namespace)
	}
	gauge.Ns(namepace)
	if gauge.Namespace != namepace {
		t.Errorf("bad ns: %s", gauge.Namespace)
	}

	// tags
	tags := map[string]string{"k1": "v1", "k2": "v2"}
	if len(gauge.Tags) != 0 {
		t.Errorf("bad default tags: %v", gauge.Tags)
	}
	gauge.AddTags(tags)
	if !reflect.DeepEqual(gauge.Tags, tags) {
		t.Errorf("bad tags: %v", gauge.Tags)
	}
}

func TestRatio(t *testing.T) {
	metric := "ratio.test"
	rt := RatioBuilder(metric, "ok")

	if rt.Metric != metric {
		t.Errorf("bad metric: %s", rt.Metric)
	}
	if rt.Value != "ok" {
		t.Errorf("bad value: %s", rt.Value)
	}
	if rt.Aggregator != "rt" {
		t.Errorf("bad aggregator: %s", rt.Aggregator)
	}

	// ns
	namepace := "1"
	if rt.Namespace != "" {
		t.Errorf("bad default ns: %s", rt.Namespace)
	}
	rt.Ns(namepace)
	if rt.Namespace != namepace {
		t.Errorf("bad ns: %s", rt.Namespace)
	}
}

func TestPercentile(t *testing.T) {
	metric := "percentile.test"
	val := 88.0
	percentiles := []string{"p99", "p75"}
	pt := PercentileBuilder(metric, val, percentiles)

	if pt.Metric != metric {
		t.Errorf("bad metric: %s", pt.Metric)
	}
	if pt.Value != fmt.Sprintf("%f", val) {
		t.Errorf("bad value: %s", pt.Value)
	}
	if pt.Aggregator != strings.Join(percentiles, ",") {
		t.Errorf("bad aggregator: %s", pt.Aggregator)
	}

	// ns
	namepace := "1"
	if pt.Namespace != "" {
		t.Errorf("bad default ns: %s", pt.Namespace)
	}
	pt.Ns(namepace)
	if pt.Namespace != namepace {
		t.Errorf("bad ns: %s", pt.Namespace)
	}

	// tags
	tags := map[string]string{"k1": "v1", "k2": "v2"}
	if len(pt.Tags) != 0 {
		t.Errorf("bad default tags: %v", pt.Tags)
	}
	pt.AddTags(tags)
	if !reflect.DeepEqual(pt.Tags, tags) {
		t.Errorf("bad tags: %v", pt.Tags)
	}
}

func TestRpc(t *testing.T) {
	// Rpc
	{
		metric := "rpc"
		caller := "caller"
		callee := "callee"
		code := "ok"
		cnt := 2
		cb := RpcBuilder(caller, callee, time.Duration(cnt*1000000), code)

		if cb.Metric != metric {
			t.Errorf("bad metric: %s", cb.Metric)
		}
		if cb.Value != fmt.Sprintf("%v,%v", cnt, code) {
			t.Errorf("bad value: %s", cb.Value)
		}
		if cb.Aggregator != "rpc" {
			t.Errorf("bad aggregator: %s", cb.Aggregator)
		}
		v, ok := cb.Tags["caller"]
		if !(ok && v == caller) {
			t.Errorf("bad caller")
		}
		v, ok = cb.Tags["callee"]
		if !(ok && v == callee) {
			t.Errorf("bad callee")
		}

		// ns
		namepace := "1"
		if cb.Namespace != "" {
			t.Errorf("bad default ns: %s", cb.Namespace)
		}
		cb.Ns(namepace)
		if cb.Namespace != namepace {
			t.Errorf("bad ns: %s", cb.Namespace)
		}

		// tags
		tags := map[string]string{"k1": "v1", "k2": "v2"}
		cb.AddTags(tags)
		tags["caller"] = caller
		tags["callee"] = callee
		if !reflect.DeepEqual(cb.Tags, tags) {
			t.Errorf("bad tags: %v", cb.Tags)
		}
	}
	// RpcMetric
	{
		metric := "rpctest"
		caller := "caller"
		callee := "callee"
		code := "ok"
		cnt := 2
		cb := RpcMetricBuilder(metric, caller, callee, time.Duration(cnt*1000000), code, DefaultRpcVersion)

		if cb.Metric != metric {
			t.Errorf("bad metric: %s", cb.Metric)
		}
		if cb.Value != fmt.Sprintf("%v,%v", cnt, code) {
			t.Errorf("bad value: %s", cb.Value)
		}
		if cb.Aggregator != "rpc" {
			t.Errorf("bad aggregator: %s", cb.Aggregator)
		}
		v, ok := cb.Tags["caller"]
		if !(ok && v == caller) {
			t.Errorf("bad caller")
		}
		v, ok = cb.Tags["callee"]
		if !(ok && v == callee) {
			t.Errorf("bad callee")
		}

		// ns
		namepace := "1"
		if cb.Namespace != "" {
			t.Errorf("bad default ns: %s", cb.Namespace)
		}
		cb.Ns(namepace)
		if cb.Namespace != namepace {
			t.Errorf("bad ns: %s", cb.Namespace)
		}

		// tags
		tags := map[string]string{"k1": "v1", "k2": "v2"}
		cb.AddTags(tags)
		tags["caller"] = caller
		tags["callee"] = callee
		if !reflect.DeepEqual(cb.Tags, tags) {
			t.Errorf("bad tags: %v", cb.Tags)
		}
	}
}

func TestCheck(t *testing.T) {
	bads := ""
	for i := 0; i < 51; i++ {
		bads = fmt.Sprintf("%s.%d", bads, i)
	}
	if len(bads) < 101 {
		t.Errorf("gen bads error")
	}

	// ns
	cb := CounterNBuilder("m", 1).Ns("")
	err := cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "empty ns")) {
		t.Errorf("bad ns check")
	}

	// metric
	ns := "1"
	cb = CounterNBuilder("", 1).Ns(ns)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "empty metric")) {
		t.Errorf("bad metric check")
	}
	cb = CounterNBuilder(bads, 1).Ns(ns)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "metric too long")) {
		t.Errorf("bad metric check")
	}

	// tags
	cb = CounterNBuilder("m", 1).Ns(ns)
	tags := map[string]string{"k": "v"}
	cb.AddTags(tags)
	err = cb.Check()
	if err != nil {
		t.Errorf("bad tags check")
	}

	cb = CounterNBuilder("m", 1).Ns(ns)
	tags = map[string]string{"": "v"}
	cb.AddTags(tags)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "empty tagk")) {
		t.Errorf("bad tags check")
	}

	cb = CounterNBuilder("m", 1).Ns(ns)
	tags = map[string]string{bads: "v"}
	cb.AddTags(tags)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "tagk too long")) {
		t.Errorf("bad tags check")
	}

	cb = CounterNBuilder("m", 1).Ns(ns)
	tags = map[string]string{"k": ""}
	cb.AddTags(tags)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "empty tagv")) {
		t.Errorf("bad tags check")
	}

	cb = CounterNBuilder("m", 1).Ns(ns)
	tags = map[string]string{"k": bads}
	cb.AddTags(tags)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "tagv too long")) {
		t.Errorf("bad tags check")
	}

	cb = CounterNBuilder("m", 1).Ns(ns)
	tags = map[string]string{}
	for i := 0; i < 9; i++ {
		kv := fmt.Sprintf("%d", i)
		tags[kv] = kv
	}
	cb.AddTags(tags)
	err = cb.Check()
	if !(err != nil && strings.Contains(err.Error(), "too many tags")) {
		t.Errorf("bad tags check")
	}
}

func TestPush(t *testing.T) {
	SetDefaultNs("1")

	Counter("niean.test")
	Counter("niean.test", map[string]string{"k": "v"})
	time.Sleep(time.Second * time.Duration(1))

	CounterN("niean.test", 2)
	CounterN("niean.test", 2, map[string]string{"k": "v"})
	time.Sleep(time.Second * time.Duration(1))

	Rpc("caller", "callee", time.Duration(10000000), "ok")
	Rpc("caller", "callee", time.Duration(10000000), "ok", map[string]string{"k": "v"})
	time.Sleep(time.Second * time.Duration(1))

	RpcMetric("rpctest", "caller", "callee", time.Duration(10000000), "ok")
	RpcMetric("rpctest", "caller", "callee", time.Duration(10000000), "ok", map[string]string{"k": "v"})
	time.Sleep(time.Second * time.Duration(1))
}
