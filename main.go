package main

import (
	"fmt"
	"net/http"
	"storage/handler"
)

func main()  {
	fmt.Println("start........")
	{
		http.HandleFunc("/file/upload", handler.UploadHandler)
		http.HandleFunc("/file/upload/suc", handler.UploadOkHandler)
		http.HandleFunc("/file/meta", handler.FileQueryHandler)
		http.HandleFunc("/file/download", handler.DownloadHandler)
		http.HandleFunc("/file/update", handler.FileUpdateMetaHandler)
		http.HandleFunc("/file/delete", handler.FileDelHandler)
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start failed: ", err.Error())
	}
}






