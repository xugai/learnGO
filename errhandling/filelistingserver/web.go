package main

import (
	"learnGO/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"
	//"io/ioutil"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			r := recover()
			if r != nil {
				log.Println("Error occured: ", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if userErr, ok := err.(userError); ok {
			http.Error(writer, userErr.Message(), http.StatusInternalServerError)
			return
		}
		code := http.StatusOK
		if err != nil {
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.FileList()))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println("Server startup error!!!")
		panic(err)
	}

}
