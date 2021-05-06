package main

import (
	"fmt"
	"net/http"
)

// 핸들러 구조체 정의 - 인스턴스르 만들고
type fooHandler struct{}

// fooHandler타입으로. - ServeHTTP라는 이름의 인터페이스를 정의
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo handler!")
}

func main() {

	/*
		net/http 모듈을 사용

		라우터 정의
		/ : 루트
		w : response를 writing, r: 사용자의 요청
	*/

	// 루트 핸들러 등록
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world go web server")
	})

	// bar 핸들러 추가
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello my bar!")
	})

	// func 타입으로 직접 등록하는 게 아니라 인스턴스 형태로 등록할 때는 Handle로!
	// 인스턴스를 등록
	http.Handle("/foo", &fooHandler{})

	// 3003 port로 server를 listen해준다.
	http.ListenAndServe(":3003", nil)
}
