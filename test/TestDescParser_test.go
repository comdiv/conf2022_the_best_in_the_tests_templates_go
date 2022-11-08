package test

import (
	"fmt"
	"github.com/spectrum-data/conf2022_the_best_in_the_tests_templates_go/input"
	"os"
	"strings"
	"testing"
	"time"
)

func CreateTempFileWithContext(context string) string {
	file, err := os.CreateTemp(os.TempDir(), "test_descriptions_*.csv")

	if err != nil {
		panic(err)
	}

	file.WriteString(context)
	file.Sync()
	file.Close()

	return file.Name()
}

var timePublish = "2022-11-09T11:01:22.123567Z"

var test1 = input.TestDesc{
	Author:           "some_author",
	Input:            "1234567890",
	Expected:         "==PASSPORT_RF:1234567890",
	IsDisabled:       false,
	CommentOnFailure: "some comment 1",
	PublishTime:      timePublish,
}

var test2 = input.TestDesc{
	Author:           "some_author",
	Input:            "6511111111",
	Expected:         "==INN_FL:123456789012",
	IsDisabled:       false,
	CommentOnFailure: "some comment 2",
	PublishTime:      timePublish,
}

var options = input.ParseOption{Author: "some_author", Publish: time.Now()}

func checkContains(t *testing.T, slice []*input.TestDesc, element input.TestDesc) {
	containsElement := false

	for _, td := range slice {
		if td.Author == element.Author &&
			td.Input == element.Input &&
			td.Expected == element.Expected &&
			td.IsDisabled == element.IsDisabled &&
			td.CommentOnFailure == element.CommentOnFailure {
			containsElement = true
		}
	}

	if !containsElement {
		t.Errorf("не содержит ожидаемый элемент - %+v", element)
	}
}

func Test_Contains(t *testing.T) {
	testDesc := input.TestDesc{
		Author:           "some_author",
		Input:            "1234567890",
		Expected:         "==PASSPORT_RF:1234567890",
		IsDisabled:       false,
		CommentOnFailure: "some comment 1",
		PublishTime:      timePublish,
	}

	theSameTestDesc := input.TestDesc{
		Author:           "some_author",
		Input:            "1234567890",
		Expected:         "==PASSPORT_RF:1234567890",
		IsDisabled:       false,
		CommentOnFailure: "some comment 1",
		PublishTime:      timePublish,
	}

	checkContains(t, []*input.TestDesc{
		&theSameTestDesc,
	}, testDesc)
}

