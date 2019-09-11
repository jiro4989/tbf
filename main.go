package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

const (
	doc = `tbf は技術書典のサイトからサークル情報の一覧を取得します。

Usage:
	tbf [options] <tbf_id>
	tbf -h | --help
	tbf -v | --version

Examples:
	tbf tbf07
	tbf tbf07 -o tbf07.tsv

Options:
	-h --help             このヘルプを出力する。
	-v --version          バージョン情報を出力する。
	-o --outfile=<FILE>   ファイルに出力する。
	`
)

const (
	ErrorCodeOK ErrorCode = iota
	ErrorCodeFailedBinding
	ErrorCodeCreateFile
)

type (
	Config struct {
		TBFID   string `docopt:"<tbf_id>"`
		OutFile string `docopt:"--outfile"`
	}
	ErrorCode int
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ErrorCodeFailedBinding
	}

	var w *os.File
	if config.OutFile == "" {
		w = os.Stdout
	} else {
		var err error
		w, err = os.Create(config.OutFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return ErrorCodeCreateFile
		}
		defer w.Close()
	}

	tbfID := config.TBFID
	baseURL := fmt.Sprintf("https://techbookfest.org/api/circle?eventID=%s&visibility=site&limit=100&onlyAdoption=true", tbfID)
	for _, c := range fetchCircleOverviews(tbfID, baseURL) {
		text := fmt.Sprintf("%s\t%s\t%s\t%s", c.Name, c.NameRuby, c.URL, c.GenreFreeFormat)
		fmt.Fprintln(w, text)
	}

	return ErrorCodeOK
}
