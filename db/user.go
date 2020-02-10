package db

import (
	"fmt"
	mydb "storage/db/mysql"
)

func UserSignUp(username string, pwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"INSERT ignore INTO tb_user(`user_name`, `user_pwd`) VALUES(?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, pwd)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	count, err := ret.RowsAffected()
	fmt.Println(count, err, username, pwd)
	if err == nil && count == 1 {
		return true
	}
	return false
}

func QueryPwd(username string) (string, error) {
	stmt, err := mydb.DBConn().Prepare("SELECT `user_pwd` FROM tb_user WHERE user_name = ?")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return "", err
	}
	defer stmt.Close()

	var pwd string
	err = stmt.QueryRow(username).Scan(&pwd)
	if err != nil{
		fmt.Println(username, err.Error())
		return "", err
	}
	return pwd, nil
}

func SaveToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare("REPLACE INTO tb_token(`user_name`, `token`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil{
		return false
	}
	return true
}

type User struct {
	Username string
	Email   string
	Phone string
	SignupAt string
	LastActiveAt string
	Status int
}


func QueryUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mydb.DBConn().Prepare(
		"SELECT `user_name`, `signup_at` FROM tb_user WHERE user_name=? limit 1")
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ", err.Error())
		return user, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		fmt.Println("faield query userinfo", err.Error())
		return user, err
	}
	return user, nil
}

