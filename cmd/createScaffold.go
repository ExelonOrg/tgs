/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/davoodharun/terragrunt-scaffolder/helpers"
	"github.com/davoodharun/terragrunt-scaffolder/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createScaffoldCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigType("json")
		viper.SetConfigName("config")
		viper.AddConfigPath(".tgs")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		var config structs.Config
		err := viper.Unmarshal(&config)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		config = helpers.ReadConfig()
		// Capture any error

		// iterate through groups (or stacks/patters i.e. production, non_production)
		for group_key, group_value := range config.Stacks {

			// read stack (file must exist TODO: create check)
			var stack = helpers.ReadStack(group_key)

			// for each environment in group
			for environment_key := range group_value.Environments {

				// iterate over application types (base_modules) in stack yaml file
				for apptype_key, apptype_value := range stack {
					// for each application within module type
					for app_key := range apptype_value {
						// create folder structure and application hcl file
						appFile := structs.AppHclFile{
							Path:            fmt.Sprintf("%s/%s/%s/%s", group_key, group_value.Environments[environment_key], apptype_key, app_key),
							Environment:     group_value.Environments[environment_key],
							Group:           group_key,
							BaseModule:      apptype_key,
							App:             app_key,
							DependencyChain: apptype_value[app_key].Dependencies,
						}
						appFile.Write()
						if myfile, err := os.Create(fmt.Sprintf("%s/%s/%s/%s/README.md", group_key, group_value.Environments[environment_key], apptype_key, app_key)); err != nil {
							log.Fatal(err)
							myfile.Close()
						}
					}
				}

				environemntHclFile := structs.EnvironmentHclFile{
					Name: group_value.Environments[environment_key],
					Path: fmt.Sprintf("%s/%s", group_key, group_value.Environments[environment_key]),
				}
				environemntHclFile.Write()

				// if myfile, err := os.Create(fmt.Sprintf("%s/%s/%s.hcl", group_key, group_value.Environments[environment_key], group_value.Environments[environment_key])); err != nil {
				// 	log.Fatal(err)
				// 	myfile.Close()
				// }

			}
			groupHclFile := structs.GroupHclFile{
				Path: group_key,
			}

			groupHclFile.Write()
			globalHclFile := structs.GlobalHclFile{}
			globalHclFile.Write()

			for k := range stack {
				var str = fmt.Sprintf("_base_modules/%s", k)
				baseModuleHclFile := structs.BaseModuleHclFile{
					Name: k,
				}

				baseModuleHclFile.Write()
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

				if myfile, err := os.Create(fmt.Sprintf("%s/providers.tf", str)); err != nil {
					log.Fatal(err)
					myfile.Close()
				}

			}

		}
		// _ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return err
		// 	}
		// 	fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		// 	return nil
		// })
		// if err != nil {
		// 	fmt.Println(err)
		// }

	},
}

func init() {
	scaffoldCmd.AddCommand(createScaffoldCmd)

}
