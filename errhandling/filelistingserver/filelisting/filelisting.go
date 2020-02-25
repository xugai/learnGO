package filelisting

import (
	"net/http"
	"os"
	"log"
	"io/ioutil"
	"strings"
	//"errors"
)

const prefix = "/list/"

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func FileList() func(writer http.ResponseWriter, request *http.Request) error {
	return func(writer http.ResponseWriter, request *http.Request) error{
		if strings.Index(request.URL.Path, prefix) == -1 {
			log.Println("Request URL path must start with " + prefix)
			return userError("Request URL path must start with " + prefix)
		}
		fileName := request.URL.Path[len(prefix):]
		file, err := os.Open(fileName)
		if err != nil {
			log.Println("Error: ", err.Error())
			return err
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("Error: ", err.Error())
			return err
		}
		writer.Write(content)
		return nil
	}
}