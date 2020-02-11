package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

/* 接口列表
init 初始化分块
[in]:
file_hash
file_size
user_name
[out]:
chunk_count
chunk_size
file_size
file_hash
uploadid

uppart 上传分块
[in]:
username
uploadid
index

complete 分块上传完成
[in]
username
file_hash
file_size
file_name

cancel 取消上传
[in]
file_hash

status 查询上传状态
[in]
file_hash
*/

type MPUploadInfo struct {
	ChunkSize int
	ChunkCount int
	FileHash   string
	FileSize int
	UploadId string
}


func MPUploadInitHandler(w http.ResponseWriter, r http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")
	filesize, _ := strconv.ParseFloat(r.Form.Get("file_size"), 32)
	username := r.Form.Get("username")

	uploadid := username + fmt.Sprintf("%x", time.Now().UnixNano())
	chunkCount := int(math.Ceil(filesize / 1024 / 1024))
	fmt.Println(filehash, uploadid, chunkCount)

}


