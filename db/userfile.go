package db

import (
	"fmt"
	mydb "storage/db/mysql"
	"time"
)

type UserFile struct {
	UserName string
	FileHash string
	FileName string
	FileSize int64
	UploadAt string
	LastUpdated string
}

func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tb_user_file(`user_name`, `file_sha1`, `file_name`, " +
			"`file_size`, `upload_at`)values(?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		return false
	}
	return true
}

func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_name, file_size, upload_at, last_update from tb_user_file where user_name = ? limit ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil {
		return nil, err
	}

	var userFiles []UserFile
	for rows.Next() {
		ufile := UserFile{}
		err = rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSize, &ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			fmt.Println("query user file failed" + err.Error())
			break
		}
		userFiles = append(userFiles, ufile)
	}
	return userFiles, nil
}


