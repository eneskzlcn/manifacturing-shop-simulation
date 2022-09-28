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

func New() *Simulation {
	return &Simulation{}
}

func (s *Simulation) init(properties Properties) error {
	if err := properties.Validate(); err != nil {
		return err
	}
	s.Properties = properties
	s.Time = 0
	s.InspectorAvailability = true
	s.Finished = false
	s.initFEL()
	s.initExamineQueue()
	return nil
}

func (s *Simulation) initExamineQueue() {
	s.ExamineQueue = priority_queue.NewPriorityQueue[EventData]()
	s.ExamineQueueStatistics = CumulativeStatistics{}
}

func (s *Simulation) initFEL() {
	s.FEL = priority_queue.NewPriorityQueue[EventData]()
	s.scheduleNextEventToArrive()
	s.FELStatistics = CumulativeStatistics{}
}

func (s *Simulation) Start(properties Properties) error {
	if err := s.init(properties); err != nil {
		return err
	}
	for !s.Finished {
		eventData := s.timeAdvanceFn()
		s.handleEvent(eventData)
	}
	report := s.GenerateReport()
	report.Print()
	return nil
}

func (s *Simulation) timeAdvanceFn() EventData {
	eventData := s.FEL.Dequeue().(EventData)
	s.Time = eventData.FinishTime
	return eventData
}

func (s *Simulation) handleEvent(eventData EventData) {
	switch eventData.Type {
	case ARRIVAL:
		if s.IsInspectorAvailable() {
			s.startExamine()
		} else {
			s.addEventToExamineQueue(eventData)
		}
		s.scheduleNextEventToArrive()
		break
	case EXAMINE:
		if s.isCurrentPartChosenAsFaulty() {
			s.handleFailPrediction()
		}
		if s.hasEventInExamineQueue() {
			s.scheduleNextEventInExamineQueueToExamine()
			break
		}
		s.MakeInspectorAvailable()
		break
	}
	s.prepareFELStatistics()
}

func (s *Simulation) IsInspectorAvailable() bool {
	return s.InspectorAvailability == true
}

func (s *Simulation) MakeInspectorAvailable() {
	s.InspectorAvailability = true
}

func (s *Simulation) MakeInspectorBusy() {
	s.InspectorAvailability = false
}

func (s *Simulation) scheduleNextEventInExamineQueueToExamine() {
	event := s.ExamineQueue.Dequeue().(EventData)
	event.Type = EXAMINE
	event.FinishTime = s.Time + randomutil.RandomInt(s.Properties.MinExamineTime, s.Properties.MaxExamineTime)
	event.StandbyDuration = event.FinishTime - event.ArrivalTime
	event.ArrivalTime = s.Time
	s.FEL.Enqueue(event)
}

func (s *Simulation) hasEventInExamineQueue() bool {
	return s.ExamineQueue.Length() > 0
}

func (s *Simulation) handleFailPrediction() {
	s.Properties.TerminateCounter--
	if s.Properties.TerminateCounter <= 0 {
		s.Finished = true
	}
}

func (s *Simulation) isCurrentPartChosenAsFaulty() bool {
	percentagePrediction := randomutil.RandomInt(0, 100)
	if percentagePrediction <= s.Properties.FailurePossibilityPercentage {
		return true
	}
	return false
}

func (s *Simulation) scheduleNextEventToArrive() {
	s.FEL.Enqueue(EventData{
		Type:        ARRIVAL,
		ArrivalTime: s.Time,
		FinishTime:  s.Time + s.Properties.PartTurnOutRate,
	})
}

func (s *Simulation) addEventToExamineQueue(eventData EventData) {
	s.ExamineQueue.Enqueue(eventData)
	s.prepareExamineStatistics(s.ExamineQueue.Length())
}

func (s *Simulation) startExamine() {
	s.FEL.Enqueue(EventData{
		Type:            EXAMINE,
		ArrivalTime:     s.Time,
		FinishTime:      s.Time + randomutil.RandomInt(s.Properties.MinExamineTime, s.Properties.MaxExamineTime),
		StandbyDuration: 0,
	})
	s.MakeInspectorBusy()
}

func (s *Simulation) prepareExamineStatistics(examineQueueLength int) {
	s.ExamineQueueStatistics.Prepare(examineQueueLength)
}

func (s *Simulation) prepareFELStatistics() {
	s.FELStatistics.Prepare(s.FEL.Length())
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
