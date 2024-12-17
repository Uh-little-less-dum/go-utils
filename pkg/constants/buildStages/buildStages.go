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
	SelectPackageManager
	CloneTemplateAppStage
	PreConflictResolveBuild
	PostConflictResolveBuild
	ResolvePluginConflicts
	End_BuildSuccess
	End_BuildFail
)
