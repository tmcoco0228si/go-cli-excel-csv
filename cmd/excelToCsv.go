/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
)

// excelToCsvCmd represents the excelToCsv command
var excelToCsvCmd = &cobra.Command{
	Use:   "excelToCsv",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("excelToCsv called")

		filePath := "サンプル.xlsx"
		// excelファイル読み込み
		f, err := excelize.OpenFile(filePath)
		if err != nil {
			return err
		}
		fmt.Println(f.GetCellValue("Sheet1", "A1"))

		// ファイル名のみを取得する
		filename := filepath.Base(filePath)

		// .xlsx拡張子を取り除いたファイル名を取得する
		newFileName := strings.TrimSuffix(filename, filepath.Ext(filename))

		// xlsxファイル名を元にCSVファイル作成
		csvFile, err := os.Create(newFileName + ".csv")
		if err != nil {
			return err
		}
		// csvFile.Close()

		// CSV書き込み
		csvWriter := csv.NewWriter(csvFile)
		//内部バッファのフラッシュ
		defer csvWriter.Flush()

		csvWriter.Write([]string{f.GetCellValue("Sheet1", "A1")})

		r := csv.NewReader(csvFile)
		rows, err := r.ReadAll() // csvを一度に全て読み込む
		if err != nil {
			log.Fatal(err)
		}

		// [][]stringなのでループする
		for _, v := range rows {
			fmt.Println(v)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(excelToCsvCmd)
	// excelToCsvCmd.PersistentFlags().String("foo", "", "A help for foo")
	// excelToCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
