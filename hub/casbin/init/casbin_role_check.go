/*
@Time   : 2022/3/8 15:44
@Author : ckx0709
@Remark : 管理平台接收更新信号，从测试环境获取swag信息（确保测试环境OK,且已更新swag），
检查系统接口变动，对casbin role表进行自动化检查更新
*/
package init

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//管理平台触发，service层手动调用
func CasRoleCheck() {
	//从测试环境拉去swag信息
	response, err := http.Get("https://community-test-api.epshealth.com:6443/swagger/doc.json")
	if err != nil {
		log.Printf("[CasRoleCheck] Get swag doc.json err: %+v", err)
		return
	}
	doc, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[CasRoleCheck] ReadAll response.Body err: %+v", err)
		return
	}
	defer response.Body.Close()

	rs := make(map[string]interface{})
	err = json.Unmarshal(doc, &rs)
	if err != nil {
		log.Printf("[CasRoleCheck] json.Unmarshal err: %+v", err)
		return
	}

	//从swag获取所有 route&action
	routersMap := make(map[string]string)
	for k2, v2 := range rs["paths"].(map[string]interface{}) {
		if strings.Index(k2, "{") > 0 {
			k2 = strings.Replace(k2, "{", ":", -1)
			k2 = strings.Replace(k2, "}", "", -1)
		}
		mmp := v2.(map[string]interface{})
		for k3, v3 := range mmp {
			if k3 == "delete" || k3 == "get" || k3 == "post" || k3 == "put" {
				mp := v3.(map[string]interface{})
				nk := fmt.Sprintf("%s,%s", k2, k3)
				routersMap[nk] = mp["description"].(string)
			}
		}
	}
	fmt.Printf("***routersMap keys :%+v", routersMap)

	//从casbin_role获取所有策略

	//比对最新swag与数据库policy，更新数据库策略&路由映射表

}
