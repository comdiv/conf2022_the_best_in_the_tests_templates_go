package conf2022_the_best_in_the_tests_templates_go

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/test"
	"os"
	"path/filepath"
	"testing"
)

// Запуск только БАЗОВЫХ ТЕСТОВ
//
// Файл с тестами - base.csv
func Test_base(t *testing.T) {
	currentDir, _ := os.Getwd()

	files := []test.TestDescFile{
		{Path: filepath.Join(currentDir, "base.csv"), Type: test.BASE},
	}

	testBase := test.TestBase{TestFiles: files}

	testBase.Run(t)
}
