package statsdlib

import (
	"fmt"
)

func init() {
	err := logger{}.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = statsdServer{}.Init()
	if err != nil {
		logger{}.Erro(err.Error())
		return
	}

	err = statsdClient{}.Init()
	if err != nil {
		logger{}.Erro(err.Error())
		return
	}

	err = serviceMeta{}.Init()
	if err != nil {
		logger{}.Erro(fmt.Sprintf("service meta init error: %s", err.Error()))
		return
	}
}
