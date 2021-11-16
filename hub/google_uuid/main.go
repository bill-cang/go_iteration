/*
@Time   : 2021/11/8 15:42
@Author : ckx0709
@Remark : google uuid
*/
package main

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func main() {
	//t := time.Unix(1000000, 0)
	t := int64(time.Now().Nanosecond())
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t)), 0)

	for i := 0; i < 10; i++ {
		fmt.Println(ulid.MustNew(ulid.Timestamp(time.Now()), entropy))
	}

}

