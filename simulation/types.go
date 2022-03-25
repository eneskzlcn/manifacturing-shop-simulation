package simulation

type CumulativeStatistics struct {
	MaxQueueLength int
	AvgQueueLength int
	QueueLengths []int
}
func (cs *CumulativeStatistics) Prepare(newQueueLength int) {
	cs.QueueLengths = append(cs.QueueLengths, newQueueLength)
	if cs.MaxQueueLength < newQueueLength {
		cs.MaxQueueLength = newQueueLength
	}
	cs.AvgQueueLength += newQueueLength
}
func (cs * CumulativeStatistics) GenerateReport() StatisticsReport{
	report := StatisticsReport{
		MaxQueueLength: cs.MaxQueueLength,
		AvgQueueLength: cs.AvgQueueLength / len(cs.QueueLengths),
	}
	return report
}
type ConditionalProperties struct {
	MinExamineTime int
	MaxExamineTime int
	TerminateCounter int
	FailurePossibilityPercentage int
}

type StatisticsReport struct {
	MaxQueueLength int
	AvgQueueLength int
}