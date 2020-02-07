package db

import (
	"database/sql"
	"fmt"
	mydb "storage/db/mysql"
)

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FilePath sql.NullString
}


func SaveFileMetaDB(filehash string, filename string, filesize int64, filepath string) bool {
	// 使用Prepared statement可以防注入、拼接更快
	stmt, err := mydb.DBConn().Prepare(
		"INSERT ignore INTO tb_file(`sha1`, `name`, `size`, `path`, `status`) VALUES (?, ?, ?, ?, 1)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()
	fmt.Println("filename", filename)
	ret, err := stmt.Exec(filehash, filename, filesize, filepath)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if count, err := ret.RowsAffected(); err == nil {
		if count <= 0 {
			fmt.Printf("File<%s> with hash: %s has uploaded\n", filename, filehash)
		}
		return true
	}
	return false
}

func UpdateFileMetaDB(filehash string, filename string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"UPDATE tb_file SET name = ? WHERE sha1 = ?")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if count, err := ret.RowsAffected(); err == nil {
		if count <= 0 {
			fmt.Printf("File with hash: %s has uploaded\n", filehash)
		}
		return true
	}
	return false
}

func GetFileMetaDB(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"SELECT `sha1`, `name`, `size`, `path` FROM tb_file WHERE sha1 = ? AND status = 1 limit 1")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileName, &tfile.FileSize, &tfile.FilePath)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}

func DelFileMetaDB(filehash string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"UPDATE tb_file SET status = 0 WHERE sha1 = ?")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if count, err := ret.RowsAffected(); err == nil {
		if count <= 0 {
			fmt.Printf("File with hash: %s has uploaded\n", filehash)
		}
		return true
	}
	return false
}
