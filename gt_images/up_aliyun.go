package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main9() {

	// 直接在代码中设置Access Key ID和Access Key Secret
	const accessKeyId = "accessKeyId"         // 替换为你的Access Key ID
	const accessKeySecret = "accessKeySecret" // 替换为你的Access Key Secret
	const endpoint = "endpoint"               // 替换为你的OSS endpoint
	const bucketName = "bucketName"
	ts := timestamp()                                                                                         // 替换为你的存储空间名称
	objectKey := "translatedir/id/" + strconv.FormatInt(ts, 10) + "/image0.png"                               // OSS上的Object完整路径
	const localFilePath = "/Users/avey/Documents/Dev/translateimage/example2/downloadDir/09343436/image0.png" // 本地文件的完整路径

	// 使用设置的凭证创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error creating OSS client:", err)
		os.Exit(-1)
	}

	// 获取存储空间Bucket实例
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error getting bucket:", err)
		os.Exit(-1)
	}

	// 上传文件到OSS
	err = bucket.PutObjectFromFile(objectKey, localFilePath)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		os.Exit(-1)
	}

	// 打印API返回的信息
	fmt.Printf("File uploaded successfully.\n")
	// fmt.Printf("ETag: %s\n", result.ETag)
	// fmt.Printf("RequestId: %s\n", result.RequestId)
	// 构造可访问的 URL
	// 替换 <YourEndpoint> 为您的 OSS Endpoint
	// 如果您的 Endpoint 支持 HTTPS，请使用 https 协议
	url := fmt.Sprintf("http://"+bucketName+".oss-cn-hongkong.aliyuncs.com/%s", objectKey)
	fmt.Println("File URL:", url)
}

func timestamp() int64 {
	// 获取当前时间
	currentTime := time.Now()

	// 生成时间戳（秒级）
	timestamp := currentTime.Unix()

	fmt.Println("当前时间:", currentTime)
	fmt.Println("时间戳（秒级）:", timestamp)

	// // 如果你需要毫秒级的时间戳
	// timestampMillis := currentTime.UnixNano() / int64(time.Millisecond)
	// fmt.Println("时间戳（毫秒级）:", timestampMillis)
	return timestamp
}
