package handler

import (
	"github.com/gorilla/mux"
	"gogogo/s3client"
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	if filename == "" {
		http.Error(w, "文件名不能為空", http.StatusBadRequest)
		return
	}
	err := s3client.DeleterFromS3(filename)
	if err != nil {
		log.Printf("刪除文件失敗: %v", err)
		http.Error(w, "刪除文件失敗", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("文件刪除成功"))

}
