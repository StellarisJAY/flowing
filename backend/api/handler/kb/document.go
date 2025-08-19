package kb

import (
	"errors"
	"flowing/api/handler/common"
	"flowing/global"
	model "flowing/internal/model/kb"
	service "flowing/internal/service/kb"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListDocument(c *gin.Context) {
	var query model.DocumentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		panic(global.ErrBadRequest(err))
		return
	}
	res, total, err := service.ListDocument(c, query)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(200, common.PageResp(res, total))
}

func UploadDocument(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	files, ok := form.File["file"]
	if !ok {
		panic(global.ErrBadRequest(errors.New("请求中没有文件")))
	}
	id, ok := form.Value["knowledgeBaseId"]
	if !ok {
		panic(global.ErrBadRequest(errors.New("请求中没有知识库id")))
	}
	knowledgeBaseId, err := strconv.ParseInt(id[0], 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	fileObj, err := files[0].Open()
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	contentType := files[0].Header.Get("Content-Type")
	req := model.UploadDocumentReq{
		KnowledgeBaseId: knowledgeBaseId,
		FileName:        files[0].Filename,
		Size:            files[0].Size,
		FileObj:         fileObj,
		ContentType:     contentType,
	}
	if err := service.UploadDocument(c, req); err != nil {
		panic(err)
	}
	c.JSON(200, common.Ok())
}

func GetDownloadURL(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		panic(global.ErrBadRequest(err))
	}
	url, err := service.GetDownloadURL(c, id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, common.OkWithData(url))
}
