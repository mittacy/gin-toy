package api

import (
	"bytes"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/base"
	"html/template"
)

var transformTemplate = `
{{- /* delete empty line */ -}}
package dp

import (
	"github.com/gin-gonic/gin"
	"github.com/mittacy/gin-toy/core/response"
)

type {{ .Name }}DP struct{}

func New{{ .Name }}DP() {{ .Name }}DP {
	return {{ .Name }}DP{}
}

func (ctl *{{ .Name }}DP) Get(c *gin.Context, data interface{}) {
	response.Success(c, data)
}

`

type Transform struct {
	Name      string
	NameLower string
	AppName   string
}

func (s *Transform) execute() ([]byte, error) {
	s.Name = base.StringFirstUpper(s.Name)
	s.NameLower = base.StringFirstLower(s.Name)

	buf := new(bytes.Buffer)

	tmpl, err := template.New("validator").Parse(transformTemplate)
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
