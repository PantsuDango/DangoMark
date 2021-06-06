package params

type ModActIndex struct {
	Module string `json:"Module" form:"Module" binding:"required"`
	Action string `json:"Action" form:"Action" binding:"required"`
}
