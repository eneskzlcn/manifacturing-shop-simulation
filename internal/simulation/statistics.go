package simulation

type CumulativeStatistics struct {
	MaxQueueLength int
	AvgQueueLength int
	QueueLengths   []int
}

func (cs *CumulativeStatistics) Prepare(newQueueLength int) {
	cs.QueueLengths = append(cs.QueueLengths, newQueueLength)
	if cs.MaxQueueLength < newQueueLength {
		cs.MaxQueueLength = newQueueLength
	}
	cs.AvgQueueLength += newQueueLength
}
func (cs *CumulativeStatistics) GenerateReport() StatisticsReport {
	report := StatisticsReport{
		MaxQueueLength: cs.MaxQueueLength,
	}
	if len(cs.QueueLengths) <= 0 {
		report.AvgQueueLength = 0
	} else {
		report.AvgQueueLength = cs.AvgQueueLength / len(cs.QueueLengths)
	}
	return report
}
