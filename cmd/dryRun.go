/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dryRunCmd represents the dryRun command
var dryRunCmd = &cobra.Command{
	Use:   "dryRun",
	Short: "Generate config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dryRun called")
	},
}

func init() {
	rootCmd.AddCommand(dryRunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dryRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dryRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
