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

// genDatCmd represents the genDat command
var genDatCmd = &cobra.Command{
	Use:   "genDat",
	Short: "generate dat file",
	Long: `generate dat file with yaml config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		genDat(os.Args)
	},
}

func genDat(args []string) {
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
	defer  f.Close()

	for i := 0;i < config_data.Row_num;i++ {
		var record string
		for _,d := range config_data.Col_info {
			format_int := "%0" + strconv.Itoa(d.Str_len) + "d"
			format_str := "%" + strconv.Itoa(d.Str_len) + "s"
			if !d.Require {
				record += fmt.Sprintf(format_str," ")	
			} else {
				record += fmt.Sprintf(format_int,GenRandomNum(d.Str_len))
			}
		}
		record += config_data.Nc
		f.WriteString(record)
	}
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
