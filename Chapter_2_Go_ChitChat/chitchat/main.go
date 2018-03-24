package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux() 		 																	//  표준 라이브러리에서 멀티 플렉서 생성
	files := http.FileServer(http.Dir(config.Static)) 							// public에 있는 존재하는 파일을 핸들러가 제공할 수 있도록  files에 저장
	mux.Handle("/static/", http.StripPrefix("/static/", files) 			// StripPrefix함수는 요청 url에서 첫번째 매개변수를 url에서 제거해줌.
	/*
	   위 두줄은 localhost/static/css/first.css 라고 요청이 오면 <application root>/public/css/first.css 파일을 제공해 줌.
	*/

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)

/*  url과 핸들러를 연결해 주기 위해 핸들 함수 이용
	위 함수는 루트로 요청이 들어오면 index 핸들러로 요청을 넘긴다.
	핸들러 함수로 매개변수를 제공하지 않는 이유는 모든 핸들러 함수는 ResponseWriter를 첫 번째 매개변수로 사용하고,
	Rrequest에 대한 포인터를 두번째 매개변수로 사요하므로 처리기 함수에 매개변수를 제공하지 않아도 됨.
	*/
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
