package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-excel-csv",
	Short: "コマンドに関する説明 short version",
	Long:  `コマンドに関する説明 long version`,

	// 実際の処理
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-excel-csv.yaml)")

	// フラグ設定
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 設定ファイルと ENV 変数を読み込む。設定ファイルと ENV 変数を読み込む。
func initConfig() {
	if cfgFile != "" {
		// コンフィグファイルのパス、名前、拡張子を明示的に定義します
		viper.SetConfigFile(cfgFile)
	} else {
		//$HOME環境変数を返します
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-cli-excel-csv" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-cli-excel-csv")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
