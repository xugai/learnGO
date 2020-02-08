package real

import (
	"net/http"
	"io/ioutil"
	"time"
)

type Retriever struct {
	Content string
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) Get (url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	r.Content = string(bytes)
	return r.Content
}
