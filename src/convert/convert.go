package convert

import (
	"bufio"
	"os"
	"sync"

	"github.com/ryokky59/td-kadai/src/domain"
)

// ConverterInterface ハッシュアルゴリズム変換用のインターフェース
type ConverterInterface interface {
	Checksum(text string) string
	Output(output string, textList map[int]string) error
}

// AlgorithmConverter 対応するハッシュアルゴリズムに変換する
type AlgorithmConverter struct {
	ConverterModule ConverterInterface
}

// NewAlgorithmConverter AlgorithmConverterのコンストラクタ
func NewAlgorithmConverter(algorithm string) *AlgorithmConverter {
	switch algorithm {
	case "sha256":
		return &AlgorithmConverter{ConverterModule: domain.Sha256{}}
	}

	return nil
}

// Exec 変換実行
func (c AlgorithmConverter) Exec(input, output string) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumber int
	textList := map[int]string{}
	wg := &sync.WaitGroup{}
	var mux sync.Mutex
	for scanner.Scan() {
		wg.Add(1)
		lineNumber++
		line := scanner.Text()
		go func(l string, ln int) {
			defer wg.Done()
			mux.Lock()
			textList[ln] = c.ConverterModule.Checksum(l)
			mux.Unlock()
		}(line, lineNumber)
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := c.ConverterModule.Output(output, textList); err != nil {
		return err
	}

	return nil
}
