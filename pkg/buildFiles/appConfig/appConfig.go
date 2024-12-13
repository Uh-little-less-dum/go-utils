package app_config

import (
	"fmt"
	"os"
	"path/filepath"

	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/classesKinda/plugin"
	schemas_app_config "github.com/Uh-little-less-dum/go-utils/pkg/schemastructs/ulldAppConfig"
	"github.com/charmbracelet/log"
)

type AppConfig struct {
	FilePath string
	Data     schemas_app_config.AppConfig
}

func NewAppConfig(targetDir string) *AppConfig {
	filePath := filepath.Join(targetDir, "appConfig.ulld.json")
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	j, err := schemas_app_config.UnmarshalAppConfig(b)
	if err != nil {
		log.Fatal(err)
	}
	return &AppConfig{
		FilePath: filePath,
		Data:     j,
	}
}

func (a *AppConfig) GatherPlugins() []ulld_plugin.Plugin {
	var data []ulld_plugin.Plugin
	for _, item := range a.Data.Plugins.UnionArray {
		data = append(data, ulld_plugin.Plugin{PluginName: item.PluginsClass.Name})
	}
	fmt.Print(a.Data)
	// Still need to gather plugins from the slots field.
	// for _, k := range a.Data.Slots {
	// if k
	// }
	return data
}
