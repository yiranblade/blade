package table

type Table struct {
	tableType  int      `json:tableType`
	indexName  []string `json:indexName`
	factorName []string `json:factorName`
}
