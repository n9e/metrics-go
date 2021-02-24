package statsdlib

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"syscall"
)

// statsd server
var _statsdServer *statsdServer = &statsdServer{Addr: "127.0.0.1:788"}

type statsdServer struct {
	Addr    string       `json:"addr"`
	UdpAddr *net.UDPAddr `json:"udp_addr"`
}

func (this statsdServer) Init() error {
	wd, _ := syscall.Getwd()

	statsdServerAddrFile := wd + "/.statsd/statsd.cfg.txt"
	c, err := _read(statsdServerAddrFile)
	if err != nil {
		logger{}.Info("use default metrics-agent addr: %s", _statsdServer.Addr)
	} else {
		logger{}.Info("use metrics-agent addr: %s", c)
		_statsdServer.Addr = c
	}

	addr, err := net.ResolveUDPAddr("udp4", _statsdServer.Addr)
	if err != nil {
		return err
	}
	_statsdServer.UdpAddr = addr
	return nil
}

// statsd client
var _statsdClient *statsdClient = &statsdClient{}

type statsdClient struct {
	UdpAddr *net.UDPAddr `json:"udp_addr"`
	UdpConn *net.UDPConn `json:"udp_conn"`
}

func (this statsdClient) Init() error {
	var (
		err error = nil
	)

	_statsdClient.UdpAddr, err = net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return fmt.Errorf("open udp connection error, [err:%s]", err.Error())
	}

	_statsdClient.UdpConn, err = net.ListenUDP("udp4", _statsdClient.UdpAddr)
	if err != nil {
		return fmt.Errorf("open udp connection error, [err:%s]", err.Error())
	}

	logger{}.Info("metric open udp on %s\n", _statsdClient.UdpAddr.String())
	return nil
}

// service meta
var _serviceMeta *serviceMeta = &serviceMeta{}

type serviceMeta struct {
	ServiceName string `json:"service_name"`
	Module      string `json:"module"`
	Cluster     string `json:"cluster"`
	Ns          string `json:"ns"`
}

func (this serviceMeta) Init() error {
	wd, _ := syscall.Getwd()

	// read metas
	deployMetaPath := wd + "/.deploy/"
	serviceNameFile := deployMetaPath + "service.service_name.txt"
	moduleFile := deployMetaPath + "service.module.txt"
	clusterFile := deployMetaPath + "service.cluster.txt"

	c, err := _read(serviceNameFile)
	if err != nil {
		return err
	}
	_serviceMeta.ServiceName = c

	c, err = _read(moduleFile)
	if err != nil {
		return err
	}
	_serviceMeta.Module = c

	c, err = _read(clusterFile)
	if err != nil {
		return err
	}
	_serviceMeta.Cluster = c

	_serviceMeta.Ns = _serviceMeta.Cluster + "." + _serviceMeta.ServiceName

	logger{}.Info("use service meta config: \nservice_name:%s\nmoduel:%s\ncluster:%s",
		_serviceMeta.ServiceName, _serviceMeta.Module, _serviceMeta.Cluster)
	return nil
}

func _read(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("read file error, [file:%s][err:%s]", filename, err.Error())
	}

	cstr := strings.TrimSpace(string(content))
	return cstr, nil
}
