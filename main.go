package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"gogogo/handler"
	"gogogo/s3client"
)

func main() {
	// 讀取設定
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("設定檔讀取失敗: %v", err)
	}

	// 初始化 S3
	s3client.InitS3()

	// 路由
	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.UploadHandler).Methods("POST")
	r.HandleFunc("/list", handler.ListObjectsHandler).Methods("GET")
	r.HandleFunc("/download", handler.DownloadHandler).Methods("GET")
	r.HandleFunc("/delete/{filename}", handler.DeleteHandler).Methods("DELETE")

	log.Println("✅ Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
