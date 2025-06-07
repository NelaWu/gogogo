package handler

import (
	"fmt"
	"net/http"

	"gogogo/s3client"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "❌ 讀取檔案失敗", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = s3client.UploadToS3(header.Filename, file, header.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ 上傳失敗: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "✅ 上傳成功！%s", header.Filename)
}
