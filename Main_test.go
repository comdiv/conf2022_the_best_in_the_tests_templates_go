package conf2022_the_best_in_the_tests_templates_go

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/test"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

// Запуск только БАЗОВЫХ, ЛОКАЛЬНЫХ тестов участника и ОБЩИХ тестов всех участников
//
// # Файл с базовыми тестами - base.csv
//
// # Файл с локальными тестами - local.csv
//
// Файл с общими тестами всех участников main.csv - выкачивается из общего репозитория
func Test_main(t *testing.T) {
	currentDir, _ := os.Getwd()

	createMainFile()

	files := []test.TestDescFile{
		{Path: filepath.Join(currentDir, "base.csv"), Type: test.BASE},
		{Path: filepath.Join(currentDir, "local.csv"), Type: test.LOCAL},
		{Path: filepath.Join(currentDir, "main.csv"), Type: test.MAIN},
	}

	testBase := test.TestBase{TestFiles: files}

	testBase.Run(t)
}

func createMainFile() {
	requestURL := "https://raw.githubusercontent.com/spectrum-data/conf2022_the_best_in_the_tests_templates_base/main/main.csv"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)

	fmt.Printf("client: got response!\n")

	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	mainFile, e := os.Create("main.csv")
	defer mainFile.Close()

	if e != nil {
		panic(e)
	}

	mainFile.Write(resBody)
}
