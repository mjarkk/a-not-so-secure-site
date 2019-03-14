package templates

import (
	"bytes"
	"errors"
	"path"
	"strings"
	"text/template"

	"unknownclouds.com/root/kpn-xml-to-conf/fs"
)

// GetTemplate returns a template
func GetTemplate(name string, data interface{}) (string, error) {
	temp, err := OpenTemplateFile(name)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBufferString("")
	err = temp.Execute(buf, data)
	return buf.String(), err
}

// OpenTemplateFile opens a template file and returns is as a template type
func OpenTemplateFile(fileName string) (*template.Template, error) {
	if strings.Contains(fileName, "/") || strings.Contains(fileName, ".") || strings.Contains(fileName, "\\") || strings.Contains(fileName, "~") {
		return nil, errors.New("not a valid file name")
	}

	content, err := fs.FileContents(path.Join("templates", fileName+".html"))
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("test").Parse(content)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
