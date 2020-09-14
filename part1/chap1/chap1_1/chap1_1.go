//chap1_1.go
package main

import (
	"fmt"
	"github.com/techauthor/gobook/part1/chap1/chap1_1/demo"
	"net/http"
	"strconv"
)

//index
func HelloGo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Golang")
}

//create user
func CreateUser(writer http.ResponseWriter, request *http.Request) {
	u := &demo.User{
		UserName: "xhtian", Pwd: "123456", Email: "txhui@xxx.com", Tel: "13888888888",
	}
	demo.DB.InsertOne(u)
	fmt.Fprintln(writer, "create user.", u)
}

//clear all users
func DeleteAll(writer http.ResponseWriter, request *http.Request) {
	sql := "delete from `t_user`"
	demo.DB.Exec(sql)
	fmt.Fprintln(writer, "delete user all.")
}

//delete user
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id, _ := strconv.ParseInt(request.Form["id"][0], 10, 64)
	var u demo.User
	demo.DB.Id(id).Delete(&u)
	fmt.Fprintln(writer, "delete user id is:", id)
}

//update user
func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id, _ := strconv.ParseInt(request.Form["id"][0], 10, 64)
	name := request.Form["name"][0]
	var u = &demo.User{
		Id:       id,
		UserName: name,
	}
	demo.DB.Id(id).Update(u)
	fmt.Fprintln(writer, "update user.", u)
}

//the main method
func main() {
	fmt.Println("hello chap1_1...")
	http.HandleFunc("/gobook/chap1_1", HelloGo)
	http.HandleFunc("/gobook/chap1_1/create", CreateUser)
	http.HandleFunc("/gobook/chap1_1/delete", DeleteUser)
	http.HandleFunc("/gobook/chap1_1/update", UpdateUser)
	http.HandleFunc("/gobook/chap1_1/deleteall", DeleteAll)
	http.ListenAndServe(":8888", nil)
}
