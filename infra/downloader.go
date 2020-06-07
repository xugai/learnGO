package infra

import (
	//"learnGO/infra"
	//"./testing"
	"fmt"
)


func getReceiver() Retriever {
	return Receiver{}
}

type Retriever interface {
	Get(string) string
}

func main() {
	var r Retriever = getReceiver()
	fmt.Println(r.Get("http://www.imooc.com"))
}
