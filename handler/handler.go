package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"storage/meta"
	"storage/util"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error"+err.Error())
			return
		}
		io.WriteString(w, string(data))

	} else if r.Method == "POST" {
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println("get form file failed" + err.Error())
			return
		}

		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: header.Filename,
			Location: "/tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create("/tmp/" + header.Filename)
		if err != nil {
			fmt.Println("create file failed: " + err.Error())
			return
		}
		defer newFile.Close()

		if fileMeta.FileSize, err = io.Copy(newFile, file); err != nil{
			fmt.Println("copy file failed: " + err.Error())
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload", http.StatusFound)
	}
}

func UploadOkHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Ok!")
}

func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0]
	fMeta, err := meta.GetFileMeta(fileHash)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(data))

}


func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")
	fMeta, err := meta.GetFileMeta(fsha1)
		if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	file, err := os.Open(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=\"%s\"", fMeta.FileName))
	w.Write(data)
}


func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fsha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	fMeta, err := meta.GetFileMeta(fsha1)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fMeta.FileName = newFileName
	meta.UpdateFileMeta(fMeta)

	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func FileDelHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fsha1 := r.Form["filehash"][0]
	meta.RemoveFileMeta(fsha1)

	fMeta, err := meta.GetFileMeta(fsha1)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	os.Remove(fMeta.Location)

	w.WriteHeader(http.StatusOK)
}
