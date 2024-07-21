package main

import "context"
import "fmt"
import "net/http"
import "os"
import "os/signal"
import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var routes []Route

type MyHandler struct{}

func (f *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Request URL:", r.URL.Path)
	for _, route := range routes {
		if route.Method == r.Method && route.Path == r.URL.Path {
			route.Handler(w, r)
			return
		}
	}
	http.NotFound(w,r)
}

func main() {
	routes = append(routes, []Route{
		{Method: "GET", Path: "/hello", Handler: func(w http.ResponseWriter, r *http.Request) {
			user := &User{
				Name: "Tom",
				Age:  10,
			}
			userBytes, _ := json.Marshal(user)
			_, err := w.Write(userBytes)
			if err != nil {
				panic(err)
			}
		}},
		{Method: "POST", Path: "/world", Handler: func(w http.ResponseWriter, r *http.Request) {
			resp := &Resp{
				Code: 200,
				Msg:  "Success",
				Data: nil,
			}
			respBytes, _ := json.Marshal(resp)
			_, err := w.Write(respBytes)
			if err != nil {
				panic(err)
			}
		}},
	}...)

	fmt.Println("welcome!")
	server := &http.Server{
		Addr:    ":9092",
		Handler: &MyHandler{},
	}
	go func() {
		fmt.Println("listening...")
		err := server.ListenAndServe()

		if err != nil {
			fmt.Println("here comes a panic")
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("server shutdown")
	server.Shutdown(context.Background())

}
