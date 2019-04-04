package table

import (
	"blade/common"
	"errors"
)

type TableFactory struct {
}

func NewTableFactory() TableFactory {
	return TableFactory{}
}

func (t TableFactory) GetOrthogonalByType(tableType int) (table Orthogonal, err error) {
	switch tableType {
	case common.L423:
		table = new(L423)
	default:
		err = errors.New("未知表")
	}
	return
}
