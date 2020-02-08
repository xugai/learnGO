package testing

import (
	//"net/http"
	//"io/ioutil"
)

type Retriever struct {

}

func (Retriever) Get (url string) string {
	return "fake content"
}
