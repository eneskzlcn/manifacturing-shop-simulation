package simulation

import (
	"fmt"
	priority_queue "github.com/eneskzlcn/manifacturing-shop-simulation/internal/priority-queue"
	"log"
	"math/rand"
	"time"
)

type Simulation struct {
	Time                   int
	Finished               bool
	FELStatistics          CumulativeStatistics
	ExamineQueueStatistics CumulativeStatistics
	Conditions             ConditionalProperties
	FEL                    *priority_queue.PriorityQueue
	ExamineQueue           *priority_queue.PriorityQueue
	InspectorAvailability  bool
}

func NewSimulation() *Simulation {
	return &Simulation{}
}
func (s *Simulation) init(terminateCounter, minExamineTime, maxExamineTime, failurePossibilityPercentage int) {
	s.Conditions = ConditionalProperties{
		MinExamineTime:               minExamineTime,
		MaxExamineTime:               maxExamineTime,
		TerminateCounter:             terminateCounter,
		FailurePossibilityPercentage: failurePossibilityPercentage,
	}
	s.Time = 0
	s.InspectorAvailability = true
	s.Finished = false
	s.FEL = priority_queue.NewPriorityQueue()
	s.ExamineQueue = priority_queue.NewPriorityQueue()
	initialEvent := EventData{
		Type:            ARRIVAL,
		ArrivalTime:     0,
		FinishTime:      5,
		StandbyDuration: 0,
	}
	s.FEL.Enqueue(initialEvent)
	s.FELStatistics = CumulativeStatistics{}
	s.ExamineQueueStatistics = CumulativeStatistics{}
}
func (s *Simulation) Start(terminateCounter, minExamineTime, maxExamineTime, failurePossibilityPercentage int) {
	s.init(terminateCounter, minExamineTime, maxExamineTime, failurePossibilityPercentage)
	for !s.Finished {
		eventData := s.timeAdvanceFn()
		s.eventHandler(eventData)
	}
	s.GenerateReport()
}
func (s *Simulation) timeAdvanceFn() EventData {
	eventData := s.FEL.Dequeue().(EventData)
	log.Printf("\n Dequeue FEL: %s, %d, %d", eventData.Type.GetString(), eventData.ArrivalTime, eventData.FinishTime)
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
				FinishTime:      s.Time + Random(s.Conditions.MinExamineTime, s.Conditions.MaxExamineTime),
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
			FinishTime:  s.Time + 5,
		})
		break
	case EXAMINE:
		if failurePrediction := Random(0, 100); failurePrediction <= 10 {
			s.Conditions.TerminateCounter--
			if s.Conditions.TerminateCounter <= 0 {
				s.Finished = true
			}
		}
		if s.ExamineQueue.Length() != 0 {
			log.Println("Queuee is not empty")
			event := s.ExamineQueue.Dequeue().(EventData)
			event.Type = EXAMINE
			event.FinishTime = s.Time + Random(s.Conditions.MinExamineTime, s.Conditions.MaxExamineTime)
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
func (s *Simulation) GenerateReport() {
	felStatisticsReport := s.FELStatistics.GenerateReport()
	examineQueueStatisticsReport := s.ExamineQueueStatistics.GenerateReport()
	fmt.Println("---------- GENERATE REPORT BEGINS ----------")
	fmt.Printf("\nStatistics\n")
	fmt.Printf("Max FEL Length: %d, Average FEL Length: %d\n", felStatisticsReport.MaxQueueLength, felStatisticsReport.AvgQueueLength)
	fmt.Printf("Max Examine Queue Length: %d, Average Examine Queue Length: %d\n", examineQueueStatisticsReport.MaxQueueLength, examineQueueStatisticsReport.AvgQueueLength)

	fmt.Printf("\nCurrent FEL List\n")
	felItems := s.FEL.GetItems()
	for index, item := range felItems {
		eventData := item.(EventData)
		fmt.Printf("%d'th Event = Type: %s, Arrival Time: %d, Finish Time: %d, Standby Duration: %d\n",
			index, eventData.Type.GetString(), eventData.ArrivalTime, eventData.FinishTime, eventData.StandbyDuration)
	}
	fmt.Println("\n\n---------- GENERATE REPORT END ----------")
}

func Random(lowerBound, upperBound int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upperBound-lowerBound) + lowerBound
}
