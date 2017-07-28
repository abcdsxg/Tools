package openload

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func ParseSunkd() ([]string, error) {
	url := "http://moviesunkd.net/%E6%AC%8A%E5%8A%9B%E7%9A%84%E9%81%8A%E6%88%B2-game-thrones-4/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Print("url出错")
		return nil, errors.New("url出错")
	}
	props := doc.Find(`div.entry>p`)
	if props.Length() == 0 {
		log.Print("没有找到")
		return nil, errors.New("没有找到")
	}
	var links []string
	props.Each(func(i int, prop *goquery.Selection) {
		prop.Find(`a`).Each(func(j int, ss *goquery.Selection) {
			result, _ := ss.Attr("href")
			if strings.Contains(result, "moviesunkd.net/?p") {
				links = append(links, result)
			}
		})

	})
	return links, nil
}

var wg sync.WaitGroup
var rw sync.Mutex
var urls []string

func GetSunkdUrls(postUrls []string) ([]string, error) {
	wg.Add(len(postUrls))
	for index, url := range postUrls {
		//go parseUrl(index, url)
		parseUrl(index, url)
	}
	wg.Wait()
	return urls, nil
}

func parseUrl(i int, u string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Printf("第%d条网址:%s没有匹配到", i, u)
	}
	req.Header.Set("Cookie", "__cfduid=dd5d66240b675cfdd8e6554b09f43bd901500699460;"+
		" wordpress_test_cookie=WP+Cookie+check;"+
		" wordpress_logged_in_5e9088018b968adb67db841f267c0900=mm11%7C1532602490%7Cn5mu5Lc5i0kpeh5GDSzknTQ7kGcnZ0ZnbRGqqfOimOX%7C5d5969731471c8e18f2910ec7725e4ac5b165ecce4d0baa4a906a4825153b41d;"+
		" _ga=GA1.2.209954273.1500699462; _gid=GA1.2.434515712.1501066384;"+
		" innity.crtg.728_90=IN2p120%2CIN2p100%2CIN2p080%2CIN2p060%2CIN2p040%2C;"+
		" innity.crtg.300_250=IN1p120%2CIN1p100%2CIN1p080%2CIN1p060%2CIN1p040%2C;"+
		" __AF=58ee5ee1-4d27-4f09-a52d-fdb4473f17e4")

	res, err := client.Do(req)
	d, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Printf("url:%s访问出错", u)
	}
	id := u[strings.Index(u, "p=")+2:]
	props := d.Find(`div.wordpress-post-tabs>#tabs_` + id + `_0>#tabs-` + id + `-0-0`)
	if props.Length() == 0 {
		log.Printf("第%d条网址:%s没有匹配到", i, u)
	}
	h, _ := props.Html()
	r := regexp.MustCompile(`src="(.*?)"`)
	f := r.FindStringSubmatch(h)
	if len(f) > 1 {
		rw.Lock()
		urls = append(urls, f[1])
		log.Printf("成功获取到第%d个:%s", i+1, u)
		rw.Unlock()
	} else {
		log.Printf("第%d条网址:%s没有匹配到", i, u)
	}
	wg.Done()
}

func RemoteUpload(urls []string) {
	for i, url := range urls {
		u := fmt.Sprintf("https://api.openload.co/1/remotedl/add?login=%s&key=%s&url=%s&folder=%s&headers=",
			apiLogin, apiKey, url, folder)
		r, err := http.Get(u)
		if err != nil {
			log.Printf("url:%s访问出错", u)
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		if strings.Contains(string(body), "\"status\":200") {
			log.Printf("第" + strconv.Itoa(i+1) + "个url:" + url + "上传成功")
			log.Print(string(body))
		} else {
			log.Printf("第" + strconv.Itoa(i+1) + "个url:" + url + "上传失败,err:" + string(body))
		}

	}
}

type FileInfo struct {
	//{"status":200,"msg":"OK","result":{"id":"76723200","folderid":"3910806"}}
	Status string
	Msg    string
	Result struct {
		Id       string
		folderid string
	}
}
