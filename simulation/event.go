package simulation

type EventType int
const (
	EXAMINE EventType = iota
)
type EventData struct {
	Type EventType
	ArrivalTime int
	StartTime int
	StandbyDuration int
}
func (e EventData) GetPriority() int {
	return e.ArrivalTime
}
type EventResult struct {
	FinishTime int
	Failure bool
	ExamineDuration int
}
