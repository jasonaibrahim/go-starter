package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetTemplateNameFromPath(pathToTemplate string, parentPath string) string {
	pathToTemplateParts := strings.Split(pathToTemplate, string(os.PathSeparator))
	containingPathParts := strings.Split(parentPath, string(os.PathSeparator))

	var newTemplateNameParts []string

	if len(pathToTemplateParts) > len(containingPathParts) {
		startingIndex := len(pathToTemplateParts) - (len(pathToTemplateParts) - len(containingPathParts))
		newTemplateNameParts = pathToTemplateParts[startingIndex:]
	} else {
		newTemplateNameParts = pathToTemplateParts
	}

	lastPath := getLastPath(newTemplateNameParts)
	if lastPath == "index" {
		newTemplateNameParts = newTemplateNameParts[:len(newTemplateNameParts)-1]
	}

	joined := strings.Join(newTemplateNameParts, string(os.PathSeparator))
	return strings.TrimSuffix(joined, filepath.Ext(joined))
}

func getLastPath(pathParts []string) string {
	last := pathParts[len(pathParts)-1]
	return strings.TrimSuffix(last, filepath.Ext(last))
}
