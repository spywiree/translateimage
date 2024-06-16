package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/spywiree/translateimage"
	// languagecodes "github.com/spywiree/langcodes"
)

type RequestData struct {
	URL             []string `json:"URL"`
	SL              string   `json:"SL"`
	TL              string   `json:"TL"`
	ID              string   `json:"ID"`
	DeleteMode      int      `json:"DeleteMode"`
	AccessKeyId     string   `json:"accessKeyId"`
	AccessKeySecret string   `json:"accessKeySecret"`
	Endpoint        string   `json:"endpoint"`
	BucketName      string   `json:"bucketName"`
}

func aliyunoss(c *gin.Context) {

	var requestData RequestData
	// 解析传入的 JSON 数据并映射到 requestData 结构体上
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	// 在这里处理请求数据，例如打印出来或进行其他操作
	fmt.Printf("Received data: %+v\n", requestData)
	//处理请求和返回数据
	downloadDir := "downloadDir" + "/" + requestData.ID
	outputDir := "outputDir" + "/" + requestData.ID

	// 将接收到的字符串转换为 LanguageCode
	SL, ok := StringToLanguageCode(requestData.SL)
	if !ok {
		// 如果转换失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language code"})
		return
	}
	// 将接收到的字符串转换为 LanguageCode
	TL, ok := StringToLanguageCode(requestData.TL)
	if !ok {
		// 如果转换失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language code"})
		return
	}
	//获取当前工作目录
	// 获取执行文件的绝对路径
	currentPath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		currentPath = "当前工作目录获取失败"
	}

	TranslateList, err := TranslateImages(requestData.URL, downloadDir, outputDir, SL, TL)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": " err", "TranslateImages": "TranslateImages 传入参数错误"})

	}
	log.Print(TranslateList)

	var urllistMap map[string]string = make(map[string]string)

	for imagename := range TranslateList {

		ts := Timestamps() // 假设这个函数返回字符串格式的时间戳
		objectKey := "translatedir/" + requestData.ID + "/" + ts + "/" + TranslateList[imagename]
		log.Print("objectKey:", objectKey)

		localFilePath := outputDir + "/" + TranslateList[imagename]
		log.Print("localFilePath:", localFilePath)
		// 上传翻译后的文件到 Alioss
		url, err := Aliyunoss(requestData.AccessKeyId, requestData.AccessKeySecret, requestData.Endpoint, requestData.BucketName, objectKey, localFilePath)
		if err != nil {
			log.Print(err)
		}
		urllistMap[imagename] = url

	}

	var delStatus string
	// 根据 DeleteMode 构建要删除的目录路径
	switch requestData.DeleteMode {
	case 1:
		// 删除当前ID目录
		DIR := downloadDir
		ODIR := outputDir
		FileDelete(DIR)
		FileDelete(ODIR)
		delStatus = "已删除商品当前下载目录"
	case 2:
		// 删除主目录
		FileDelete("downloadDir")
		FileDelete("outputDir")
		delStatus = "已删除商品下载主目录"
	default:
		// 不删除数据，直接返回或做其他处理
		log.Println("No directory will be deleted.")
		delStatus = "未删除商品下载目录"
	}

	// 返回成功响应或其他你需要的响应数据
	c.JSON(200, gin.H{"status": "success", "ossurl": urllistMap, "currentPath": currentPath, "outputDir": outputDir, "TL_Name": TranslateList, "delstatus": delStatus})

}

func translation(c *gin.Context) {

	var requestData RequestData
	// 解析传入的 JSON 数据并映射到 requestData 结构体上
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	// 在这里处理请求数据，例如打印出来或进行其他操作
	fmt.Printf("Received data: %+v\n", requestData)
	//处理请求和返回数据
	downloadDir := "downloadDir" + "/" + requestData.ID
	outputDir := "outputDir" + "/" + requestData.ID

	// 将接收到的字符串转换为 LanguageCode
	SL, ok := StringToLanguageCode(requestData.SL)
	if !ok {
		// 如果转换失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language code"})
		return
	}
	// 将接收到的字符串转换为 LanguageCode
	TL, ok := StringToLanguageCode(requestData.TL)
	if !ok {
		// 如果转换失败，返回错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language code"})
		return
	}
	//获取当前工作目录
	// 获取执行文件的绝对路径
	currentPath, err := os.Executable()
	if err != nil {
		log.Println("Error getting executable path:", err)
		currentPath = "当前工作目录获取失败"
	}

	TranslateList, err := TranslateImages(requestData.URL, downloadDir, outputDir, SL, TL)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": " err", "TranslateImages": "TranslateImages 传入参数错误"})

	}
	log.Print(TranslateList)

	var delStatus string
	// 根据 DeleteMode 构建要删除的目录路径
	switch requestData.DeleteMode {
	case 1:
		// 删除当前ID目录
		DIR := downloadDir
		ODIR := outputDir
		FileDelete(DIR)
		FileDelete(ODIR)
		delStatus = "已删除商品当前下载目录"
	case 2:
		// 删除主目录
		FileDelete("downloadDir")
		FileDelete("outputDir")
		delStatus = "已删除商品下载主目录"
	default:
		// 不删除数据，直接返回或做其他处理
		log.Println("No directory will be deleted.")
		delStatus = "未删除商品下载目录"
	}

	// 返回成功响应或其他你需要的响应数据
	c.JSON(200, gin.H{"status": "success", "currentPath": currentPath, "outputDir": outputDir, "TL_Name": TranslateList, "delstatus": delStatus})

}

// 删除指定目录的函数
func FileDelete(Dir string) {

	err := os.RemoveAll(Dir)
	if err != nil {
		// 处理错误
		log.Printf("Error deleting download directory: %v\n", err)
	} else {
		log.Printf("Download directory %s has been deleted successfully.\n", Dir)
	}
}

// 时间戳函数
func Timestamps() string {
	// 获取当前时间
	currentTime := time.Now()

	// 如果你需要毫秒级的时间戳
	timestampMillis := currentTime.UnixNano() / int64(time.Millisecond)
	log.Print("时间戳（毫秒级）:", timestampMillis)

	// 将毫秒级时间戳转换为字符串
	timestampStr := strconv.FormatInt(timestampMillis, 10) // 第二个参数是基数，10表示十进制
	return timestampStr
}
