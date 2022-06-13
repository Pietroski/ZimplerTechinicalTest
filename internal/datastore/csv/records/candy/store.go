package candy_store

import "io"

type (
	Store interface {
		Querier
	}

	candyRecord struct {
		fileOperator *io.Reader
	}
)

//func NewCandyStore(
//	file io.Reader,
//	fileOperator func(r io.Reader) *io.Reader,
//) Store {
//	cr := &candyRecord{
//		fileOperator: fileOperator(file),
//	}
//
//	return cr
//}
