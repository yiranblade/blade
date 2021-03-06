package table

type L423 struct {
	row    int
	column int
	data   [][]int
}

func NewL423() *L423 {
	var l423 = new(L423)
	l423.row = 4
	l423.column = 3
	l423.data = [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0}};
	return l423
}

func (t L423) GetTestCaseData(table Table) (testCase [][]string, err error) {
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
	}
	for i := 1; i < row; i++ {
		for j := 0; j < column; j++ {
			factorIndex := tableData[i-1][j]
			testCase[i][j] = factorName[factorIndex]
		}

	}
	return
}
