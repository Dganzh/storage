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
		http.HandleFunc("/file/fastupload", handler.TryFastUploadHandler)
		http.HandleFunc("/file/upload/suc", handler.UploadOkHandler)
		http.HandleFunc("/file/meta", handler.FileQueryHandler)
		http.HandleFunc("/file/download", handler.DownloadHandler)
		http.HandleFunc("/file/update", handler.FileUpdateMetaHandler)
		http.HandleFunc("/file/delete", handler.FileDelHandler)
		http.HandleFunc("/file/query", handler.FileQueryHandler)
	}

	{
		http.HandleFunc("/user/signup", handler.SignUpHandler)
		http.HandleFunc("/user/signin", handler.SigninHandler)
		http.HandleFunc("/user/info", handler.UserInfoHandler)
	}

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("start failed: ", err.Error())
	}
}






