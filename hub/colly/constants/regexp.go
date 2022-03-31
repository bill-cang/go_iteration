/*
@Time   : 2022/3/31 10:37
@Author : ckx0709
@Remark :
*/
package constants

import (
	"fmt"
	"regexp"
)

type Expr string

const (
	ExprDefault     Expr = "[\\s|\\S]+"
	ExprShareDayMsg Expr = "{[\\s|\\S]+}"
	ExprHtmlTable1  Expr = "</td><td>"
	ExprProxy1      Expr = "[\\d]{2,3}\\.[\\d]{2,3}\\.[\\d]{2,3}\\.[\\d]{2,3}:[\\d]{2,5}"
)

func (e Expr) GetRegexp() (rxp *regexp.Regexp) {
	rxp, _ = regexp.Compile(fmt.Sprintf("%s", e))
	if rxp == nil {
		rxp, _ = regexp.Compile(fmt.Sprintf("%s", ExprDefault))
	}
	return
}
