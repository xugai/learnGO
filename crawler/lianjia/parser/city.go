package parser

import (
	"fmt"
	"learnGO/crawler/engine"
	"regexp"
	"strings"
)

const content = `<a
      class="content__list--item--aside" target="_blank"      href="/zufang/AQ2464281991093428224.html"
      title="整租·北正街 1室1厅 南/北">
        <img
          alt="整租·北正街 1室1厅 南/北_北正街租房"
          src="https://s1.ljcdn.com/matrix_pc/dist/pc/src/resource/default/250-182.png?_v=20200507152818ad2"
          data-src="https://image1.ljcdn.com/340800-inspection/68b7fa46-97b5-4ab1-9246-2cb504536e7d.jpg!m_fill,w_250,h_182,l_flianjia_black,o_auto"
          class="lazyload">
        <!-- 是否展示vr图片 -->
                <!-- 广告标签 -->
            </a>
    <div class="content__list--item--main">
      <p class="content__list--item--title twoline">
        <a target="_blank" href="/zufang/AQ2464281991093428224.html">
          整租·北正街 1室1厅 南/北        </a>
      </p>
      <p class="content__list--item--des">
        <a target="_blank" href="/zufang/daguanqu/">大观区</a>-<a href="/zufang/longshanlu/" target="_blank">龙山路</a>-<a title="北正街" href="/zufang/c8827134158651696/" target="_blank">北正街</a>`


const mainHouseRege = `<a
      class="content__list--item--aside" target="_blank"      href="(/zufang/[0-9a-zA-Z]+\.html)"
      title="([^>]+)">`
const streetRege = `<p class="content__list--item--des">
        <a target="_blank" href="/zufang/[0-9a-zA-Z]+/">([^<]+)</a>-<a href="/zufang/[0-9a-zA-Z]+/" target="_blank">([^<]+)</a>-<a title="[^"]+" href="/zufang/[0-9a-zA-Z]+/" target="_blank">([^<]+)</a>
        <i>/</i>
        ([^<]+)
        <i>/</i>([^<]+)<i>/</i>
          ([^<]+)<span class="hide">
          <i>/</i>
          ([^<]+)
                  </span>
      </p>`
const priceRege = `<span class="content__list--item-price"><em>(-?[0-9]+)</em> 元/月</span>`

func ParseCity(content [] byte, prefixUrl string) engine.ParseResult {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Abnormal Error occured: ", err)
		}
	}()

	mainHouseR := regexp.MustCompile(mainHouseRege)
	streetR := regexp.MustCompile(streetRege)
	priceR := regexp.MustCompile(priceRege)
	mainHouseMatches := mainHouseR.FindAllSubmatch(content, -1)
	streetmatches := streetR.FindAllSubmatch(content, -1)
	priceMatches := priceR.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for i, m := range mainHouseMatches {
		houseName := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: prefixUrl + string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseHouse(bytes, houseName)
			},
		})

		result.Items = append(result.Items,
			string(m[2]) + " " + string(streetmatches[i][1]) + string(streetmatches[i][2]) + string(streetmatches[i][3]) + " " + string(streetmatches[i][4]) + " " +
			strings.Trim(string(streetmatches[i][5]), " ") + strings.Trim(string(streetmatches[i][6]), " ") + " " + strings.Join(strings.Fields(string(streetmatches[i][7])), "") +
			" " + string(priceMatches[i][1]) + "元/月")
	}
	return result
}
