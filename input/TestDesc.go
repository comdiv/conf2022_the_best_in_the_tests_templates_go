package input

import (
	"encoding/csv"
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/output"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TestDesc
// Описание тестов.
// Входной файл должен выглядеть так:
//
//	author | number | stringToProcessed | isDisabled | commentOnFailure | publishTime
//	harisov | 1 | паспорт Харисов Д.И. 1009 123848==PASSPORT:1009123848 | false | Не удалось определить корректный паспорт ФЛ |
//	harisov | 2 | Паспорт Харисов Д.И. 10090 123848=?PASSPORT:1009123848 | false | Не удалось определить некорректный паспорт ФЛ |
type TestDesc struct {
	Author            string
	Number            int
	StringToProcessed string
	IsDisabled        bool
	CommentOnFailure  string
}

// ParseFromFile получает набор описаний тестов из файла
func ParseFromFile(filePath string) []TestDesc {
	file, _ := os.Open(filePath)
	defer file.Close()

	var reader = BuildCsvReader(file)

	var wasHeaderRead = false
	var result []TestDesc

	for {
		record, e := reader.Read()

		if e == io.EOF {
			break
		} else if e != nil {
			panic(e)
		}

		if !wasHeaderRead {
			wasHeaderRead = true
		} else {
			// Для каждого поля - убираем пробелы с начала и с конца
			for i, field := range record {
				record[i] = strings.TrimSpace(field)
			}

			// Парсим номер в int
			number, numberParseError := strconv.Atoi(record[1])

			if numberParseError != nil {
				panic(numberParseError)
			}

			// Парсим параметр отключения теста в bool
			isDisabled := strings.EqualFold(record[3], strconv.FormatBool(true))

			result = append(result, TestDesc{
				Author:            record[0],
				Number:            number,
				StringToProcessed: record[2],
				IsDisabled:        isDisabled,
				CommentOnFailure:  record[4],
			})
		}
	}

	return result
}

// TestDescFileValidateResult Результат валидации входного файла.
type TestDescFileValidateResult struct {
	IsValid bool
	errors  []string
}

/*
	Validate

Валидирует входной файл:
 1. наличие хедера
 2. кол-во разделителей в каждой строке = кол-ву разделителей в хедере
 3. тип полей Number и isDisabled
 4. что идентификатор теста (Author + Number) уникальный
*/
func Validate(filePath string) TestDescFileValidateResult {
	file, _ := os.Open(filePath)
	defer file.Close()

	var reader = BuildCsvReader(file)

	var errorMessages []string
	var testIdToLineNumber = make(map[string]int)

	fileInfo, _ := os.Stat(filePath)

	if fileInfo.Size() == 0 {
		errorMessages = append(errorMessages, "File is empty!")
	}

	var wasHeaderRead = false

	currentLineNumber := 1

	for {
		record, e := reader.Read()

		if e == io.EOF {
			break
		}

		if e != nil {
			errorMessages = append(errorMessages, FormErrorMessage(currentLineNumber, e.Error()))
		} else {
			if !wasHeaderRead {
				actualHeader := strings.Join(record, string(DEFAULT_COLUMN_DELIMITER))

				if actualHeader != DEFAULT_HEADER {
					errorMessages = append(errorMessages, FormErrorMessage(currentLineNumber, "Файл с описанием тест-кейсов не содержит ожидаемый заголовок `"+DEFAULT_HEADER+"'"))
				}

				wasHeaderRead = true
			} else {
				// Для каждого поля - убираем пробелы с начала и с конца
				for i, field := range record {
					record[i] = strings.TrimSpace(field)
				}

				match, _ := regexp.MatchString(output.INPUT_STRUCTURE_REGEX, record[2])

				if !match {
					errorMessages = append(errorMessages, FormErrorMessage(currentLineNumber, fmt.Sprintf("Описание теста не удолетворяет структуре - %s", output.INPUT_STRUCTURE_REGEX)))
				}

				if lineNumberFromMap, ok := testIdToLineNumber[record[0]+record[1]]; ok {
					errorMessages = append(errorMessages, FormErrorMessage(currentLineNumber, fmt.Sprintf("В строках с номерами %d и %d совпадает связка Author+Number", lineNumberFromMap, currentLineNumber)))
				} else {
					testIdToLineNumber[record[0]+record[1]] = currentLineNumber
				}

				// Пытаемся спарсить в int
				_, numberParseError := strconv.Atoi(record[1])

				if numberParseError != nil {
					errorMessages = append(errorMessages, FormErrorMessage(currentLineNumber, numberParseError.Error()))
				}

			}

			currentLineNumber += 1
		}
	}

	return TestDescFileValidateResult{
		IsValid: len(errorMessages) == 0,
		errors:  errorMessages,
	}
}

// PrintErrors выводит список ошибок при валидации
func (validation TestDescFileValidateResult) PrintErrors() {
	for _, errorMessage := range validation.errors {
		fmt.Println(errorMessage)
	}
}

func (validation TestDescFileValidateResult) GetErrors() *[]string {
	return &validation.errors
}

// FormErrorMessage формирует сообщение об ошибке, указывая номер строки
func FormErrorMessage(lineNumber int, message string) string {
	return fmt.Sprintf("В строке с номером %d ошибка - '%s'", lineNumber, message)
}

// BuildCsvReader формирует экземпляр csv.Reader с дефолтными настройками
func BuildCsvReader(file *os.File) *csv.Reader {
	reader := csv.NewReader(file)

	reader.Comment = '#'
	reader.FieldsPerRecord = DEFAULT_FIELDS_COUNT
	reader.Comma = DEFAULT_COLUMN_DELIMITER

	return reader
}

// DEFAULT_HEADER - заголовок всех файлов с описаниями тестов
const DEFAULT_HEADER = "author|number|stringToProcessed|isDisabled|commentOnFailure|publishTime"

// DEFAULT_FIELDS_COUNT - кол-во полей в каждой строке с описаниями тестов
const DEFAULT_FIELDS_COUNT = 6

// DEFAULT_COLUMN_DELIMITER - разделитель, использующийся в файлах с описаниями тестов
const DEFAULT_COLUMN_DELIMITER = '|'
