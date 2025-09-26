package entity

// PageRequest
//
// Request page params.
type PageRequest struct {
	Current  int      `form:"current" binding:"min=1"`
	PageSize int      `form:"pageSize" binding:"min=1,max=100"`
	Asc      []string `form:"asc"`
	Desc     []string `form:"desc"`
}
