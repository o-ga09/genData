# Usage for genData

- ## genCSV

概要

- 指定桁数の文字列をカンマ区切りでcsvファイルに出力する
- 改行文字を指定することが可能

  
実行形式 - how to use
  ```bash
    genData genCSV [config file] [output file name]
  ```
  
設定ファイルの書式 - config file format
  ```yaml
    row_num: 100 # データの行数 line of deta
    nc: "\n"     # 改行文字 newline character 
    Col_info:    # データの項目情報 data infomation
        - {
            col_name: "test1", #項目名1 column name 1
            str_len: 10,       #項目1の文字列長 length column name 1
          }
        - {
            col_name: "test2", #項目名2 column name 2
            str_len: 20,       #項目2の文字列長 length column name 2
          }
        - {
            col_name: "test3", #項目名3 column name 3
            str_len: 5,        #項目3の文字列長 length column name 3
          }
  ```

  - config file, output file ともにカレントディレクトリに配置、出力されます
  

- ## genDat

概要

- 指定桁数の文字列を複数結合した固定長のdatファイルに出力する
- 改行文字を指定することが可能

  実行形式 - how to use
  ```bash
    genData genDat [config file] [output file name]
  ```
  
  設定ファイルの書式 - config file format
  ```yaml
    row_num: 100 # データの行数 line of deta
    nc: "\n"     # 改行文字 newline character 
    Col_info:    # データの項目情報 data infomation
        - {
            col_name: "test1", #項目名1 column name 1
            str_len: 10,       #項目1の文字列長 length column name 1
            require: false     #項目名1が必要化否かのフラグ flag of required one
          }
        - {
            col_name: "test2", #項目名2 column name 2
            str_len: 20,       #項目2の文字列長 length column name 2
            require: false     #項目名2が必要化否かのフラグ flag of required one
          }
        - {
            col_name: "test3", #項目名3 column name 3
            str_len: 5,        #項目3の文字列長 length column name 3
            require: false     #項目名3が必要化否かのフラグ flag of required one
          }
  ```

  - config file, output file ともにカレントディレクトリに配置、出力されます

- ## genAddress

概要

- 指定したドメインのダミーメールアドレスを出現割合を指定してtxtファイルに出力する
- 改行文字を指定することが可能

  実行形式 - how to use
  ```bash
    genData genAddress [config file] [output file name]
  ```
  
  設定ファイルの書式 - config file format
  ```yaml
    row_num: 100 # データの行数 line of deta
    nc: "\n"     # 改行文字 newline character 
    Col_info:    # データの項目情報 data infomation
        - {
            domain_name: "@icloud.com", #ドメイン名1 domain name1
            ratio: 0.5,               #ドメイン名1の出現割合 ratio of domain name1
          }
        - {
            domain_name: "@ezweb.ne.jp", #ドメイン名1 domain name2
            ratio: 0.25,               #ドメイン名2の出現割合 ratio of domain name2
          }
        - {
            domain_name: "@i.softbank.jp", #ドメイン名1 domain name3
            ratio: 0.25,               #ドメイン名3の出現割合 ratio of domain name3
          }
  ```

- config file, output file ともにカレントディレクトリに配置、出力されます
- ドメインの出現割合は合計で1となるように設定してください
  

copyright 2023 @o-ga09