package infra

import (
	"net/http"
	"io/ioutil"
)

type Receiver struct {

}

func (Receiver) Get (url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
