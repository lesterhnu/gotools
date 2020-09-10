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
	baseUrl = "https://www.boc.cn/sourcedb/whpj/index.html"
)

var pattern = regexp.MustCompile(`(?s)<tr>[.\s]*<td>(.*?)</td>[.\s]*<td>(.*?)</td>[.\s]*<td>(.*?)</td>[.\s]*<td>(.*?)</td>[.\s]*<td>(.*?)</td>[.\s]*<td>(.*?)</td>[.\s]*<td class="pjrq">(.*?)</td>`)

type MoneyType struct {
	//目标币种
	Target string
	//现汇买入价
	BuyingRate string //
	//现钞买入价
	CashPurchasePrice string
	//现汇卖出价
	SellingRateOfSpotExchange string
	//现钞卖出价
	CashSellingRate string
	//中行折算价
	BocConversionPrice string
	//发布日期
	Datetime string
}

func GetRate(target string) MoneyType {
	targetPattern := regexp.MustCompile(fmt.Sprintf(`<tr>[.\n\s]*<td>%s</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td>(.*?)</td>[.\n\s]*<td class="pjrq">(.*?)</td>`, target))
	html := getHtml(baseUrl)
	res := targetPattern.FindAllStringSubmatch(html, len(html))
	var result MoneyType
	if len(res) > 0 {
		result = MoneyType{

			BuyingRate:                res[0][1],
			CashPurchasePrice:         res[0][2],
			SellingRateOfSpotExchange: res[0][3],
			CashSellingRate:           res[0][4],
			BocConversionPrice:        res[0][5],
			Datetime:                  res[0][6],
		}
	}
	return result

}
func GetAllRate() []MoneyType {
	html := getHtml(baseUrl)
	res := pattern.FindAllStringSubmatch(html, len(html))
	list := make([]MoneyType, len(res))
	if len(res) > 0 {
		for k, v := range res {
			list[k] = MoneyType{
				Target:                    v[1],
				BuyingRate:                v[2],
				CashPurchasePrice:         v[3],
				SellingRateOfSpotExchange: v[4],
				CashSellingRate:           v[5],
				BocConversionPrice:        v[6],
				Datetime:                  v[7],
			}
		}
	}
	return list
}
func getHtml(url string) (html string) {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	html = string(body)
	return
}
