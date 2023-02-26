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
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Col_info struct {
	Col_name string `yaml:"col_name"`
	Str_len int `yaml:"str_len"`
	Require bool `yaml:"require"`
	Domain_name string `yaml:"domain_name"`
	Ratio float32 `yaml:"ratio"`
}

type Config struct {
	Row_num int `yaml:"row_num"`
	Col_info []Col_info `yaml:"Col_info"`
	Nc string `yaml:"nc"`
}

// genCsvCmd represents the genCsv command
var genCsvCmd = &cobra.Command{
	Use:   "genCsv",
	Short: "generate csv file",
	Long: `generate csv file with yaml config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		genCSV(os.Args)
	},
}

func genCSV(args []string) {
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

	w := csv.NewWriter(f)

	// csvに書き込むデータを作成する
	for i := 0; i < config_data.Row_num; i++ {
		var record []string
		for _,d := range config_data.Col_info {
			if i == 0 {
				record = append(record,d.Col_name)
				continue
			}
			format := "%0" + strconv.Itoa(d.Str_len) + "d"
			record = append(record,fmt.Sprintf(format,GenRandomNum(d.Str_len)))
		}
		w.Write(record)
	}
	w.Flush()

}

func ReadConfigfile(fb []byte) (Config,error) {
	var data Config
	err := yaml.Unmarshal(fb,&data)

	if err != nil {
		fmt.Println(err)
	}

	return data,nil
}

func init() {
	rootCmd.AddCommand(genCsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCsvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GenRandomNum(r int) int64 {
	var res int64

	rand.Seed(time.Now().UnixNano())
	res = int64(rand.Intn(r))

	return res
}
