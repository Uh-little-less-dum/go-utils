package build_config

import (
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	build_stages "github.com/Uh-little-less-dum/go-utils/pkg/constants/buildStages"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type BuildConfigOpts struct {
	TargetDir    string
	InitialStage build_stages.BuildStage
	AppConfig    *app_config.AppConfig
}

func (b *BuildConfigOpts) SetDataFromViper() {
	targetDir := viper.GetViper().GetString("targetDir")
	if targetDir != "" {
		b.TargetDir = targetDir
	}
}

// TODO: Set up the rest of the config flags here. Handle what you can with viper first though. Might not even need this struct if viper can handle everything.
func (c *BuildConfigOpts) Init(cmd *cobra.Command, targetDir string) {
	if len(targetDir) > 0 {
		c.TargetDir = targetDir
		// c.InitialStage = constants.CloneTemplateAppStage
	}
	c.InitialStage = build_stages.ConfirmCurrentDirStage
}
