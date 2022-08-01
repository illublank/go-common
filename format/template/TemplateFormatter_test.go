package template_test

import (
  "fmt"
  "testing"

  "github.com/illublank/go-common/format/template"
)

func Test(t *testing.T) {
  tpl, _ := template.NewTemplateFormatter("aaaa{{.name}}")
  params := map[string]any{"name": "111"}

  fmt.Println(tpl.Format(params))
}
