package cmd

import (
	"fmt"

	"github.com/silfoxs/silgo/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	silgoApiCmd = &cobra.Command{
		Use:   "silgo-api",
		Short: "A silgo-api project",
		Long:  `A silgo-api project`,
		Run: func(cmd *cobra.Command, args []string) {
			app.Run()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./configs/dev.yaml", "config file (default is $HOME/prod.yaml)")
	rootCmd.AddCommand(silgoApiCmd)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("config file:", viper.ConfigFileUsed())
		fmt.Println("env:", viper.GetString("app.env"))
		fmt.Println("port:", viper.GetString("app.port"))
	}
}
