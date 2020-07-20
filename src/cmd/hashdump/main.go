package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ryokky59/td-kadai/src/convert"
)

const (
	// ExitCodeSuccess 処理が成功したときのコード
	ExitCodeSuccess = 0
	// ExitCodeError 何らかの原因でエラーが返ってきたときのコード
	ExitCodeError = 1
	// ExitCodeInvalidFileError 指定したファイルが無効のときのコード
	ExitCodeInvalidFileError = 2
	// ExitCodeInvalidAlgorithmError 指定したハッシュアルゴリズムが無効のときのコード
	ExitCodeInvalidAlgorithmError = 3
)

var (
	input     string
	output    string
	algorithm string
)

func init() {
	flag.StringVar(&input, "input", "../../testdata/in.txt", "入力")
	flag.StringVar(&input, "i", "../../testdata/in.txt", "入力 (short)")
	flag.StringVar(&output, "output", "../../testdata/out.txt", "出力")
	flag.StringVar(&output, "o", "../../testdata/out.txt", "出力 (short)")
	flag.StringVar(&algorithm, "algorithm", "sha256", "アルゴリズム")
	flag.StringVar(&algorithm, "a", "sha256", "アルゴリズム (short)")
	flag.Parse()
}

func run() int {
	if input == "" {
		fmt.Fprintln(os.Stderr, "Input target file.\n  ex) `--input=./sha256.txt`")
		return ExitCodeInvalidFileError
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Target input file is not found.")
		return ExitCodeInvalidFileError
	}

	converter := convert.NewAlgorithmConverter(algorithm)
	if converter == nil {
		fmt.Fprintln(os.Stderr, "Target algorithm is not supported.")
		return ExitCodeInvalidAlgorithmError
	}

	if err := converter.Exec(input, output); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeError
	}

	return ExitCodeSuccess
}

func main() {
	os.Exit(run())
}
