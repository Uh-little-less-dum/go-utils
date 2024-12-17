package app_config

import (
	"os"

	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/classesKinda/plugin"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type AppConfig struct {
	filePath string
	hasRead  bool
	data     gjson.Result
}

func NewAppConfig(filePath string) *AppConfig {
	return &AppConfig{
		filePath: filePath,
		hasRead:  false,
	}
}

func (a *AppConfig) Json() gjson.Result {
	if a.hasRead {
		return a.data
	}
	b, err := os.ReadFile(a.filePath)
	if err != nil {
		log.Fatal(err)
	}
	j := gjson.ParseBytes(b)
	a.data = j
	return j
}

func (a *AppConfig) GatherPlugins() []*ulld_plugin.Plugin {
	var res []*ulld_plugin.Plugin
	j := a.Json()
	plugins := j.Get("plugins").Array()
	for _, p := range plugins {
		res = append(res, ulld_plugin.PluginJsonToStruct(p))
	}
	j.Get("slots").ForEach(func(key, value gjson.Result) bool {
		res = append(res, ulld_plugin.SlotJsonToStruct(key, value)...)
		return true
	})
	return res
}
