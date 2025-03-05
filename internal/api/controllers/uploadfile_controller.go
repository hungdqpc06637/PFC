package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

type UploadFileController struct {
	*BaseController
}

var UploadFile = &UploadFileController{}
var urlDir = `\\192.168.123.113\pfc\PFCModel_Picture\`

func (c *UploadFileController) UploadFilePFCModel(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	now := time.Now()
	timestampNano := now.UnixNano()
	destDir := urlDir

	fileNameRender := fmt.Sprintf("%d_%s", timestampNano, file.Filename)
	destPath := filepath.Join(destDir, fileNameRender)

	src, err := file.Open()
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, fileNameRender)
}

func (c *UploadFileController) UploadFilePFCModelFromFolderPFCModel(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Invalid form data")
		return
	}

	ModelName := ctx.PostForm("ModelName")
	ModelName = strings.ReplaceAll(ModelName, `/`, ``)
	if ModelName == "" {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, "Missing ModelName")
		return
	}

	now := time.Now()
	timestampNano := now.UnixNano()

	destDir := filepath.Join(urlDir, ModelName)

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Failed to create directory")
			return
		}
	}

	fileNameRender := fmt.Sprintf("%d_%s", timestampNano, file.Filename)
	destPath := filepath.Join(destDir, fileNameRender)

	src, err := file.Open()
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Failed to read file content")
		return
	}

	minSize := 30 * 1024
	if len(fileContent) < minSize {
		paddingSize := minSize - len(fileContent)
		padding := make([]byte, paddingSize)
		for i := range padding {
			padding[i] = ' '
		}
		fileContent = append(fileContent, padding...)
	}

	dst, err := os.Create(destPath)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	defer dst.Close()

	if _, err := dst.Write(fileContent); err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, "Failed to write file")
		return
	}

	response.OkWithData(ctx, fileNameRender)
}

func (c *UploadFileController) DownloadFilePFCModel(ctx *gin.Context) {
	var query struct {
		Filename string `form:"filename" binding:"required"`
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'filename' query parameter"})
		return
	}

	baseDir := urlDir

	filePath := filepath.Join(baseDir, query.Filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing file"})
		return
	}

	ctx.FileAttachment(filePath, query.Filename)
}

func (c *UploadFileController) DownloadFilePFCModelFromFolderPFCModel(ctx *gin.Context) {
	var query struct {
		ModelName string `form:"ModelName" binding:"required"`
		Filename  string `form:"filename" binding:"required"`
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid parameters"})
		return
	}

	baseDir := urlDir
	ModelName := strings.ReplaceAll(query.ModelName, `/`, ``)

	destDir := filepath.Join(baseDir, ModelName)
	filePath := filepath.Join(destDir, query.Filename)

	cleanFilePath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanFilePath, baseDir) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}

	if _, err := os.Stat(cleanFilePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing file"})
		return
	}

	ctx.FileAttachment(cleanFilePath, query.Filename)
}

func (c *UploadFileController) DeleteFilePFCModel(ctx *gin.Context) {
	var query struct {
		Filename string `form:"filename" binding:"required"`
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'filename' query parameter"})
		return
	}

	baseDir := urlDir

	filePath := filepath.Join(baseDir, query.Filename)

	err := os.Remove(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to delete file: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

func (c *UploadFileController) DeleteFilePFCModelFromFolderPFCModel(ctx *gin.Context) {
	var query struct {
		ModelName string `form:"ModelName" binding:"required"`
		Filename  string `form:"filename" binding:"required"`
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid parameters"})
		return
	}

	baseDir := urlDir
	ModelName := strings.ReplaceAll(query.ModelName, `/`, ``)
	destDir := filepath.Join(baseDir, ModelName)

	filePath := filepath.Join(destDir, query.Filename)

	cleanFilePath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanFilePath, baseDir) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}

	if _, err := os.Stat(cleanFilePath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing file"})
		return
	}

	if err := os.Remove(cleanFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete file: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

func (c *UploadFileController) DeleteModelDirectory(ctx *gin.Context) {
	var query struct {
		ModelName string `form:"ModelName" binding:"required"`
	}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'ModelName' parameter"})
		return
	}

	baseDir := urlDir
	ModelName := strings.ReplaceAll(query.ModelName, `/`, ``)
	dirPath := filepath.Join(baseDir, ModelName)

	cleanDirPath := filepath.Clean(dirPath)
	if !strings.HasPrefix(cleanDirPath, baseDir) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid directory path"})
		return
	}

	if _, err := os.Stat(cleanDirPath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Directory not found"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing directory"})
		return
	}

	if err := os.RemoveAll(cleanDirPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete directory: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Directory deleted successfully"})
}
