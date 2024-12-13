package run_status

type RunStatus int

const (
	NotStarted RunStatus = iota
	Running
	Complete
)

func HasNotRun(currentStatus RunStatus) bool {
	return currentStatus == NotStarted
}
