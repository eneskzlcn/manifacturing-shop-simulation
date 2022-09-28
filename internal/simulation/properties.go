package simulation

import (
	"errors"
	"fmt"
)

const (
	PartTurnOutRateLowerBound = 5
	PartTurnOutRateUpperBound = 10

	MaximumExamineTimeLowerBound = 8
	MaximumExamineTimeUpperBound = 4

	MinimumExamineTimeLowerBound = 1
	//minimumExamineTimeUpperBound depends on maximumExamineTime value dynamically(must be lower than it.)

	FaultyInspectionPercentageLowerBound = 10
	FaultyInspectionPercentageUpperBound = 90

	FaultyPartCountToTerminateLowerBound = 100
	FaultyPartCountToTerminateUpperBound = 500
)

var (
	InvalidPartTurnOutRate = errors.New(fmt.Sprintf("invalid part out rate. must be in range [%d, %d]",
		PartTurnOutRateLowerBound, PartTurnOutRateUpperBound))
	InvalidMinimumExamineTime = errors.New(
		fmt.Sprintf("must be greater or equal than %d and lower than Maximum Examine Time",
			MinimumExamineTimeLowerBound))
	InvalidMaximumExamineTime = errors.New(fmt.Sprintf("invalid maximum examine time. must be in range [%d, %d]",
		MaximumExamineTimeLowerBound, MaximumExamineTimeUpperBound))
	InvalidFaultyInspectionPercentage = errors.New(fmt.Sprintf("invalid faulty inspection percentage. must be in range [%d, %d]",
		FaultyInspectionPercentageLowerBound, FaultyInspectionPercentageUpperBound))
	InvalidFaultyPartCountToTerminate = errors.New(
		fmt.Sprintf("invalid faulty part count to terminate. must be in range [%d, %d]",
			FaultyPartCountToTerminateLowerBound, FaultyPartCountToTerminateUpperBound))
)

type Properties struct {
	MinExamineTime               int
	MaxExamineTime               int
	TerminateCounter             int
	FailurePossibilityPercentage int
	PartTurnOutRate              int
}

func (cp Properties) Validate() error {
	if !isValidPartTurnOutRate(cp.PartTurnOutRate) {
		return InvalidPartTurnOutRate
	}
	if !isValidMinimumExamineTime(cp.MinExamineTime, cp.MaxExamineTime) {
		return InvalidMinimumExamineTime
	}
	if !isValidMaximumExamineTime(cp.MaxExamineTime) {
		return InvalidMaximumExamineTime
	}
	if !isValidTerminateCounter(cp.TerminateCounter) {
		return InvalidFaultyPartCountToTerminate
	}
	if !isValidFailurePossibilityPercentage(cp.FailurePossibilityPercentage) {
		return InvalidFaultyInspectionPercentage
	}
	return nil
}
func isValidPartTurnOutRate(partTurnOutRate int) bool {
	return partTurnOutRate <= PartTurnOutRateUpperBound && partTurnOutRate >= PartTurnOutRateLowerBound
}
func isValidMinimumExamineTime(minimumExamineTime, maximumExamineTime int) bool {
	return minimumExamineTime < maximumExamineTime && minimumExamineTime >= MinimumExamineTimeLowerBound
}
func isValidMaximumExamineTime(maximumExamineTime int) bool {
	return maximumExamineTime <= MaximumExamineTimeUpperBound && maximumExamineTime >= MaximumExamineTimeLowerBound
}
func isValidTerminateCounter(terminateCounter int) bool {
	return terminateCounter <= FaultyPartCountToTerminateUpperBound &&
		terminateCounter >= FaultyPartCountToTerminateLowerBound
}
func isValidFailurePossibilityPercentage(failurePossibilityPercentage int) bool {
	return failurePossibilityPercentage >= FaultyInspectionPercentageLowerBound &&
		failurePossibilityPercentage <= FaultyInspectionPercentageUpperBound
}
