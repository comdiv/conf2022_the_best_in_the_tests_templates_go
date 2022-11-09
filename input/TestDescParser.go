package input

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseOption - параметра конфигурации парсера входных файлов с описаниями тестов
type ParseOption struct {
	Author  string
	Publish time.Time
}

// Parse - основной метод парсинга входных файлов
// Входе - файл с описаниями тестов (в строгом csv формате, в упрощенном формате)
// Выход - набор описаний тестов или ошибка
func Parse(reader io.Reader, options *ParseOption) ([]*TestDesc, error) {
	scanner := bufio.NewScanner(reader)

	var format ParserFormat = UNDEFINED
	lineCounter := 0
	var line string
	var mainError error

	var result []*TestDesc

	for scanner.Scan() {
		line = scanner.Text()
		lineCounter += 1

		if len(line) == 0 || isComment(line) {
			continue
		}

		if format == UNDEFINED {
			f, e := detektFormat(line, options)

			if f == SKIP {
				return result, nil
			}

			if e != nil {
				mainError = e
				break
			}

			format = f

			if format == MAIN {
				// skip header
				continue
			}
		}

		td, e := parseLine(line, lineCounter, format, options)

		if e != nil {
			mainError = e
			break
		}

		result = append(result, td)
	}

	groupedByBizKey := map[string][]*TestDesc{}
	var duplicateBizKeys []string

	for _, element := range result {
		groupedByBizKey[element.BizKey()] = append(groupedByBizKey[element.BizKey()], element)
	}

	for key, value := range groupedByBizKey {
		if len(value) > 1 {
			duplicateBizKeys = append(duplicateBizKeys, key)
		}
	}

	if len(duplicateBizKeys) > 0 {
		mainError = errors.New(
			fmt.Sprintf("Обнаружены тесты дубли: [%s]", strings.Join(duplicateBizKeys, ", ")),
		)
	}

	return result, mainError
}

func parseLine(line string, lineNumber int, format ParserFormat, options *ParseOption) (*TestDesc, error) {
	switch format {
	case MAIN:
		return parseLineMain(line, lineNumber, options)
	case LOCAL:
		return parseLineLocal(line, lineNumber, options)
	default:
		return nil, errors.New("UNDEFINED no implements parsing")
	}
}

func parseLineMain(line string, lineNumber int, options *ParseOption) (*TestDesc, error) {
	fields := strings.Split(line, DEFAULT_COLUMN_DELIMITER)

	if len(fields) != DEFAULT_FIELDS_COUNT {
		return nil, errors.New(invalidLineStructureMessage(lineNumber))
	}

	author := strings.TrimSpace(fields[0])
	input := strings.TrimSpace(fields[1])
	expected := strings.TrimSpace(fields[2])
	isDisabled, boolParseError := strconv.ParseBool(strings.TrimSpace(fields[3]))
	commentOnFailure := strings.TrimSpace(fields[4])
	publishTime := strings.TrimSpace(fields[5])

	if boolParseError != nil {
		return nil, errors.New(invalidLineStructureMessage(lineNumber))
	}

	result := TestDesc{
		Author:           author,
		Input:            input,
		Expected:         expected,
		IsDisabled:       isDisabled,
		CommentOnFailure: commentOnFailure,
		PublishTime:      publishTime,
	}

	fullExpectetionParseError := checkExpectedResult(fmt.Sprintf("%s%s", input, expected), lineNumber)

	return &result, fullExpectetionParseError
}

func parseLineLocal(line string, lineNumber int, options *ParseOption) (*TestDesc, error) {
	commentFields := strings.Split(line, "#")
	var commentOnFailure string

	if len(commentFields) > 2 {
		return nil, errors.New(invalidLineStructureMessage(lineNumber))
	} else if len(commentFields) == 2 {
		commentOnFailure = strings.TrimSpace(commentFields[1])
	} else {
		commentOnFailure = ""
	}

	inputAndExpected := strings.Split(commentFields[0], "->")

	if len(inputAndExpected) != 2 || len(inputAndExpected[0]) == 0 || len(inputAndExpected[1]) == 0 {
		return nil, errors.New(invalidLineStructureMessage(lineNumber))
	}

	inputCandidate := strings.TrimSpace(inputAndExpected[0])
	expectedCandidate := strings.TrimSpace(inputAndExpected[1])

	if strings.HasPrefix(expectedCandidate, "==") ||
		strings.HasPrefix(expectedCandidate, "~=") ||
		strings.HasPrefix(expectedCandidate, "=?") ||
		strings.HasPrefix(expectedCandidate, "~?") {
		expectedCandidate = expectedCandidate
	} else {
		expectedCandidate = fmt.Sprintf("%s%s", "==", expectedCandidate)
	}

	isDisabled := false

	if strings.HasPrefix(inputCandidate, DISABLED_TEST_SYMBOL) {
		isDisabled = true
		inputCandidate = strings.TrimPrefix(inputCandidate, DISABLED_TEST_SYMBOL)
	}

	result := TestDesc{
		Author:           options.Author,
		Input:            inputCandidate,
		Expected:         expectedCandidate,
		IsDisabled:       isDisabled,
		CommentOnFailure: commentOnFailure,
		PublishTime:      options.Publish.Format(time.RFC3339),
	}

	fullExpectetionParseError := checkExpectedResult(fmt.Sprintf("%s%s", inputCandidate, expectedCandidate), lineNumber)

	return &result, fullExpectetionParseError
}

func detektFormat(line string, options *ParseOption) (ParserFormat, error) {
	if line == "404: Not Found" {
		if os.Getenv("IS_LOCAL_TEST_MODE") == "true" {
			return SKIP, nil
		} else {
			return UNDEFINED, fmt.Errorf("NOT_FOUND")
		}
	}
	var format ParserFormat = UNDEFINED
	var err error

	switch {
	case regexp.MustCompile(MAIN).MatchString(line):
		format = MAIN
	case regexp.MustCompile(LOCAL).MatchString(line):
		format = LOCAL
	}

	if format == UNDEFINED {
		err = errors.New("Не могу определить формат файла")
	}

	if format == LOCAL && len(options.Author) == 0 {
		err = errors.New("Не указано имя автора при парсинге из локального файла")
	}

	return format, err
}

func isComment(line string) bool {
	return regexp.MustCompile("^\\s*#.*$").MatchString(line)
}

func invalidLineStructureMessage(lineNumber int) string {
	return fmt.Sprintf("Неправильная структура строки %d", lineNumber)
}

func checkExpectedResult(fullExpectation string, lineNumber int) (error error) {
	defer func() {
		if err := recover(); err != nil {
			error = errors.New(fmt.Sprintf("Не правильный синтаксис проверки строка %d - '%s'", lineNumber, fullExpectation))
		}
	}()

	output.ParseExpectedResult(fullExpectation)

	return error
}

// Формат парсинга - либо MAIN (строгий csv-файл), либо LOCAL (упрощенный формат)
type ParserFormat string

const (
	UNDEFINED = ""
	LOCAL     = "^([\\s\\S]*?[^~=?]+)->(==|~=|=\\?|~\\?)?([^~=?]+[\\s\\S]*?)$"
	MAIN      = "^\\s*author\\s*\\|\\s*input.*$"
	SKIP      = ""
)
