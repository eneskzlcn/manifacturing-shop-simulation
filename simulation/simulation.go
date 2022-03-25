package simulation

import (
	"ManifacturingShopSimulation/priority-queue"
	"fmt"
	"math/rand"
	"time"
)

type Simulation struct {
	Time int
	Finished bool
	Statistics CumulativeStatistics
	Conditions ConditionalProperties
	FEL * priority_queue.PriorityQueue
}
func NewSimulation() *Simulation {
	return &Simulation{}
}
func (s *Simulation) init(terminateCounter, minExamineTime, maxExamineTime,failurePossibilityPercentage int) {
	s.Conditions = ConditionalProperties{
		MinExamineTime:   minExamineTime,
		MaxExamineTime:   maxExamineTime,
		TerminateCounter: terminateCounter,
		FailurePossibilityPercentage: failurePossibilityPercentage,
	}
	s.Time = 0
	s.Finished = false
	s.FEL = priority_queue.NewPriorityQueue()
	initialEvent := EventData{
		Type:            EXAMINE,
		ArrivalTime:     0,
		StartTime:       0,
		StandbyDuration: 0,
	}
	s.FEL.Enqueue(initialEvent)
	s.Statistics = CumulativeStatistics{}

}
func (s * Simulation) Start(terminateCounter, minExamineTime,maxExamineTime,failurePossibilityPercentage int) {
	s.init(terminateCounter,minExamineTime,maxExamineTime,failurePossibilityPercentage)
	for !s.Finished {
		eventData := s.timeAdvanceFn()
		s.eventHandler(eventData)
	}
	s.GenerateReport()
}
func (s *Simulation) timeAdvanceFn () EventData {
	eventData := s.FEL.Dequeue().(EventData)
	s.Time+=5
	return eventData
}
func (s * Simulation) eventHandler(eventData EventData) {
	var eventResult EventResult
	switch eventData.Type {
	case EXAMINE:
		eventResult = s.ExamineEvent(eventData)
		break
	}
	nextEventData := EventData{
		Type:            EXAMINE,
		ArrivalTime:     s.Time,
		StartTime:       eventResult.FinishTime,
		StandbyDuration:  eventResult.FinishTime - s.Time,
	}
	s.FEL.Enqueue(nextEventData)
	s.prepareStatistics(s.FEL.Length())

}
func (s * Simulation) prepareStatistics(newFelLength int) {
	s.Statistics.Prepare(newFelLength)
}
func (s * Simulation) GenerateReport() {
	statisticsReport := s.Statistics.GenerateReport()
	fmt.Println("---------- GENERATE REPORT BEGINS ----------")
	fmt.Printf("\nStatistics\n")
	fmt.Printf("Max FEL Length: %d, Average FEL Length: %d\n",statisticsReport.MaxQueueLength,statisticsReport.AvgQueueLength)
	fmt.Printf("\nCurrent FEL List\n")
	felItems := s.FEL.GetItems()
	for index,item := range felItems {
		eventData := item.(EventData)
		fmt.Printf("%d'th Event = Type: %d, Arrival Time: %d, Start Time: %d, Standby Duration: %d,",
			index,eventData.Type, eventData.ArrivalTime, eventData.StartTime, eventData.StandbyDuration)
	}
	fmt.Println("\n\n---------- GENERATE REPORT END ----------")
}

func (s * Simulation) ExamineEvent(eventData EventData) EventResult {
	examineDuration := Random(s.Conditions.MinExamineTime, s.Conditions.MaxExamineTime)
	failure := false
	if failurePrediction := Random(0, 100); failurePrediction <= 10 {
		failure = true
	}
	eventResult := EventResult{
		FinishTime:      eventData.StartTime + examineDuration,
		Failure:         failure,
		ExamineDuration: examineDuration,
	}
	if failure {
		s.Conditions.TerminateCounter -= 1
	}
	if s.Conditions.TerminateCounter <= 0 {
		s.Finished = true
	}
	return eventResult
}
func Random(lowerBound, upperBound int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upperBound -lowerBound) + lowerBound
}