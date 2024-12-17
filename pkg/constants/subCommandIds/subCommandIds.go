package sub_command_ids

type SubCommandId int

const (
	GatherAppConfig SubCommandId = iota
	GatherPlugins
	GatherRootPackageJson
	InstallDependencies
)
