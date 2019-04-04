package table

import "fmt"

const (
	column = 3
	row    = 4
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

	indexName := table.IndexName
	factorName := table.FactorName
	testCase = make([][]string, row)
	for i := 0; i < row; i++ {
		testCase[i] = make([]string, column)
	}
	for i := 0; i < column; i++ {
		testCase[0][i] = indexName[i]
		fmt.Println(i)
	}
	for i := 1; i < row; i++ {
		for j := 0; j < column; j++ {
			factorIndex := tableData[i-1][j]
			testCase[i][j] = factorName[factorIndex]
		}

	}
	return
}
