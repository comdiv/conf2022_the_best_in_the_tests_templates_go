package test

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/input"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/parser"
	"os"
	"regexp"
	"strings"
	"testing"
)

type TestFileType int

const (
	BASE  TestFileType = 0
	LOCAL TestFileType = 1
	MAIN  TestFileType = 2
)

type TestDescFile struct {
	Path string
	Type TestFileType
}

type TestBase struct {
	TestFiles []TestDescFile
}

var MY_LOGIN = "harisov"
var docParser = parser.RandomSuccessfulParser{}

func (base *TestBase) Run(t *testing.T) {
	stat := TestStatistics{
		OwnerLogin:   MY_LOGIN,
		IsBasePass:   true,
		LocalResults: []TestResult{},
		MainResults:  []TestResult{},
	}

	for _, testFile := range base.TestFiles {
		switch testFile.Type {
		case BASE:
			stat.runBaseTest(testFile.Path, t)
		case LOCAL:
			stat.runLocalTest(testFile.Path, t)
		case MAIN:
			stat.runMainTest(testFile.Path, t)

			stat.makeReport()
		}

	}
}

func runTest(t *testing.T, desc input.TestDesc) bool {
	expected := output.Parse(desc.StringToProcessed)
	testName := fmt.Sprintf("Входная строка - %s. Ожидаемый список доков - %+v", desc.StringToProcessed, expected.Docs)

	return t.Run(testName, func(innerT *testing.T) {
		actual := docParser.Parse(desc.StringToProcessed)

		if !expected.Match(actual) {
			fmt.Println(desc.CommentOnFailure)
			fmt.Printf("Входная строка - %s\n", desc.StringToProcessed)
			fmt.Printf("Ожидаемый список доков - %v\n", expected.Docs)
			fmt.Printf("Актуальный список доков - %v\n", actual)

			innerT.Fail()
		}
	},
	)
}

func (stat *TestStatistics) runBaseTest(path string, t *testing.T) {
	testDescriptions := input.ParseFromFile(path)

	t.Run("Базовый функционал - базовые тесты", func(innerT *testing.T) {
		for _, testDesc := range testDescriptions {
			if !runTest(innerT, testDesc) {
				stat.IsBasePass = false
			}
		}
	})
}

func (stat *TestStatistics) runLocalTest(path string, t *testing.T) {
	testDescriptions := input.ParseFromFile(path)

	t.Run("Запуск локальных тестов", func(innerT *testing.T) {
		for _, testDesc := range testDescriptions {

			testResult := runTest(innerT, testDesc)

			stat.LocalResults = append(
				stat.LocalResults,
				TestResult{
					Author:            MY_LOGIN,
					StringToProcessed: testDesc.StringToProcessed,
					IsPass:            testResult,
				})
		}
	})
}

func (stat *TestStatistics) runMainTest(path string, t *testing.T) {
	testDescriptions := input.ParseFromFile(path)

	groupByAuthor := make(map[string][]input.TestDesc)

	for _, testDesc := range testDescriptions {
		groupByAuthor[testDesc.Author] = append(groupByAuthor[testDesc.Author], testDesc)
	}

	for author, testDescs := range groupByAuthor {
		if strings.ToLower(author) == strings.ToLower(MY_LOGIN) {
			t.Run("Запуск своих тестов, которые есть в общих, но которых нет в локальном файле", func(innerT *testing.T) {
				for _, testDesc := range testDescs {
					if !contains(stat.LocalResults, testDesc.StringToProcessed) {
						testResult := runTest(innerT, testDesc)

						stat.LocalResults = append(
							stat.LocalResults,
							TestResult{
								Author:            MY_LOGIN,
								StringToProcessed: testDesc.StringToProcessed,
								IsPass:            testResult,
							})
					}
				}
			})
		} else {
			t.Run(fmt.Sprintf("Тесты от %s", author), func(innerT *testing.T) {

				for _, testDesc := range testDescs {
					testResult := runTest(innerT, testDesc)

					stat.MainResults = append(
						stat.MainResults,
						TestResult{
							Author:            author,
							StringToProcessed: testDesc.StringToProcessed,
							IsPass:            testResult,
						})
				}
			})
		}
	}

}

func (stat *TestStatistics) makeReport() {
	mainFile, _ := os.Create("report.md")
	defer mainFile.Close()

	appendAndPrint(mainFile, fmt.Sprintf("##### Owner`s login:%s", stat.OwnerLogin))
	if stat.IsBasePass {
		appendAndPrint(mainFile, "##### All basic tests were passed")
	} else {
		appendAndPrint(mainFile, "##### All basic tests were NOT passed")
	}
	appendAndPrint(mainFile, "")

	localPassedCount := countPassed(stat.LocalResults)
	appendAndPrint(mainFile, fmt.Sprintf("##### Your own tests: %d/%d", localPassedCount, len(stat.LocalResults)))
	appendAndPrint(mainFile, fmt.Sprintf("##### So, %d test(s) can get you points", localPassedCount))
	appendAndPrint(mainFile, "")

	groupByAuthor := make(map[string][]TestResult)

	for _, testResult := range stat.MainResults {
		groupByAuthor[testResult.Author] = append(groupByAuthor[testResult.Author], testResult)
	}

	appendAndPrint(mainFile, "##### Competitors:")
	for author, testResults := range groupByAuthor {
		countPassed := countPassed(testResults)

		appendAndPrint(mainFile, fmt.Sprintf("##### %s: you passed  %d/%d", author, countPassed, len(testResults)))
	}
	appendAndPrint(mainFile, "")

	appendLine(mainFile, "##### FULL_INFO")
	appendLine(mainFile, "|author|input|expected|result|")
	appendLine(mainFile, "|-----|-----|-----|-----|")

	for _, testResult := range stat.LocalResults {
		appendTestResult(mainFile, testResult)
	}

	for _, testResults := range groupByAuthor {
		for _, testResult := range testResults {
			appendTestResult(mainFile, testResult)
		}
	}
}

func appendTestResult(file *os.File, result TestResult) {
	splitStringToProcessed := regexp.MustCompile(output.INPUT_STRUCTURE_REGEX).FindStringSubmatch(result.StringToProcessed)
	stringToWrite := fmt.Sprintf(
		"|%s|%s|%s%s|%v|\n",
		result.Author,
		splitStringToProcessed[1],
		splitStringToProcessed[2],
		splitStringToProcessed[3],
		result.IsPass,
	)

	file.WriteString(stringToWrite)
}

func appendAndPrint(file *os.File, input string) {
	appendLine(file, input)
	fmt.Println(input)
}

func appendLine(file *os.File, input string) {
	file.WriteString(fmt.Sprintf("%s\n", input))
}

func countPassed(testResults []TestResult) int {
	count := 0
	for _, testResult := range testResults {
		if testResult.IsPass {
			count = +1
		}
	}

	return count
}

func contains(results []TestResult, stringToSearch string) bool {
	for _, testResult := range results {
		if strings.TrimSpace(testResult.StringToProcessed) == strings.TrimSpace(stringToSearch) {
			return true
		}
	}

	return false
}
