package table

type Orthogonal interface {
	GetTestCaseData(table Table) (testCase [][]string, err error)
}
