package exec

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/ggicci/httpin"
)

func WebServiceV1Alpha1(urlStr string) (*restful.WebService, error) {
	sCmd, err := OpenSSHCmd(urlStr)
	if err != nil {
		return nil, fmt.Errorf("open ssh cmd: %v", err)
	}
	r := &CmdResource{SSHCmd: sCmd}
	return r.WebServiceV1Alpha1(), nil
}

type CmdResource struct {
	*SSHCmd
}

func (r *CmdResource) WebServiceV1Alpha1() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/api/v1alpha1/cmds").
		Consumes(restful.MIME_OCTET).
		Produces(restful.MIME_JSON)
	ws.Route(ws.
		POST("/{path}:run").
		Filter(restful.HttpMiddlewareHandlerToFilter(httpin.NewInput(Cmd{}))).
		To(r.runCmd))
	return ws
}

func (r *CmdResource) WebServiceV1Beta1() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/api/v1beta1/cmds").
		Consumes(restful.MIME_OCTET).
		Produces(restful.MIME_JSON)
	ws.Route(ws.
		POST("/{path}:run").
		Filter(restful.HttpMiddlewareHandlerToFilter(httpin.NewInput(Cmd{}))).
		To(r.runCmd))
	return ws
}

func (r *CmdResource) runCmd(req *restful.Request, res *restful.Response) {
	ctx := req.Request.Context()
	cmd := ctx.Value(httpin.Input).(*Cmd)
	cmd.Path = req.PathParameter("path")
	cRes, err := r.Run(ctx, cmd)
	if err != nil {
		code := http.StatusInternalServerError
		res.WriteError(code, restful.NewError(code, err.Error()))
		return
	}
	res.WriteEntity(cRes)
}
