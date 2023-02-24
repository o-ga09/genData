/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genDatCmd represents the genDat command
var genDatCmd = &cobra.Command{
	Use:   "genDat",
	Short: "generate dat file",
	Long: `generate dat file with yaml config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("genDat called")
	},
}

func init() {
	rootCmd.AddCommand(genDatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genDatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genDatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
