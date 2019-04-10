package table

import "fmt"

type L827 struct {
	row    int
	column int
	data   [][]int
}
type Asd struct {
}

func NewL827() *L827 {
	var l827 = new(L827)
	l827.row = 8
	l827.column = 7
	l827.data = [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 1, 1},
		{0, 1, 1, 1, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 1, 0},
		{1, 1, 0, 1, 0, 0, 1}};
	return l827
}

func (t L827) GetTestCaseData(table Table) (testCase [][]string, err error) {
	row := t.row
	column := t.column
	indexName := table.IndexName
	factorName := table.FactorName
	tableData := t.data
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
