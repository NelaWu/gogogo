package handler

import (
	"fmt"
	"gogogo/s3client"
	"io"
	"net/http"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "❌ 必須提供文件名", http.StatusBadRequest)
		return
	}

	//從s3獲取文件
	reader, contentType, err := s3client.DownloadFromS3(filename)
	if err != nil {
		http.Error(w, "❌ 下載失敗", http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	// 下載圖片，contentType=下載檔案類型
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", contentType)
	// 將文件內容複製到響應中
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, fmt.Sprintf("傳輸失敗: %v", err), http.StatusInternalServerError)
		return
	}

}
