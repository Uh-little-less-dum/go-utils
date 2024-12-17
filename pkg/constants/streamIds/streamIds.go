package stream_ids

type StreamId int

const (
	PreConflictResolve StreamId = iota
	PostConflictResolve
)
