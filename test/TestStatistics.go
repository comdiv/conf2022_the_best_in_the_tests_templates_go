package test

type TestStatistics struct {
	OwnerLogin   string
	IsBasePass   bool
	LocalResults []TestResult
	MainResults  []TestResult
}

type TestResult struct {
	Author            string
	StringToProcessed string
	IsPass            bool
}
