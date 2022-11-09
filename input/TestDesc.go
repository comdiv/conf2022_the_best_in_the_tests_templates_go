package input

import (
	"fmt"
	"strconv"
	"strings"
)

// TestDesc
// Описание тестов.
// Входной файл должен выглядеть так:
//
//author |input                            |stringToProcessed		|isDisabled	|commentOnFailure                                		|publishTime
//harisov|паспорт Харисов Д.И. 1008 123848 |==PASSPORT_RF:1009123848|false     	|Не удалось определить паспорт РФ ФЛ  |2022-01-01T00:12:28.642595Z
//harisov|Паспорт Харисов Д.И. 10090 123848|=?PASSPORT_RF:1009123848|false     	|Не удалось определить паспорт РФ ФЛ|2022-01-01T00:22:28.642595Z
//
// Или в упрощенном формате так:
// паспорт Харисов Д.И. 1008 123848 -> PASSPORT_RF:1009123848 #Не удалось определить паспорт РФ ФЛ
// Паспорт Харисов Д.И. 10090 123848 -> =?PASSPORT_RF:1009123848 #Не удалось определить паспорт РФ ФЛ

type TestDesc struct {
	// Автора теста
	Author string

	// Вход
	Input string

	// Ожидаемый результат (в нормализованном виде), дополняется == если нет префикса
	Expected string

	// Признак отключения (или исчез в исходниках или иная причина)
	IsDisabled bool

	// Комментарий к тесту
	CommentOnFailure string

	// Время публикации теста
	// Техническое поле, необходимое для сведения
	// При работе участников - не используется
	PublishTime string
}

// DEFAULT_HEADER - заголовок всех файлов с описаниями тестов
const DEFAULT_HEADER = "author|input|expected|isDisabled|commentOnFailure|publishTime"

// DEFAULT_FIELDS_COUNT - кол-во полей в каждой строке с описаниями тестов
var DEFAULT_FIELDS_COUNT = strings.Count(DEFAULT_HEADER, DEFAULT_COLUMN_DELIMITER) + 1

// DEFAULT_COLUMN_DELIMITER - разделитель, использующийся в файлах с описаниями тестов
const DEFAULT_COLUMN_DELIMITER = "|"

// DISABLED_TEST_SYMBOL Префикс ОПИСАНИЯ теста - тест отключен
const DISABLED_TEST_SYMBOL = "!"

func (td *TestDesc) ToCsvString() string {
	return strings.Join(
		[]string{
			td.Author,
			td.Input,
			td.Expected,
			strconv.FormatBool(td.IsDisabled),
			td.CommentOnFailure,
			td.PublishTime,
		},
		DEFAULT_COLUMN_DELIMITER)
}

func (td *TestDesc) ToLocalString() string {
	if len(td.CommentOnFailure) != 0 {
		return fmt.Sprintf("%s -> %s # %s", td.Input, td.Expected, td.CommentOnFailure)
	} else {
		return fmt.Sprintf("%s -> %s", td.Input, td.Expected)
	}
}

func (td *TestDesc) BizKey() string {
	return fmt.Sprintf("%s:%s->%s", td.Author, td.Input, td.Expected)
}
