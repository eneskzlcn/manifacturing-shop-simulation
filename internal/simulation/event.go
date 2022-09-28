package simulation

import "fmt"

type EventType int

func (e EventType) GetString() string {
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
	Type            EventType
	ArrivalTime     int
	FinishTime      int
	StandbyDuration int
}

func (e EventData) Print(eventOrder int) {
	fmt.Printf("%d'th Event = Type: %s, Arrival Time: %d, Finish Time: %d, Standby Duration: %d\n",
		eventOrder, e.Type, e.ArrivalTime, e.FinishTime, e.StandbyDuration)
}

func (e EventData) GetPriority() int {
	return e.FinishTime
}
