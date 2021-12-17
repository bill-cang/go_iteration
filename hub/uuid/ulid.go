/*
@Time   : 2021/11/8 15:42
@Author : ckx0709
@Remark : google uuid
*/
package uuid

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

var (
	randSource = int64(time.Now().Nanosecond())
	entropy    = ulid.Monotonic(rand.New(rand.NewSource(randSource)), 0)
)

func GetOrderlyUUID() (uuid string) {
	mustNew := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return mustNew.String()
}
