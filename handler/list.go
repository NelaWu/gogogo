package handler

import (
	"fmt"
	"gogogo/s3client"
	"net/http"
)

func ListObjectsHandler(w http.ResponseWriter, r *http.Request) {
	result, err := s3client.ListObjects()
	if err != nil {
		http.Error(w, fmt.Sprintf("列出文件失敗: %v", err), http.StatusInternalServerError)
		return
	}

	// 將輸出寫入 HTTP 響應
	for _, obj := range result.Contents {
		fmt.Fprintf(w, "文件: %s , 大小: %d\n", *obj.Key, *obj.Size)
	}
	fmt.Fprintf(w, "✅ 列出成功！\n")
}
