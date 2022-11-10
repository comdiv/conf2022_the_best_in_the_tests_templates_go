package conf2022_the_best_in_the_tests_templates_go

import (
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/test"
	"os"
	"path/filepath"
	"testing"
)

// Запуск только БАЗОВЫХ и ЛОКАЛЬНЫХ тестов участника
//
// # Файл с базовыми тестами - base.csv
//
// Файл с локальными тестами - local.csv
func Test_local(t *testing.T) {
	// Раскомментировать если хочется отдельно тестировать base.csv + local.csv
	currentDir, _ := os.Getwd()

	files := []test.TestDescFile{
		{Path: filepath.Join(currentDir, "base.csv"), Type: test.BASE},
		{Path: filepath.Join(currentDir, "local.csv"), Type: test.LOCAL},
	}

	testBase := test.TestBase{TestFiles: files}

	testBase.Run(t)

}
