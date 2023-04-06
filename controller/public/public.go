package public

import (
	"dang_go/internal/database"
	"dang_go/internal/model/shop"
	"dang_go/tools/app"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

/*UploadImage
* @Description: 文件上传
* @param ctx
 */
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
	// 检查uploads目录是否存在，如果不存在则创建该目录
	errs := os.MkdirAll("uploads", os.ModePerm)
	if errs != nil {
		log.Fatal(err)
	}
	// 创建一个新的文件
	ext := filepath.Ext(fname)
	filename := fname[0 : len(fname)-len(ext)]
	timestamp := time.Now().Format("20060102_150405")
	newFilename := fmt.Sprintf("%s_%s%s", filename, timestamp, ext)

	newFilePath := filepath.Join("uploads", newFilename)
	//把文件上传到哪里
	out, err := os.OpenFile("uploads/"+newFilename, os.O_WRONLY|os.O_CREATE, 0666)
	//out, err := os.Create("uploads/" + newFilename)
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

/*Home
* @Description: 首页统计(平台)
* @param ctx
 */
func Home(ctx iris.Context) {
	var merchant shop.Merchant // 商家
	var shops shop.Shop        // 商品
	var order shop.Order

	var tableMerchants int64
	database.DB.Model(&merchant).Count(&tableMerchants)
	var tableShops int64
	database.DB.Model(&shops).Count(&tableShops)
	var onShelves int64
	database.DB.Model(&shops).Where("status = 0").Count(&onShelves)
	var offTheShelf int64
	database.DB.Model(&shops).Where("status = 2").Count(&offTheShelf)
	var tableOrder int64
	database.DB.Model(&order).Count(&tableOrder)

	count := map[string]interface{}{
		"allMerchant": tableMerchants,
		"allShops":    tableShops,
		"allOrders":   tableOrder,
		"onShelves":   onShelves,
		"offTheShelf": offTheShelf,
		"a":           0,
		"b":           0,
		"c":           0,
		"d":           0,
		"e":           0,
		"f":           0,
		"j":           0,
		"h":           0,
		"i":           0,
		"g":           0,
		"k":           0,
		"l":           0,
		"m":           0,
		"n":           0,
	}
	app.OK(ctx, count, "")
}
