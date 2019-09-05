package test

import (
	"github.com/jasonaibrahim/go-starter/src/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplateNameFromDirectory(t *testing.T) {
	templatesDir := "/go/src/github.com/foo/go-starter/src/app/views"
	templatePath := "/go/src/github.com/foo/go-starter/src/app/views/foo/bar.html"
	templateName := util.GetTemplateNameFromPath(templatePath, templatesDir)
	assert.Equal(t, "foo/bar", templateName)
}

func TestArbitraryFileExtensions(t *testing.T) {
	templatesDir := "/go/src/github.com/foo/go-starter/src/app/views"
	templatePath := "/go/src/github.com/foo/go-starter/src/app/views/foo/bar.xyzasdbfb"
	templateName := util.GetTemplateNameFromPath(templatePath, templatesDir)
	assert.Equal(t, "foo/bar", templateName)
}

func TestTemplateNameFromDirectoryIgnoringIndex(t *testing.T) {
	templatesDir := "/go/src/github.com/foo/go-starter/src/app/views"
	templatePath := "/go/src/github.com/foo/go-starter/src/app/views/home/index.html"
	templateName := util.GetTemplateNameFromPath(templatePath, templatesDir)
	assert.Equal(t, "home", templateName)

	templatePath = "/go/src/github.com/foo/go-starter/src/app/views/home/index.tmpl"
	templateName = util.GetTemplateNameFromPath(templatePath, templatesDir)
	assert.Equal(t, "home", templateName)
}
