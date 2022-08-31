/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/davoodharun/terragrunt-scaffolder/cmd"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	cmd.Execute()

}
