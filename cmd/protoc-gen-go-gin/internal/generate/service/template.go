package service

import (
	"math/rand"
	"text/template"
	"time"
)

func init() {
	var err error
	serviceLogicTmpl, err = template.New("serviceLogic").Parse(serviceLogicTmplRaw)
	if err != nil {
		panic(err)
	}
	routerTmpl, err = template.New("serviceRouter").Parse(routerTmplRaw)
	if err != nil {
		panic(err)
	}
	rpcErrCodeTmpl, err = template.New("rpcErrCode").Parse(rpcErrCodeTmplRaw)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
}

var (
	serviceLogicTmpl    *template.Template
	serviceLogicTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	serverNameExampleV1 "moduleNameExample/api/serverNameExample/v1"
	//"moduleNameExample/internal/rpcclient"
)

{{- range .PbServices}}

var _ serverNameExampleV1.{{.Name}}Logicer = (*{{.LowerName}}Client)(nil)

type {{.LowerName}}Client struct {
	// define rpc clients interface here
	// example:
	//	    {{.LowerName}}Cli {{.LowerCutServiceName}}V1.{{.Name}}Client
}

// New{{.Name}}Client create a client
func New{{.Name}}Client() serverNameExampleV1.{{.Name}}Logicer {
	return &{{.LowerName}}Client{
		// example:
		//	    {{.LowerName}}Cli: {{.LowerCutServiceName}}V1.New{{.Name}}Client(rpcclient.Get{{.CutServiceName}}RPCConn()),
	}
}

{{- range .Methods}}

{{.Comment}}
func (c *{{.LowerServiceName}}Client) {{.MethodName}}(ctx context.Context, req *serverNameExampleV1.{{.Request}}) (*serverNameExampleV1.{{.Reply}}, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	//
	//     reply, err := c.{{.LowerServiceName}}Cli.{{.MethodName}}(ctx, &{{.LowerCutServiceName}}V1.{{.Request}}{
{{- range .RequestFields}}
	//     	{{.Name}}: req.{{.Name}},
{{- end}}
	//     })
	//     if err != nil {
	//     	return nil, err
	//     }
	//
	//     return &serverNameExampleV1.{{.Reply}}{
{{- range .ReplyFields}}
	//     	{{.Name}}: reply.{{.Name}},
{{- end}}
	//     }, nil
}

{{- end}}

// ---------- Do not delete or move this split line, this is the merge code marker ----------

{{- end}}
`

	routerTmpl    *template.Template
	routerTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	serverNameExampleV1 "moduleNameExample/api/serverNameExample/v1"
	"moduleNameExample/internal/service"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
{{- range .PbServices}}
		{{.LowerName}}Middlewares(c)
{{- end}}
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
{{- range .PbServices}}
			{{.LowerName}}Router(r, groupPathMiddlewares, singlePathMiddlewares, service.New{{.Name}}Client())
{{- end}}
		})
}

{{- range .PbServices}}

func {{.LowerName}}Router(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService serverNameExampleV1.{{.Name}}Logicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			// set metadata to be passed from http to rpc
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c), // request_id
			//middleware.HeaderAuthorizationKey: c.GetHeader(middleware.HeaderAuthorizationKey),  // authorization
		})
		return metadata.NewOutgoingContext(c, md)
	}

	serverNameExampleV1.Register{{.Name}}Router(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		serverNameExampleV1.With{{.Name}}RPCResponse(),
		serverNameExampleV1.With{{.Name}}Logger(logger.Get()),
		serverNameExampleV1.With{{.Name}}RPCStatusToHTTPCode(
			// Set some error codes to standard http return codes,
			// by default there is already ecode.StatusInternalServerError and ecode.StatusServiceUnavailable
			// example:
			// 	ecode.StatusUnimplemented, ecode.StatusAborted,
		),
		serverNameExampleV1.With{{.Name}}WrapCtx(ctxFn),
	)
}

// you can set the middleware of a route group, or set the middleware of a single route, 
// or you can mix them, pay attention to the duplication of middleware when mixing them, 
// it is recommended to set the middleware of a single route in preference
func {{.LowerName}}Middlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/{{.LowerName}}/:id  will take effect
	// c.setGroupPath("/api/v1/{{.LowerName}}", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
{{- range .Methods}}
	//c.setSinglePath("{{.Method}}", "{{.Path}}", middleware.Auth())
{{- end}}
}

// ---------- Do not delete or move this split line, this is the merge code marker ----------

{{- end}}
`

	rpcErrCodeTmpl    *template.Template
	rpcErrCodeTmplRaw = `// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

{{- range .PbServices}}

// {{.LowerName}} rpc service level error code
var (
	_{{.LowerName}}NO       = {{.RandNumber}} // number range 1~100, if there is the same number, trigger panic.
	_{{.LowerName}}Name     = "{{.LowerName}}"
	_{{.LowerName}}BaseCode = errcode.RCode(_{{.LowerName}}NO)
// --blank line--
{{- range $i, $v := .Methods}}
	Status{{.MethodName}}{{.ServiceName}}   = errcode.NewRPCStatus(_{{.LowerServiceName}}BaseCode+{{$v.AddOne $i}}, "failed to {{.MethodName}} "+_{{.LowerServiceName}}Name)
{{- end}}
	// add +1 to the previous error code
)

// ---------- Do not delete or move this split line, this is the merge code marker ----------

{{- end}}
`
)
