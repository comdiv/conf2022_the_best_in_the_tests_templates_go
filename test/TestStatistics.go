package test

// TestStatistics - статистика запусков тестов
//
// # OwnerLogin - логин участника
//
// # IsBasePass - пройдены ли базовые тесты
//
// # LocalResults - результаты запуска локальных тестов
//
// # MainResults - Результаты запуска общих тестов
type TestStatistics struct {
	OwnerLogin   string
	IsBasePass   bool
	LocalResults []TestResult
	MainResults  []TestResult
}

// TestResult - результат одиночного запуска теста
//
// # Author - Логин автора теста
//
// # StringToProcessed - Входная строка
//
// # IsPass - Пройден ли тест
type TestResult struct {
	Author            string
	StringToProcessed string
	IsPass            bool
}
