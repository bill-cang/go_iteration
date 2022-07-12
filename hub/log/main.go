/*
@Time   : 2022/6/16 16:39
@Author : ckx0709
@Remark :
*/
package main

import log "github.com/sirupsen/logrus"

func main() {

	for i := 0; i < 20; i++ {
		log.Printf("%d伟大的领袖毛主席万岁", i)
	}

}
