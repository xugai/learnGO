package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

const url = "https://www.lianjia.com/city/"
const resultSize = 119

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		log.Printf("Fetch url %s error: %v", url, err)
	}

	expectedUrl := [] string {
		"https://bj.lianjia.com/zufang/",
		"https://sh.lianjia.com/zufang/",
		"https://sz.lianjia.com/zufang/",
	}

	expectedCity := [] string {
		"北京", "上海", "深圳",
	}

	// verify parse result
	result := ParseCityList(content)
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d requests", resultSize, len(result.Items))
	}

	// verify url
	for i, url := range expectedUrl {
		if result.Requests[i].Url != url {
			t.Errorf("expected the #%d url is: %s, but get %s", i, url, result.Requests[i].Url)
		}
	}

	// verify city
	for i, city := range expectedCity {
		if result.Items[i].(string) != city {
			t.Errorf("expected the #%d city is: %s, but get %s", i, city, result.Items[i].(string))
		}
	}
}
