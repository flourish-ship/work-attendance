package pagination

type Pagination struct {
	Records  interface{} `json:"records"`
	Total    int         `json:"total"`
	Pages    int         `json:"pages"`
	PageSize int         `json:"pageSize"`
	Current  int         `json:"current"`
}

func (pagination *Pagination) Offset() int {
	if pagination.Current == 0 {
		return 0
	}

	return (pagination.Current - 1) * pagination.PageSize
}

func (pagination *Pagination) SetTotal(total int64) {
	if total == 0 {
		return
	}

	pagination.Total = int(total)
	if pagination.Total%pagination.PageSize > 0 {
		pagination.Pages = int(pagination.Total/pagination.PageSize) + 1
	} else {
		pagination.Pages = int(pagination.Total / pagination.PageSize)
	}
}
