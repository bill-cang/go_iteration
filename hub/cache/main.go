/*
@Time   : 2022/7/15 14:02
@Author : ckx0709
@Remark :
*/
package main

import (
	"github.com/patrickmn/go-cache"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(cache.DefaultExpiration, 0)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

}
