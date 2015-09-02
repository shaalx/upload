package test

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego/httplib"
	"testing"
)

func TestT(t *testing.T) {

	b, err := ioutil.ReadFile("C 学习笔记.pdf")
	if checkerr(err) {
		return
	}
	// req := httplib.Put("http://localhost:8080/upload/tests/upload_test.go")
	req := httplib.Put("http://upload.daoapp.io/upload/qyuhen/C 学习笔记.pdf")
	req.Body(b)
	resp, err := req.Response()
	if checkerr(err) {
		return
	}
	b2, err := ioutil.ReadAll(resp.Body)
	if checkerr(err) {
		return
	}
	fmt.Println(string(b2))
}

func checkerr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
