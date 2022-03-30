package manifacturing_shop_simulation

type EventType int

func (e EventType) GetString() string{
	switch e {
	default:
		return ""
	case ARRIVAL:
		return "ARRIVAL"
	case EXAMINE:
		return "EXAMINE"
	}
}
const (
	EXAMINE EventType = iota
	ARRIVAL
)

type EventData struct {
	Type        EventType
	ArrivalTime int
	FinishTime int
	StandbyDuration int
}
func (e EventData) GetPriority() int {
	return e.FinishTime
}
