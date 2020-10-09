package main

import (
	"fmt"
	"github.com/eden-framework/plugins"
	"path"
)

var Plugin GenerationPlugin

type GenerationPlugin struct {
}

func (g *GenerationPlugin) GenerateEntryPoint(opt plugins.Option, cwd string) string {
	globalPkgPath := path.Join(opt.PackageName, "internal/global")
	globalFilePath := path.Join(cwd, "internal/global")
	tpl := fmt.Sprintf(`,
		{{ .UseWithoutAlias "github.com/eden-framework/eden-framework/pkg/application" "" }}.WithConfig(&{{ .UseWithoutAlias "%s" "%s" }}.EventConfig)`, globalPkgPath, globalFilePath)
	return tpl
}

func (g *GenerationPlugin) GenerateFilePoint(opt plugins.Option, cwd string) []*plugins.FileTemplate {
	file := plugins.NewFileTemplate("global", path.Join(cwd, "internal/global/event.go"))
	file.WithBlock(`
var EventConfig = struct {
	Event *{{ .UseWithoutAlias "github.com/eden-framework/plugin-event/event" "" }}.MessageBus
}{
	Event: &{{ .UseWithoutAlias "github.com/eden-framework/plugin-event/event" "" }}.MessageBus{
		Driver: event.EVENT_DRIVER__BUILDIN,
	},
}
`)

	return []*plugins.FileTemplate{file}
}
