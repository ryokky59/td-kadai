package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"sort"
)

// Sha256 SHA-256のハッシュアルゴリズムの構造体
type Sha256 struct{}

// Checksum チェックサムを付与する
func (s Sha256) Checksum(text string) string {
	sha256 := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sha256[:])
}

// Output 出力する
func (s Sha256) Output(output string, textList map[int]string) error {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	var keys []int
	for k := range textList {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		_, err := out.WriteString(textList[k] + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
