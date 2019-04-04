package table

const (
	column = 4
	row    = 5
)

var tableData = [][]int{
	{0, 0, 0},
	{0, 1, 1},
	{1, 0, 1},
	{1, 1, 0}}

type L423 struct {
	data  [][]int
	index [column]string
}

func (t L423) init() {
	t.data = tableData
}

func (t L423) GetTestCaseData(table Table) (testCase [][]string, err error) {

	indexName := table.indexName
	factorName := table.factorName
	for i := 0; i < column; i++ {
		testCase[1][i+1] = indexName[i]
	}
	for i := 2; i < row; i++ {
		for j := 1; j < column; j++ {
			factorIndex := tableData[i-2][j-1]
			testCase[i][j] = factorName[factorIndex]
		}

	}
	return
}
