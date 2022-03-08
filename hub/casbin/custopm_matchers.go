/*
@Time   : 2022/3/3 10:29
@Author : ckx0709
@Remark :
*/
package main

import (
	"errors"
	"fmt"
)

//桩函数
func RegxSub(args ...interface{}) (interface{}, error) {
	if err := validateVariadicArgs(2, args...); err != nil {
		return false, fmt.Errorf("%s: %s", "keyMatch2", err)
	}

	name1 := args[0].(string)
	name2 := args[1].(string)

	fmt.Println(name1, name2)
	return true, nil
}

func validateVariadicArgs(expectedLen int, args ...interface{}) error {
	if len(args) != expectedLen {
		return fmt.Errorf("Expected %d arguments, but got %d", expectedLen, len(args))
	}

	for _, p := range args {
		_, ok := p.(string)
		if !ok {
			return errors.New("Argument must be a string")
		}
	}

	return nil
}
