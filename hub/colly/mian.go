/*
@Time   : 2022/3/22 15:20
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go_iteration/hub/colly/client"
	"go_iteration/hub/colly/constants"
	"net/http"
	"time"
)

func main() {

	c := client.GetCollyClient()
	c.SetRequestTimeout(10 * time.Second)
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		log.Debugf("***OnHTML client = %+v", *c)
	})

	c.OnResponse(getShareDayMsg)

	c.OnRequest(func(r *colly.Request) {
		//设置代理
		SetCollyProxy(c)
	})

	addr := `http://push2.eastmoney.com/api/qt/stock/get?ut=fa5fd1943c7b386f172d6893dbfba10b&invt=2&fltt=2&fields=f43,f57,f58,f169,f170,f46,f44,f51,f168,f47,f164,f163,f116,f60,f45,f52,f50,f48,f167,f117,f71,f161,f49,f530,f135,f136,f137,f138,f139,f141,f142,f144,f145,f147,f148,f140,f143,f146,f149,f55,f62,f162,f92,f173,f104,f105,f84,f85,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f107,f111,f86,f177,f78,f110,f262,f263,f264,f267,f268,f250,f251,f252,f253,f254,f255,f256,f257,f258,f266,f269,f270,f271,f273,f274,f275,f127,f199,f128,f193,f196,f194,f195,f197,f80,f280,f281,f282,f284,f285,f286,f287,f292&cb=jQuery112407430974032649418_1642642459884&_=%d&secid=1.600353`
	addr = fmt.Sprintf(addr, time.Now().UnixMilli())
	for {
		err := c.Visit(addr)
		if err == nil {
			break
		} else {
			log.Errorf("Visit err =%+v\n", err)
			//移除无用的代理
			RemoveCollyProxy(c)
		}
		<-time.After(2 * time.Second)
		addr = fmt.Sprintf(addr, time.Now().UnixMilli())
	}

}

func getShareDayMsg(rsp *colly.Response) {
	log.Debugf("***OnHTML client = %+v", rsp)
	if rsp.StatusCode != http.StatusOK {
		//TODO：添加上下文追踪
		log.Errorf("[getShareDayMsg] colly.Response.StatusCode :%d", rsp.StatusCode)
		return
	}

	//fmt.Printf("[getShareDayMsg] rsp.Body =%s", rsp.Body)
	rxp := constants.ExprShareDayMsg.GetRegexp()
	findString := rxp.FindString(fmt.Sprintf("%s", rsp.Body))
	fmt.Printf("***:%s\n", findString)

}
