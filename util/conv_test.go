/*
@Time   : 2021/10/22 9:42
@Author : ckx0709
@Remark : 转换
*/
package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestConv2DEC(t *testing.T) {
	dec := Conv2DEC("1011")
	fmt.Print(dec)
}

func TestConv2DEC2(t *testing.T) {
	/*	fmt.Printf("%.2f\n", 1.479+0.005)
		fmt.Printf("%.2f\n", 1.475+0.005)
		fmt.Printf("%.2f\n", 1.470+0.005)*/

	//fmt.Printf("%.2f\n", FloatPositionRetain(18, 2))
	fmt.Println(FloatPositionRetain(1.4799, 2, 0))
	fmt.Println(FloatPositionRetain(1.4755, 2, 0))
	fmt.Println(FloatPositionRetain(1.4701, 2, 0))

	fmt.Println("****")
	fmt.Println(FloatPositionRetain(1.4799, 2, 1))
	fmt.Println(FloatPositionRetain(1.4755, 2, 1))
	fmt.Println(FloatPositionRetain(1.4710, 2, 1))
	fmt.Println(FloatPositionRetain(1.4701, 2, 1))

}

func TestConv2DEC4(t *testing.T) {
	holidays := GetYearHolidays("2021-03-23")
	fmt.Println(holidays)
}

func TestConv2DEC5(t *testing.T) {
	times := GetTimes(10)
	fmt.Println(times)
}

func TestFileTouch(t *testing.T) {
	unescape, err := url.QueryUnescape("http://92.push2his.eastmoney.com/api/qt/stock/kline/get?cb=jQuery112408165682304265207_1641795511100&secid=0.002625&ut=fa5fd1943c7b386f172d6893dbfba10b&fields1=f1%2Cf2%2Cf3%2Cf4%2Cf5%2Cf6&fields2=f51%2Cf52%2Cf53%2Cf54%2Cf55%2Cf56%2Cf57%2Cf58%2Cf59%2Cf60%2Cf61&klt=101&fqt=0&end=20500101&lmt=120&_=1641795511166")
	if err != nil {
		return
	}

	fmt.Println(unescape)

}

func TestConv2DEC6(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://99.push2.eastmoney.com/api/qt/stock/sse?ut=fa5fd1943c7b386f172d6893dbfba10b&fltt=2&fields=f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f206,f207,f208,f209,f210,f211,f212,f213,f214,f215,f86,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f292&secid=0.002625", nil)
	request.Header.Add("Accept", "text/event-stream")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Host", "91.push2.eastmoney.com")
	request.Header.Add("Origin", "http://quote.eastmoney.com")
	request.Header.Add("Referer", "http://quote.eastmoney.com/")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36 Edg/97.0.1072.55")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(all))


}
