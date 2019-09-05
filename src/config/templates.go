package config

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jasonaibrahim/go-starter/src/util"
	"os"
	"path/filepath"
)

func LoadHtmlTemplates(router *gin.Engine) error {
	cwd, _ := os.Getwd()
	templatesPath := filepath.Join(cwd, "src", "app", "views")

	render, err := loadTemplates(templatesPath)
	if err != nil {
		return err
	}

	router.HTMLRender = render

	return nil
}

func loadTemplates(templatesDir string) (multitemplate.Renderer, error) {
	renderer := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		return nil, err
	}

	views, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		return nil, err
	}

	for _, view := range views {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, view)
		templateName := util.GetTemplateNameFromPath(view, templatesDir)
		renderer.AddFromFiles(templateName, files...)
	}

	return renderer, nil
}
