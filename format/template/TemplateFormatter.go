package template

import (
  "bytes"
  "fmt"
  "text/template"

  "github.com/illublank/go-common/format"
)

type TemplateFormatter struct {
  format.Formatter
  formatStr string
  tpl       *template.Template
}

func NewTemplateFormatter(formatStr string) (*TemplateFormatter, error) {

  tpl, err := template.New("common").Parse(formatStr)
  if err != nil {
    return nil, err
  }
  return &TemplateFormatter{
    formatStr: formatStr,
    tpl:       tpl,
  }, nil
}

func (s *TemplateFormatter) Format(params any) string {
  buf := &bytes.Buffer{}
  err := s.tpl.Execute(buf, params)
  if err != nil {
    fmt.Println(err)
  }
  return buf.String()
}
