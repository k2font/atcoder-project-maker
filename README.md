# AtCoder Project Maker
## Features
- GoでAtCoderの問題を解く際に素早くプロジェクトを立ち上げるためのCLIアプリケーション
- C++もしくはGo用のプロジェクトを作成できます
- また、プロジェクト全体のほかに、ファイル単体(=1つの問題用のファイル)を作成することもできます

## Requirements
- spf13/cobra@1.7.0

## Usage
```bash
$ go build && ./mkatcproj setup --help                           
Setup a project or a file

Usage:
  mkatcproj setup [flags]

Aliases:
  setup, s

Flags:
  -h, --help          help for setup
  -l, --lang string   Languages you want to use
  -n, --name string   Name of project or file
  -t, --type string   Objects you want to make
```

## Example
- `$ mkatcproj setup --lang go --type project --name abc301`
  - Go用のAtCoder用プロジェクトを作成します
  ```bash
  TODO: コマンドの出力結果を記載
  ```

- `$ mkatcproj setup --lang cpp --type file --name abc301_d`
  - C++用のAtCoder用ファイルを作成します
  ```bash
  TODO: コマンドの出力結果を記載
  ```