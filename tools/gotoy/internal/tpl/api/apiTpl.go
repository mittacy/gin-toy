package api

import (
	"bytes"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/base"
	"html/template"
)

var apiTemplate = `
{{- /* delete empty line */ -}}
package api

import (
	"github.com/gin-gonic/gin"
	"{{ .AppName }}/{{ .TargetDir }}/dp"
	"{{ .AppName }}/{{ .TargetDir }}/service"
	"{{ .AppName }}/{{ .TargetDir }}/validator/{{ .NameLower }}Vdr"
	"github.com/mittacy/gin-toy/core/response"
	"github.com/mittacy/gin-toy/core/singleton"
)

var {{ .Name }} {{ .NameLower }}Api

type {{ .NameLower }}Api struct {
	dp dp.{{ .Name }}DP
}

func init() {
	singleton.Register(func() {
		{{ .Name }} = {{ .NameLower }}Api{
			dp: dp.New{{ .Name }}DP(),
		}
	})
}

func (ctl *{{ .NameLower }}Api) Get(c *gin.Context) {
	req := {{ .NameLower }}Vdr.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErr(c, err)
		return
	}

	{{ .NameLower }}, err := service.{{ .Name }}.GetById(c, req.Id)
	if err != nil {
		response.FailCheckBizErr(c, "查询记录错误", err)
		return
	}

	ctl.dp.Get(c, {{ .NameLower }})
}
`

type Api struct {
	AppName   string
	Name      string
	NameLower string
	TargetDir string
}

func (s *Api) execute() ([]byte, error) {
	s.Name = base.StringFirstUpper(s.Name)
	s.NameLower = base.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("api").Parse(apiTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
