package controller

import (
	"DangoMark/model/params"
	"DangoMark/model/result"
	"DangoMark/model/tables"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// 新建职位
func (Controller Controller) Init(ctx *gin.Context, user tables.User) {

	JSONSuccess(ctx, http.StatusOK, user)
}

// 获取图片数据
func (Controller Controller) GetImageData(ctx *gin.Context, user tables.User) {

	var GetImageData params.GetImageData
	if err := ctx.ShouldBindBodyWith(&GetImageData, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	count, image_data := Controller.DangoDB.SelectImageDataByLanguageAndStatus(GetImageData.Language, GetImageData.Status)
	if image_data.ID == 0 {
		JSONSuccess(ctx, http.StatusOK, make(map[int]int))
	} else {
		var ImageDataResult result.ImageDataResult
		ImageDataResult.Total = count
		ImageDataResult.Data.ID = image_data.ID
		ImageDataResult.Data.Url = image_data.Url
		ImageDataResult.Data.Language = image_data.Language
		json.Unmarshal([]byte(image_data.Suggestion), &ImageDataResult.Data.Suggestion)
		ImageDataResult.Data.MarkResult = image_data.MarkResult
		ImageDataResult.Data.QualityResult = image_data.QualityResult
		ImageDataResult.Data.CoordinateJson = image_data.CoordinateJson
		ImageDataResult.Data.Status = image_data.Status
		JSONSuccess(ctx, http.StatusOK, ImageDataResult)
	}

}
