/*
@Time   : 2022/3/9 17:15
@Author : ckx0709
@Remark :
*/
package util

import (
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	/*
	 *src 要加密的字符串
	 *key 用来加密的** **长度可以是128bit、192bit、256bit中的任意一个
	 *16位key对应128bit
	 */
	key := "iUa82<xua<2sf!(="

	deName := "刘魁申"
	deIdNo := "K90892987"
	src := "{\n" +
		"  \"name\":\"" + deName + "\"," +
		"  \"idNo\":\"" + deIdNo + "\"" +
		"}"

	crypted := AesEncrypt(src, key)
	encode := Base64UrlSafeEncode(crypted)
	fmt.Print(encode)
	//bytes := AesDecrypt(crypted, []byte(key))
	//decode, err := Base64URLDecode("39W7dWTd_SBOCM8UbnG6qA")

}

func TestAesEncrypt(t *testing.T) {
	key := "fu*#)shfs3fKeia@"
	str := "o13y6n0rnq-K-w7XFbH6u43-hKQZ_IpaLRhc1VL6g64bw3raITd0Wm_nqwqg5_-twns2FOWctFb7QJcWBtlzzZP6ZukP9q8M0weTfIUEwCwHLuvhdmf2oWjdc-DhqkchLYx3e_QuUGHZNeX7oEckX8l09EsCGUyQMuCJyYmk3HYsdDcyqu0-L5tCgXDK5IqKFEZcgxe7frn7rca4bcoZEsgDlHVV2WctcG2oG0I0l1aY25QzldfdGci_P_cY01Ra31ylJoSyNJuctlwY7VV2_bz0etXhGWSpm6eQtBJsPHVMGyUuuZy2mxY9zGW46NrBTJCz_7EGNMLXZW0ISqUx0cta2Te8wa7tA7Dd26QDzQlYzNopsqIcPUeXclaHu_Ye3o3w4zOppBm5miaY0rJBQfy4CQGQvBWJz9PUqtWhhY7Pm3n5gVQJVsWVLO7uGn2joWcsQ8UWMd2X9ZQIR3mjl4erkeSRLK73MBr1fhB7ZJzyRMpijXN3IaIJmp1wbGWtf8Dc5ZJotLHRYafB5CpMWx95B98yWk1Ko4sm_TQ1F3oKvkaSbXzGH7gAMJucEVKVNkDX6hyPNtf6UP7q66aWpySabYmmeTwFZpJfV1cGB15klHvCkuMFnZgcMvT5vl5GVzBQ1P7urAEgdeU9qIra6LgoXD2umwVnJP50Iv0zalyBDcZIxO-wpMGoCFZ9KtcWZhX6D5XEu3MCDgYnLzBHcutN--TMGgIsfAIrqZTY3jCklqeZW9eYViRIipkYftV2YVhBH-aAk7Erc9Ngvge0uf1wlUB2p58E18gcEByW908G6NHHJvdBK9vzsNFVsdxtVsqudnfYAkiXVKzX_m6VHw=="
	bytes, _ := Base64URLDecode(str)
	decrypt := AesDecrypt(bytes, []byte(key))
	fmt.Printf("****AesDecrypt: %s\n", decrypt)
}
