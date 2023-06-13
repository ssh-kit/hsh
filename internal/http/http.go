package http

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"github.com/batx-dev/batcmd/internal/exec"
)

func ListenAndServe(addr string, urlStr string) error {
	ewsv1a1, err := exec.WebServiceV1Alpha1(urlStr)
	if err != nil {
		return fmt.Errorf("open web service: %v", err)
	}

	cont := restful.NewContainer()
	cont.Add(ewsv1a1)

	return http.ListenAndServe(addr, cont)
}
