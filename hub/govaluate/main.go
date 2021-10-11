/*
@Time   : 2021/10/11 11:12
@Author : ckx0709
@Remark : https://darjun.github.io/2020/04/01/godailylib/govaluate/
*/

package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{})
	parameters["foo"] = -1
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	parameters = make(map[string]interface{})
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	parameters = make(map[string]interface{})
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}
