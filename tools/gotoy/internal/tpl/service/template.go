package service

import (
	"bytes"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/base"
	"text/template"
)

var serviceTemplate = `
{{- /* delete empty line */ -}}
package service

import (
	"github.com/gin-gonic/gin"
	"{{ .AppName }}/{{ .TargetDir }}/data"
	"{{ .AppName }}/{{ .TargetDir }}/model"
	"github.com/mittacy/gin-toy/core/singleton"
	"github.com/pkg/errors"
)

// 一般情况下service应该只引用并控制自己的data模型，需要其他服务的功能请service.Xxx调用服务而不是引入其他data模型

// {{ .Name }} 服务说明注释
var {{ .Name }} {{ .NameLower }}Service

type {{ .NameLower }}Service struct {
	data data.{{ .Name }}
}

func init() {
	singleton.Register(func() {
		{{ .Name }} = {{ .NameLower }}Service{
			data: data.New{{ .Name }}(),
		}
	})
}

func (ctl *{{ .NameLower }}Service) GetById(c *gin.Context, id int64) (*model.{{ .Name }}, error) {
	{{ .NameLower }}, err := ctl.data.Get(c, id)
	if err != nil {
		return nil, errors.WithMessage(err, "查询记录错误")
	}

	return {{ .NameLower }}, nil
}
`

type Service struct {
	AppName   string
	Name      string
	NameLower string
	TargetDir string
}

func (s *Service) execute() ([]byte, error) {
	s.Name = base.StringFirstUpper(s.Name)
	s.NameLower = base.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
