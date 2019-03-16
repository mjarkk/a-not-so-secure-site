package templates

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path"
	"strings"
	"text/template"
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

	content, err := FileContents(path.Join("templates", fileName+".html"))
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("test").Parse(content)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// FileContents returns the file content
func FileContents(file string) (string, error) {
	out, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
