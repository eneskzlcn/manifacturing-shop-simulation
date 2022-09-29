/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cobra

import (
	"fmt"
	"github.com/eneskzlcn/manufacturing-shop-simulation/internal/simulation"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "simulation",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "manufacturing shop simulation",
	Long: `
A machine tool in a manufacturing shop is turning out parts at the 
rate of every x minutes. As they are finished, the parts are sent to an 
inspector who takes at most y, at least z minutes (uniform distribution) to examine each 
one and rejects about d% of the parts as faulty. You are asked to run 
the simulation for m faulty parts to leave the system. Gather the 
statistics about average and maximum queue lengths.

Constraints: 
5 <= x <= 10, 
4<= y <= 8, 
1< =z < y,
20 <= d <= 90, 
100 <= m <= 500`,
	RunE: func(cmd *cobra.Command, args []string) error {
		manufacturingShopSimulation := simulation.New()
		err := manufacturingShopSimulation.Start(simulation.Properties{
			MinExamineTime:               minimumExamineTime,
			MaxExamineTime:               maximumExamineTime,
			TerminateCounter:             faultyPartCountToTerminate,
			FailurePossibilityPercentage: faultyInspectionPercentage,
			PartTurnOutRate:              partRateValue,
		})
		return err
	},
	Example: "simulation -x=3, -y=5, -z= 2, -d=20, -m=100",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var partRateValue int = 5
var maximumExamineTime int = 4
var minimumExamineTime int = 1
var faultyInspectionPercentage int = 10
var faultyPartCountToTerminate int = 100

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVarP(&partRateValue, "partRate", "x", 5, "Part turning out rate. Constraint: 5 <= rate <= 10")
	rootCmd.Flags().IntVarP(&maximumExamineTime, "maxExTime", "y", 4, "Maximum examine time of a part. Constraint: 4 <= maxExamineTime <= 8")
	rootCmd.Flags().IntVarP(&minimumExamineTime, "minExTime", "z", 1, "Minimum examine time of a part. Constraint: 1 <= minExamineTime <= maxExamineTime")
	rootCmd.Flags().IntVarP(&faultyInspectionPercentage, "faultyInspectionRate", "d", 10, "Percentage to examine a part as faulty. Constraint 10 <= faultyInspectionPercentage <= 90")
	rootCmd.Flags().IntVarP(&faultyPartCountToTerminate, "maxFaultyPartCount", "m", 100, "Maximum count of faulty part to simulation get terminated. Constraint 100 <= faultyInspectionPercentage <= 400")

}
