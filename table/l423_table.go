package table

const (
	column = 4
)

var tableData = [][]int{
	{1, 1, 1},
	{1, 2, 2},
	{2, 1, 2},
	{2, 2, 1}}

type L423 struct {
	data  [][]int
	index [column]string
}

func (t L423) init() {
	t.data = tableData
}

func (t L423) getTestCaseData(table Table) (testCase [][]string, err error) {

	for ; ; {

	}
	return
}
