package db

import (
	"fmt"
	mydb "storage/db/mysql"
)

func init()  {
	fmt.Println("init???????????")
	TestQuery()
	fmt.Println("init end........")
}


func TestInsert(name string) bool {
	stmt, err := mydb.DBConn().Prepare("INSERT ignore INTO tb_t(`name`) VALUES(?)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(name)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	count, err := ret.RowsAffected()
	fmt.Println(count, err)
	fmt.Println("end=================")
	if err == nil && count == 1 {
		return true
	}
	return false
}

func TestQuery() {
	ret, err := GetFileMetaDB("bfc2ac5571d293ec5a5b9e611a332de82dcdcbae")
	fmt.Println(ret, "=========")
	if err != nil {
		fmt.Println("TestQuery: ", err)
	}
}


