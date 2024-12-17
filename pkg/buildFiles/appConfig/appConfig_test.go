package app_config_test

import (
	"os"
	"path/filepath"
	"testing"

	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
)

func Test_BuildConfigGatherPlugins(t *testing.T) {
	var vals = []struct {
		name       string
		cwd        string
		configPath string
	}{
		{name: "Should gather plugins from environment config", cwd: "/Users/bigsexy/Desktop/current/ulldCliTest", configPath: filepath.Join(os.Getenv("ULLD_ADDITIONAL_SOURCES"), "appConfig.ulld.json")},
	}
	for _, tt := range vals {
		t.Run(tt.name, func(t *testing.T) {
			a := app_config.NewAppConfig(tt.configPath)
			plugins := a.GatherPlugins()
			if len(plugins) <= 0 {
				t.Fail()
			}
			foundInSlots := false
			foundInPlugins := false
			for _, p := range plugins {
				if p.Name == "@ulld/pdf" {
					foundInSlots = true
				} else if p.Name == "@ulld/api" {
					foundInPlugins = true
				}
			}
			if !foundInPlugins {
				t.Log("Failed to gather plugins from plugins field")
				t.Fail()
			}
			if !foundInSlots {
				t.Log("Failed to gather plugins from slots field")
				t.Fail()
			}
		})
	}
}
