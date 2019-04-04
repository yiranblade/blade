package table

type Table struct {
	TableType  int      `json:tableType`
	IndexName  []string `json:indexName`
	FactorName []string `json:factorName`
}
