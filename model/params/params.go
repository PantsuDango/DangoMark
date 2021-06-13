package params

type ModActIndex struct {
	Module string `json:"Module" form:"Module" binding:"required"`
	Action string `json:"Action" form:"Action" binding:"required"`
}

type FileInfo struct {
	Retcode int    `json:"retcode" form:"retcode"`
	URL     string `json:"url"     form:"url"      binding:"required"`
}
