package req

type TPageReq struct {
	Offset int `form:"offset" json:"offset" binding:"required" msg:"offset is required"`
	Limit  int `form:"limit" json:"limit" binding:"required" msg:"limit is required"`
}
