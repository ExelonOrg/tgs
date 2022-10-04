/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/davoodharun/terragrunt-scaffolder/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var stack Stack

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Mkdir(".tgs", os.ModePerm); err != nil {
			log.Fatal(err)
		}
		viper.SetConfigType("json")
		viper.SetConfigName("config")
		initialConfig := []byte(`{
			"stacks": {
			  "non_production": {
				"Environments": ["dev", "test", "stage"],
				"StateStorageAccountName": "eutfstatestorage",
				"StateStorageRG": "XZC-E-N-EUWEB00-S-RGP-10",
				"ContainerName": "example",
				"KeyPrefix": "example"
			  }
			}
		  }`)

		data := structs.Config{}

		_ = json.Unmarshal([]byte(initialConfig), &data)
		// for k, v := range data.Groups {
		// 	fmt.Printf(k, v)
		// }
		viper.SetDefault("stacks", data.Stacks)
		viper.AddConfigPath(".tgs")
		viper.SafeWriteConfig()

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
