package build_stages

type BuildStage int

const (
	ConfirmCurrentDirStage BuildStage = iota
	PickTargetDirStage
	ConfirmConfigLocFromEnv
	PickConfigLoc
	ConfirmWaitForConfigMove
	WaitForConfigMove
	ChooseWaitOrPickConfigLoc
	CloneTemplateAppStage
	PreConflictResolveBuild
	ResolvePluginConflicts
	End_BuildSuccess
	End_BuildFail
)
