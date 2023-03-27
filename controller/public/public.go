package public

import (
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func UploadImage(ctx iris.Context) {
	//获取上传文件的信息
	file, info, err := ctx.FormFile("file")
	if err != nil {
		app.Error(ctx, -1, err, "文件上传失败")
		return
	}
	//defer file.Close()
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()
	// 获取文件名称
	fname := info.Filename
	//把文件上传到哪里
	//out, err := os.OpenFile("uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	// 创建一个新的文件
	ext := filepath.Ext(fname)
	filename := fname[0 : len(fname)-len(ext)]
	timestamp := time.Now().Format("20060102_150405")
	newFilename := fmt.Sprintf("%s_%s%s", filename, timestamp, ext)

	newFilePath := filepath.Join("uploads", newFilename)

	out, err := os.Create("uploads/" + newFilename)
	if err != nil {
		app.Error(ctx, -1, err, "")
		return
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	// 拷贝文件位置
	_, err = io.Copy(out, file)
	if err != nil {
		app.Error(ctx, -1, err, "文件上传失败!")
	} else {
		app.OK(ctx, filepath.ToSlash(newFilePath), "文件上传成功")
	}
}
