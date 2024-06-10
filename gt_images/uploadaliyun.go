package main

import (
	"fmt"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Aliyunoss(accessKeyId, accessKeySecret, endpoint, bucketName, objectKey, localFilePath string) (string, error) {

	// 使用设置的凭证创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Print("Error creating OSS client:", err)
	}

	// 获取存储空间Bucket实例
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Print("Error getting bucket:", err)
	}

	// objectKey := "translatedir/" + objectKeyID + "/" + strconv.FormatInt(timestamps(), 10) + objectKeyimagename
	// 上传文件到OSS
	err = bucket.PutObjectFromFile(objectKey, localFilePath)
	if err != nil {
		log.Print("Error uploading file:", err)

	}

	// 打印API返回的信息
	fmt.Printf("File uploaded successfully.\n")
	// fmt.Printf("ETag: %s\n", result.ETag)
	// fmt.Printf("RequestId: %s\n", result.RequestId)
	// 构造可访问的 URL
	// 替换 <YourEndpoint> 为您的 OSS Endpoint
	// 如果您的 Endpoint 支持 HTTPS，请使用 https 协议
	url := fmt.Sprintf("https://"+bucketName+"."+endpoint+"/%s", objectKey)
	log.Print("File URL:", url)
	return url, err
}
