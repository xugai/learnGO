package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"learnGO/crawler_distributed/config"
	"net/http"
	"time"
)

func determineEncoding(r * bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Errorf("Peek byte error: %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func getUTF8Reader(r io.Reader, t transform.Transformer) * transform.Reader {
	newReader := transform.NewReader(r, t)
	return newReader
}

var rateLimiter = time.Tick(time.Second / config.QPS)
func Fetch(url string) ([] byte, error){
	<- rateLimiter
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:77.0) Gecko/20190101 Firefox/77.0")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("Request " + url + "error.")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong request, status code is %d", response.StatusCode)
	}

	bodyReader := bufio.NewReader(response.Body)
	utf8Reader := getUTF8Reader(bodyReader, determineEncoding(bodyReader).NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
