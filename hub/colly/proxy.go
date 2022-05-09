/*
@Time   : 2022/3/31 12:19
@Author : ckx0709
@Remark :
*/
package main

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"go_iteration/hub/colly/client"
	"go_iteration/hub/colly/constants"
	"go_iteration/util"
	"sync"
)

var (
	ProxyListPool *arraylist.List
	proxySource   = []string{
		"https://proxy.seofangfa.com/",
		"https://ip.jiangxianli.com/?page=1",
		"https://ip.jiangxianli.com/?page=2",
	}
	ProxyInvalidMap = sync.Map{}
)

func init() {
	ProxyListPool = arraylist.New()

	collyClient := client.GetCollyClient()
	collyClient.OnResponse(func(rsp *colly.Response) {
		body := rsp.Body
		if body == nil {
			return
		}
		bodyStr := string(body)
		rxp0 := constants.ExprHtmlTable1.GetRegexp()
		tmp := rxp0.ReplaceAllString(bodyStr, ":")

		rxp := constants.ExprProxy1.GetRegexp()
		allString := rxp.FindAllString(tmp, -1)
		for _, str := range allString {
			str = "http://" + str
			ProxyListPool.Add(str)
		}
	})

	for _, sou := range proxySource {
		_ = collyClient.Visit(sou)
	}

	if ProxyListPool.Size() == 0 {
		log.Warnf("ProxyListPool is null !!!")
	}

}

//设置代理
func SetCollyProxy(c *colly.Collector) {
	if proxy, index := getRandProxy(); proxy != "" {
		//_ = c.SetProxy("http://182.34.206.193:25624")
		ProxyInvalidMap.Store(c.ID, index)
		_ = c.SetProxy(proxy)
	}
}

//移除无用的代理
func RemoveCollyProxy(c *colly.Collector) {
	if load, ok := ProxyInvalidMap.Load(c.ID); ok {
		index := load.(int)
		if pxy, ok := ProxyListPool.Get(index); ok {
			log.Warnf("[RemoveInvalidProxy] remove proxy %s.", pxy.(string))
			ProxyListPool.Remove(index)
		}
	}
}

func getRandProxy() (string, int) {
	if ln := ProxyListPool.Size(); ln > 0 {
		rand := util.GetRand(ln)
		pxy, _ := ProxyListPool.Get(rand)
		return pxy.(string), rand
	}
	log.Warnf("[getRandProxy] proxy is poor.")
	return "", 0
}
