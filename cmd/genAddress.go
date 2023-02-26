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
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// genAddressCmd represents the genAddress command
var genAddressCmd = &cobra.Command{
	Use:   "genAddress",
	Short: "generate dammy mail address",
	Long: `generate dammy mail address. file foemat csv.`,
	Run: func(cmd *cobra.Command, args []string) {
		genAddress(os.Args)
	},
}

func genAddress(args []string) {
	//引数チェック
	if len(os.Args) < 4 {
		fmt.Println("need to config file")
		os.Exit(1)
	}
	//コンフィグファイルを開く
	in_files := args[2]
	buf ,_:= ioutil.ReadFile(in_files)
	config_data, err := ReadConfigfile(buf)
	if err != nil {
		os.Exit(1)
	}
	//データを書き込むファイルを作成する
	out_file := args[3]

	f,err := os.Create(out_file)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	for _,d := range config_data.Col_info {
		for i := 0;i < int(float32(config_data.Row_num) * d.Ratio);i++ {
			address := fmt.Sprintf("%020s",strconv.Itoa(i)) + d.Domain_name + config_data.Nc
			f.WriteString(address)
		}
	}
}

func init() {
	rootCmd.AddCommand(genAddressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genAddressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genAddressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
