package vo

// 分页参数
type PageVo struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	Start       int   `json:"start"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
}

func GetPageParam(param *PageVo) (res *PageVo) {
	if param == nil {
		res.CurrentPage = 0
		res.PageSize = 10
		res.Limit = 10
		res.Start = 0

		return
	}

	var (
		current int
		size    int
	)
	res = new(PageVo)

	if param.CurrentPage < 0 {
		res.CurrentPage = 0
		current = 0
	} else {
		res.CurrentPage = param.CurrentPage
		current = param.CurrentPage - 1
	}

	if param.PageSize < 0 {
		res.PageSize = 10
		size = 10
	} else {
		res.PageSize = param.PageSize
		size = param.PageSize
	}
	res.Start = current * size
	res.Limit = param.PageSize
	res.Total = param.Total

	return
}

type PageResult struct {
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
}
