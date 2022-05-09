/*
@Time   : 2022/3/31 12:05
@Author : ckx0709
@Remark :
*/
package client

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

func GetCollyClient() (collector *colly.Collector) {
	//终端伪装
	//userAgent := colly.UserAgent(constants.GetRandUserAgent())
	collector = colly.NewCollector()
	//伪装终端
	extensions.RandomUserAgent(collector)
	//伪装请求头部
	//extensions.Referer(collector)
	return collector
}
