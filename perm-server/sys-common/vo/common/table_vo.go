package common

type TableVO struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}

func NewTableVO(total int64, rows interface{}) *TableVO {
	return &TableVO{
		Total: total,
		Rows:  rows,
	}
}
