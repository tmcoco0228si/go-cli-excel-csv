/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
)

// エクセルからCSVに変換するコマンド
var excelToCsvCmd = &cobra.Command{
	Use:   "excelToCsv",
	Short: "エクセルからCSVに変換するコマンドです。",
	Long: `エクセルからCSVに変換するコマンドです。
					複数行入力されたセルがある状態でも行読み込みして全てCSVとして書き込みを行います。					
				`,
	RunE: func(cmd *cobra.Command, args []string) error {

		filePath := "test.xlsx"
		f, err := excelize.OpenFile(filePath)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 拡張子（.xlsx）を取得する
		extension := filepath.Base(filePath)

		// .xlsx拡張子を取り除いたファイル名を取得する
		fileName := strings.TrimSuffix(extension, filepath.Ext(extension))

		// xlsxファイル名を元にCSVファイル作成
		csvFile, err := os.Create(fileName + ".csv")
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer csvFile.Close()

		csvWriter := csv.NewWriter(csvFile) // CSVライターの作成、
		defer csvWriter.Flush()             //内部バッファのフラッシュ

		rows := f.GetRows("Sheet1")

		// 1行ずつ読み取り、行に対する複数セル
		for _, row := range rows {
			var record []string
			for _, cellValue := range row {
				fmt.Println(cellValue)
				record = append(record, cellValue)
			}
			fmt.Println(record)
			err := csvWriter.Write(record)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(excelToCsvCmd)
}
