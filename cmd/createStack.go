/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/davoodharun/terragrunt-scaffolder/structs"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createStackCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		// if err := os.Mkdir(".tgs", os.ModePerm); err != nil {
		// 	log.Fatal(err)
		// }
		viper.SetConfigType("yaml")
		viper.SetConfigName(args[0])
		initialConfig := []byte(`
        functionapp:
            app1:
                dependencies:
                - servicebus.east
            app2:
                dependencies:
                - servicebus.east
        servicebus:
            east:
                dependencies: []
        `)

		data := structs.Pattern{}

		if err := yaml.Unmarshal(initialConfig, &data); err != nil {
			log.Fatal(err)
		}

		// for k, v := range data {
		// 	fmt.Printf(k, v)
		// }
		viper.SetDefault("functionapp", data["functionapp"])
		viper.SetDefault("servicebus", data["servicebus"])

		viper.AddConfigPath(".tgs")
		viper.SafeWriteConfig()
	},
}

func init() {
	stackCmd.AddCommand(createStackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stackCmd.PersistentFlags().String("name", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
