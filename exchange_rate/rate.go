package exchange_rate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

/**
获取人民币兑各种币的汇率
*/

const (
	baseUrl = "https://www.boc.cn/sourcedb/whpj/index"
)

var pattern = regexp.MustCompile(`<tr>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td class="pjrq">(.*?)</td>`)

type MoneyType struct {
	Target                    string
	BuyingRate                string
	CashPurchasePrice         string
	SellingRateOfSpotExchange string
	CashSellingRate           string
	BocConversionPrice        string
	Datetime                  string
}

func init() {
	log.Println("导入了exchange_rate ")
}

func GetRate(target string) float32 {
	html := getHtml()
	res := pattern.FindAllStringSubmatch(html,len(html))
	if len(res)>0{
		for _,v:= range res{
			log.Println(v[1],v[2],v[3])
		}
	}
	return 1.2
}
func GetAllRate()[]MoneyType{
	html := getHtml()
	res := pattern.FindAllStringSubmatch(html,len(html))
	list := make([]MoneyType,len(res))
	if len(res)>0{
		for k,v:= range res{
			list[k] = MoneyType{
				CashPurchasePrice:v[1],
				SellingRateOfSpotExchange:v[2],
				CashSellingRate:v[3],
				BocConversionPrice:v[4],
				Datetime:v[5],
			}
		}
	}
	return list
}
func getHtml() (html string) {
	method := "GET"
	url := ""
	client := &http.Client{
	}
	for i := 0; i <=10 ; i++ {
		if i==0{
			url = baseUrl+".html"
		}else{
			url = baseUrl +fmt.Sprintf("_%d.html",i)
		}
		log.Println(url)
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			log.Println(err)
		}
		res, err := client.Do(req)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		html += string(body)
	}


	return html
}
