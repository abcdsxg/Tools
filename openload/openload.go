package openload

import (
	"flag"
	"log"
)

var (
	apiLogin = "你的apiLogin"
	apiKey   = "你的apikey"
	folder   = "文件夹id"
	from     = 0
	to       = -1
)

func GetGameOfThrones() {
	flag.StringVar(&apiLogin, "l", apiLogin, "你的apiLogin值")
	flag.StringVar(&apiKey, "k", apiKey, "你的apikey")
	flag.StringVar(&folder, "id", folder, "文件夹id")
	flag.IntVar(&from, "f", from, "从第几集开始")
	flag.IntVar(&to, "t", to, "到第几集结束")
	flag.Parse()
	if apiLogin == "你的apiLogin" || apiKey == "你的apikey" || folder == "文件夹id" {
		log.Print("你还没有改参数呢,例如go run main.go -l 你的apiLogin值 -k 你的apikey -id 文件夹id")
		return
	}
	postLinks, err := ParseSunkd()
	if err != nil {
		log.Print(err)
		return
	}
	urls, err := GetSunkdUrls(postLinks, from, to)
	if err != nil {
		log.Print(err)
		return
	}
	RemoteUpload(urls)
}
