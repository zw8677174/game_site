package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rd/pkg/app"
	"rd/pkg/e"
	"rd/pkg/logging"
	"rd/pkg/upload"
)

func UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		app.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		app.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		app.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		app.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	app.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
