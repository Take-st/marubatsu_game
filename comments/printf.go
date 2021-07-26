package comments

func Hoge() string{
	return "hoge"
}

import (
	"fmt"
)

func whereToPutCommet() {
	fmt.Println("どこに置きますか？")
	fmt.Println("行列で入力してください")
	fmt.Println("(ex) 12 =>1行目2列目")
}