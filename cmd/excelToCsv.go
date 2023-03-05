/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// excelToCsvCmd represents the excelToCsv command
var excelToCsvCmd = &cobra.Command{
	Use:   "excelToCsv",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("excelToCsv called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(excelToCsvCmd)
	// excelToCsvCmd.PersistentFlags().String("foo", "", "A help for foo")
	// excelToCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
