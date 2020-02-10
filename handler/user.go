package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"storage/db"
	"storage/util"
)



func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			io.WriteString(w, "internal server error"+err.Error())
			return
		}
		io.WriteString(w, string(data))
		return
	}

	r.ParseForm()

	name := r.Form.Get("username")
	pwd := r.Form.Get("password")

	encPasswd := util.Sha1(pwd)
	ok := db.UserSignUp(name, encPasswd)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SUCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}


func SigninHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	pwd := r.Form.Get("password")

	// 1.check pwd
	encPwd, err := db.QueryPwd(username)
	if err != nil {
		w.Write([]byte("LOGIN QUERY FAILED"))
		return
	}
	if !util.CheckPwd(encPwd, pwd) {
		w.Write([]byte("LOGIN FAILED"))
		return
	}
	// 2.genr token
	token := util.GenToken(username, pwd)

	// 3. save token
	if db.SaveToken(username, token) {
		resp := util.RespMsg{
			Code: 0,
			Msg: "OK",
			Data: struct{
				Location string
				Username string
				Token string
			} {
				Location: "http://192.168.19.132:8687/static/view/home.html",
				Username: username,
				Token: token,
			},
		}
		w.Write(resp.JSONBytes())
	} else {
		w.Write([]byte("sign in failed"))
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")

	// check token
	//token := r.Form.Get("token")

	// 2. 查询用户信息
	user, err := db.QueryUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 3.组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}
