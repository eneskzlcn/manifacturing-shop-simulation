package simulation

import "fmt"

type StatisticsReport struct {
	MaxQueueLength int
	AvgQueueLength int
}

func (s StatisticsReport) Print(statisticName string) {
	fmt.Printf("Max %s Length: %d, Average %s Length: %d\n", statisticName, s.MaxQueueLength, statisticName, s.AvgQueueLength)
}

type Report struct {
	FELStatisticReport          StatisticsReport
	ExamineQueueStatisticReport StatisticsReport
	FELEventData                []EventData
}

func (r Report) Print() {
	r.printHeader()
	r.printStatisticsSection()
	r.printFELEventDataSection()
	r.printFooter()
}
func (r Report) printHeader() {
	fmt.Println("---------- GENERATE REPORT BEGINS ----------")
}
func (r Report) printFooter() {
	fmt.Println("\n\n---------- GENERATE REPORT END ----------")
}
func (r Report) printFELEventDataSection() {
	fmt.Printf("\nCurrent FEL List\n")
	for eventOrder, eventData := range r.FELEventData {
		eventData.Print(eventOrder)
	}
}
func (r Report) printStatisticsSection() {
	fmt.Printf("\nStatistics\n")
	r.ExamineQueueStatisticReport.Print("Examine Queue")
	r.FELStatisticReport.Print("FEL")
}
