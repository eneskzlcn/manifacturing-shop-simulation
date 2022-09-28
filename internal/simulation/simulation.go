package simulation

import (
	priority_queue "github.com/eneskzlcn/manifacturing-shop-simulation/internal/priority-queue"
	"github.com/eneskzlcn/manifacturing-shop-simulation/internal/util/convertutil"
	"github.com/eneskzlcn/manifacturing-shop-simulation/internal/util/randomutil"
)

type Simulation struct {
	Time                   int
	Finished               bool
	FELStatistics          CumulativeStatistics
	ExamineQueueStatistics CumulativeStatistics
	Properties             Properties
	FEL                    *priority_queue.PriorityQueue[EventData]
	ExamineQueue           *priority_queue.PriorityQueue[EventData]
	InspectorAvailability  bool
}

func NewSimulation() *Simulation {
	return &Simulation{}
}
func (s *Simulation) init(properties Properties) {
	s.Properties = properties
	s.Time = 0
	s.InspectorAvailability = true
	s.Finished = false
	s.FEL = priority_queue.NewPriorityQueue[EventData]()
	s.ExamineQueue = priority_queue.NewPriorityQueue[EventData]()
	initialEvent := EventData{
		Type:            ARRIVAL,
		ArrivalTime:     0,
		FinishTime:      s.Properties.PartTurnOutRate,
		StandbyDuration: 0,
	}
	s.FEL.Enqueue(initialEvent)
	s.FELStatistics = CumulativeStatistics{}
	s.ExamineQueueStatistics = CumulativeStatistics{}
}
func (s *Simulation) Start(properties Properties) {
	s.init(properties)
	for !s.Finished {
		eventData := s.timeAdvanceFn()
		s.eventHandler(eventData)
	}
	report := s.GenerateReport()
	report.Print()
}
func (s *Simulation) timeAdvanceFn() EventData {
	eventData := s.FEL.Dequeue().(EventData)
	s.Time = eventData.FinishTime
	return eventData
}
func (s *Simulation) eventHandler(eventData EventData) {
	switch eventData.Type {
	case ARRIVAL:
		if s.InspectorAvailability == true {
			s.FEL.Enqueue(EventData{
				Type:            EXAMINE,
				ArrivalTime:     s.Time,
				FinishTime:      s.Time + randomutil.RandomInt(s.Properties.MinExamineTime, s.Properties.MaxExamineTime),
				StandbyDuration: 0,
			})
			s.InspectorAvailability = false
		} else {
			s.ExamineQueue.Enqueue(eventData)
			s.prepareExamineStatistics(s.ExamineQueue.Length())
		}
		s.FEL.Enqueue(EventData{
			Type:        ARRIVAL,
			ArrivalTime: s.Time,
			FinishTime:  s.Time + s.Properties.PartTurnOutRate,
		})
		break
	case EXAMINE:
		if failurePrediction := randomutil.RandomInt(0, 100); failurePrediction <= s.Properties.FailurePossibilityPercentage {
			s.Properties.TerminateCounter--
			if s.Properties.TerminateCounter <= 0 {
				s.Finished = true
			}
		}
		if s.ExamineQueue.Length() != 0 {
			event := s.ExamineQueue.Dequeue().(EventData)
			event.Type = EXAMINE
			event.FinishTime = s.Time + randomutil.RandomInt(s.Properties.MinExamineTime, s.Properties.MaxExamineTime)
			event.StandbyDuration = event.FinishTime - event.ArrivalTime
			event.ArrivalTime = s.Time
			s.FEL.Enqueue(event)
			break
		}
		s.InspectorAvailability = true
		break
	}
	s.prepareFELStatistics(s.FEL.Length())
}
func (s *Simulation) prepareExamineStatistics(examineQueueLength int) {
	s.ExamineQueueStatistics.Prepare(examineQueueLength)
}
func (s *Simulation) prepareFELStatistics(newFelLength int) {
	s.FELStatistics.Prepare(newFelLength)
}
func (s *Simulation) GenerateReport() Report {
	felStatisticsReport := s.FELStatistics.GenerateReport()
	examineQueueStatisticsReport := s.ExamineQueueStatistics.GenerateReport()

	felEventData, err := convertutil.AnyTo[[]EventData](s.FEL.GetItems())
	if err != nil {
		return Report{}
	}

	return Report{
		FELStatisticReport:          felStatisticsReport,
		ExamineQueueStatisticReport: examineQueueStatisticsReport,
		FELEventData:                felEventData,
	}
}
