package tmplUtil

import (
	"bytes"
	"errors"
	"go-utils/structUtil"
	"html/template"
	"os"
)

type Tmpl struct {
	st *structUtil.Struct
}

func NewTemp(st *structUtil.Struct) *Tmpl {
	return &Tmpl{st: st}
}
func (t *Tmpl) Write(TmplFile string, out string) error {
	if _, err := os.Stat(out); err == nil {
		return errors.New("文件已存在")
	}
	content := &bytes.Buffer{}
	err := template.Must(template.ParseFiles(TmplFile)).Execute(content, t.st)
	if err != nil {
		return err
	}
	return os.WriteFile(out, content.Bytes(), 0644)
}
