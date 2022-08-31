/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/davoodharun/terragrunt-scaffolder/helpers"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// viper.SetConfigType("json")
		// viper.SetConfigName("config")
		// viper.AddConfigPath(".tgs")
		// if err := viper.ReadInConfig(); err != nil {
		// 	log.Fatalf("Error reading config file, %s", err)
		// }
		// var config structs.Config
		// err := viper.Unmarshal(&config)
		// if err != nil {
		// 	log.Fatalf("unable to decode into struct, %v", err)
		// }

		var config = helpers.ReadConfig()

		for group_key, v := range config.Groups {
			var group []string
			for environment_key, environment_value := range v {
				for apptype_key, apptype_value := range environment_value {
					for app_key := range apptype_value {
						group = append(group, fmt.Sprintf("%s/%s/%s/%s", group_key, environment_key, apptype_key, app_key))

					}
				}
			}
			for i := 0; i < len(group); i++ {
				if err := os.MkdirAll(group[i], os.ModePerm); err != nil {
					log.Fatal(err)
				}

				if myfile, err := os.Create(fmt.Sprintf("%s/terragrunt.hcl", group[i])); err != nil {
					log.Fatal(err)
					myfile.Close()
				}
			}

			if myfile, err := os.Create(fmt.Sprintf("%s/%s.hcl", group_key, group_key)); err != nil {
				log.Fatal(err)
				myfile.Close()
			}
		}

		for i := 0; i < len(config.BaseModules); i++ {
			var str = fmt.Sprintf("_base_modules/%s", config.BaseModules[i])
			if err := os.MkdirAll(str, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			if myfile, err := os.Create(fmt.Sprintf("%s/main.tf", str)); err != nil {
				log.Fatal(err)
				myfile.Close()
			}

			if myfile, err := os.Create(fmt.Sprintf("%s/outputs.tf", str)); err != nil {
				log.Fatal(err)
				myfile.Close()
			}

			if myfile, err := os.Create(fmt.Sprintf("%s/variables.tf", str)); err != nil {
				log.Fatal(err)
				myfile.Close()
			}

			if myfile, err := os.Create(fmt.Sprintf("%s/%s.hcl", str, config.BaseModules[i])); err != nil {
				log.Fatal(err)
				myfile.Close()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
