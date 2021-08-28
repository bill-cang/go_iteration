/*
@Time   : 2021/8/25 16:36
@Author : ckx0709
@Remark :
encoding/json 文档有如下说明：
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*	var p *int
		b, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))*/

	mmp := map[string]interface{}{"name": "张三", "age": 18}
	bytes, err := json.Marshal(mmp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
