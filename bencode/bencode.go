package bencode

import (
	//"fmt"
	"strconv"


)

func String(str string) string{
	res := strconv.Itoa(len(str))+":"+ str
	return res

}

func Int(i int) string{
	res := "i"+strconv.Itoa(i)+"e"
	return res
}

func List()