package main

import (
	"fmt"
	"./real"
	"./mock"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download (r Retriever) string {
	return r.Get(url)
}

func post (p Poster) {
	p.Post(url, map[string]string{
		"course": "golang",
		"teacher": "ccmouse",
	})
}

func session (s RetrieverPoster) string{
	s.Post(url, map[string]string{
		"content": "This is a fake session",
	})
	return s.Get(url)
}


// type assertion 1
func inspect(r Retriever) {
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Printf("%v\n", v.Content)
	case *real.Retriever:
		fmt.Printf("%v\n", v.UserAgent)
	}
}

func main() {

	var r Retriever

	r = &mock.Retriever{"This is a fake Retriever"}
	fmt.Printf("{r.Content:  %s}\n", r)
	inspect(r)

	r = &real.Retriever{
		"content",
		"Google Chrome 5.0",
		time.Minute,
	}
	inspect(r)
	// type assertion 2
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Printf("%v\n", realRetriever.Content)
	} else {
		fmt.Printf("r's type is not mock.Retriever")
	}

	fmt.Println("Try session:")
	retrieverPoster := &mock.Retriever{}
	fmt.Println(session(retrieverPoster))
	//fmt.Println(download(r))
}