func Test_ParserTest(t *testing.T) {
	t.Run("обычные кейсы", func(contextT *testing.T) {
		contextT.Run("чтение полного и нормального файла", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					input.DEFAULT_HEADER,
					test1.ToCsvString(),
					test2.ToCsvString(),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			result, err := input.Parse(file, &options)

			if err != nil {
				testT.Errorf("Не ожидалась ошибка - %+v", err)
			}

			checkContains(testT, result, test1)
			checkContains(testT, result, test2)
		})

		contextT.Run("чтение локального и нормального файла", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					test1.ToLocalString(),
					test2.ToLocalString(),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			result, err := input.Parse(file, &options)

			if err != nil {
				testT.Errorf("Не ожидалась ошибка - %+v", err)
			}

			checkContains(testT, result, test1)
			checkContains(testT, result, test2)
		})
	})

	t.Run("relax кейсы", func(contextT *testing.T) {
		contextT.Run("разрешены пустые строки в перемешку с комментариями ", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					"# а вот и наш main",
					"",
					"",
					input.DEFAULT_HEADER,
					"",
					" # а вот и тесты нашего молодца",
					test1.ToCsvString(),
					"", "",
					test2.ToCsvString(),
					"",
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			result, err := input.Parse(file, &options)

			if err != nil {
				testT.Errorf("Не ожидалась ошибка - %+v", err)
			}

			checkContains(testT, result, test1)
			checkContains(testT, result, test2)
		})

		contextT.Run("в локальных тестах можно пропускать `==`", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					"1111111111 -> PASSPORT_RF+:1111111111",
					"2222222222 -> ==PASSPORT_RF+:2222222222",
					"3333333333 -> ~=PASSPORT_RF+:3333333333",
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			result, err := input.Parse(file, &options)

			if err != nil {
				testT.Fatalf("Не ожидалась ошибка - %+v", err)
			}

			if result[0].Expected != "==PASSPORT_RF+:1111111111" {
				testT.Fail()
			}
			if result[1].Expected != "==PASSPORT_RF+:2222222222" {
				testT.Fail()
			}
			if result[2].Expected != "~=PASSPORT_RF+:3333333333" {
				testT.Fail()
			}
		})
	})

	t.Run("ошибки", func(contextT *testing.T) {
		contextT.Run("неведомый локальный формат", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					"я хочу == так писать тесты",
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)

			if err == nil || err.Error() != "Не могу определить формат файла" {
				testT.Errorf("Ожидалась ошибка - 'Не могу определить формат файла'")
			}
		})

		contextT.Run("неведомый полный формат, сбитый CSV", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					"some" + input.DEFAULT_HEADER,
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)

			if err == nil || err.Error() != "Не могу определить формат файла" {
				testT.Errorf("Ожидалась ошибка - 'Не могу определить формат файла'")
			}
		})

		contextT.Run("завалена одна из строк", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					input.DEFAULT_HEADER,
					test1.ToCsvString(),
					strings.Replace(test2.ToCsvString(), input.DEFAULT_COLUMN_DELIMITER, "~", 1),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)

			if err == nil || err.Error() != "Неправильная структура строки 3" {
				testT.Errorf("Ожидалась ошибка - 'Неправильная структура строки 3'")
			}
		})

		contextT.Run("неправильное условие - мейн", func(testT *testing.T) {
			incorrectExpected := test1
			incorrectExpected.Expected = "==труляля:1111"
			incorrectExpected.Author = "y"

			content := strings.Join(
				[]string{
					input.DEFAULT_HEADER,
					test1.ToCsvString(),
					test2.ToCsvString(),
					incorrectExpected.ToCsvString(),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)

			if err == nil || err.Error() != "Не правильный синтаксис проверки строка 4 - '1234567890==труляля:1111'" {
				testT.Errorf("Ожидалась ошибка - 'Не правильный синтаксис проверки строка 4 - '1234567890==труляля:1111''")
			}
		})

		contextT.Run("неправильное условие - локальный", func(testT *testing.T) {
			incorrectExpected := test1
			incorrectExpected.Expected = "==труляля:1111"
			incorrectExpected.Author = options.Author

			content := strings.Join(
				[]string{
					test1.ToLocalString(),
					test2.ToLocalString(),
					incorrectExpected.ToLocalString(),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)

			if err == nil || err.Error() != "Не правильный синтаксис проверки строка 3 - '1234567890==труляля:1111'" {
				testT.Errorf("Ожидалась ошибка - 'Не правильный синтаксис проверки строка 3 - '1234567890==труляля:1111''")
			}
		})

		contextT.Run("задвоения запрещены", func(testT *testing.T) {
			content := strings.Join(
				[]string{
					input.DEFAULT_HEADER,
					test1.ToCsvString(),
					test1.ToCsvString(),
				},
				"\n")

			filePath := CreateTempFileWithContext(content)
			file, _ := os.Open(filePath)
			defer file.Close()

			_, err := input.Parse(file, &options)
			if err == nil || err.Error() != fmt.Sprintf("Обнаружены тесты дубли: [%s:%s->%s]", test1.Author, test1.Input, test1.Expected) {
				testT.Errorf("Ожидалась ошибка - 'Обнаружены тесты дубли: [some_author:1234567890->==PASSPORT_RF:1234567890]'")
			}
		})
	})
}
