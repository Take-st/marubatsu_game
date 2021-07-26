package main

import (
	"fmt"
	"strings"
	// "marubatsu_game/comments"
)

func whereToPutCommet() {
	fmt.Println("どこに置きますか？")
	fmt.Println("行列で入力してください")
	fmt.Println("(ex) 12 =>1行目2列目")
}

func finishComment(side string) {
	fmt.Printf("%sの勝ちです！\n", side)
}

func isWinnerHorizontal(side string, table [][]string) (isWin bool) {
	isWin = false
	length := len(table)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if table[i][j] != side {
				isWin = false
				break
			}
			isWin = true
		}
		if isWin == true {
			break
		}
	}
	return
}

func isWinnerVertical(side string, table [][]string) (isWin bool) {
	isWin = false
	length := len(table)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if table[j][i] != side {
				isWin = false
				break
			}
			isWin = true
		}
		if isWin == true {
			break
		}
	}
	return
}

func isWinnerDiagonalRight(side string, table [][]string) (isWin bool) {
	isWin = false
	length := len(table)
	for i := 0; i < length; i++ {
		if table[i][i] != side {
			isWin = false
			break
		}
		isWin = true
	}
	return
}

func isWinnerDiagonalLeft(side string, table [][]string) (isWin bool) {
	isWin = false
	length := len(table)
	for i, j := 0, length-1; i < length && j >= 0; i, j = i+1, j-1 {
		if table[i][j] != side {
			isWin = false
			break
		}
		isWin = true
	}
	return
}

func isWinner(side string, table [][]string) bool {
	return isWinnerHorizontal(side, table) || isWinnerVertical(side, table) || isWinnerDiagonalRight(side, table) || isWinnerDiagonalLeft(side, table)
}

func splitDigit(i int) (k, l int) {
	if i >= 10 || i <= 99 {
		l = i % 10
		k = (i / 10) % 10
	} else {
		fmt.Println("いたずらしないでね")
	}
	return
}

func putStone(side string, table [][]string) [][]string {
	whereToPutCommet()
	var a, c, r int
	fmt.Scan(&a)
	c, r = splitDigit(a)
	table[c-1][r-1] = side
	return table
}

func printTable(table [][]string) {
	fmt.Println("")
	fmt.Println("現在の盤面")
	fmt.Println("*************")
	fmt.Println("  1 2 3")
	for i := 0; i < len(table); i++ {
		fmt.Printf("%d %s\n", i+1, strings.Join(table[i], " "))
	}
	fmt.Println("*************")
}

func switchTurn(turn string) string {
	if turn == "○" {
		return "×"
	} else {
		return "○"
	}
}

func main() {

	// fmt.Println(comments.Hoge())

	// type threeArray [3][3]string
	// var x threeArray = [3]string{"xx", "yy", "zz"}
	// fmt.Println(x)

	// var table threeArray = [3][3]string{
	// 	{"-", "-", "-"},
	// 	{"-", "-", "-"},
	// 	{"-", "-", "-"},
	// }
	// fmt.Println(table)
	var table = [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	var winner string
	turn := "○"

	fmt.Println("先攻: ○, 後攻:×です。")
	printTable(table)

	for !(isWinner("○", table) || isWinner("×", table)) {
		fmt.Printf("%sのターンです\n", turn)
		table = putStone(turn, table)
		printTable(table)
		turn = switchTurn(turn)
	}

	if isWinner("○", table) {
		winner = "○"
	} else if isWinner("×", table) {
		winner = "×"
	}

	finishComment(winner)

	// for i, v := range []string{"foo", "bar", "baz"} {
	// 	fmt.Println(i, v)
	// }

	// arg := []string{"a", "b", "c"}
	// for i, v := range arg {
	// 	fmt.Println(i, v)
	// }

}
