package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"storage/db"
	"storage/meta"
	"storage/util"
	"strconv"
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
		fmt.Println(fileMeta.FileSha1, "||||||")

		meta.SaveFileMeta(fileMeta)

		r.ParseForm()
		username := r.Form.Get("username")
		ok := db.OnUserFileUploadFinished(username, fileMeta.FileSha1,
			fileMeta.FileName, fileMeta.FileSize)
		if ok {
			http.Redirect(w, r, "/file/upload", http.StatusFound)
		} else {
			w.Write([]byte("Upload Failed."))
		}
	}
}

func UploadOkHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Ok!")
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

func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	username := r.Form.Get("username")
	userFiles, err := db.QueryUserFileMetas(username, limitCnt)
	if err != nil {
		fmt.Println("query user file failed", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(userFiles)
	if err != nil {
		fmt.Println("query user file json decode failed", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func TryFastUploadHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filename := r.Form.Get("filename")
	filesize, _ := strconv.Atoi(r.Form.Get("filesize"))

	fileMeta, err := meta.GetFileMeta(filehash)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	if fileMeta.FileName == "" {
		resp := util.RespMsg{
			Code: -1,
			Msg: "秒传失败",
		}
		w.Write(resp.JSONBytes())
		return
	}
	ok := db.OnUserFileUploadFinished(username, filehash, filename, int64(filesize))
	if ok {
		resp := util.RespMsg{
			Code: 0,
			Msg: "秒传成功",
		}
		w.Write(resp.JSONBytes())
	} else {
		resp := util.RespMsg{
			Code: 2,
			Msg: "上传失败，请稍后重试",
		}
		w.Write(resp.JSONBytes())
	}

}



func FileQuery1Handler(w http.ResponseWriter, r *http.Request) {
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
