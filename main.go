package main

import (
	"fmt"
	"os"
	"time"
)

// timeはパッケージ名で、time.Timeで型を指定する

func main() {
	var args = os.Args
	var worktime time.Time

	switch len(args){
	case 1:
		fmt.Println("サブコマンドが必要です。'start'か'end'サブコマンドを入力してください")

	case 2:
		var subcommand string = args[1]

		// string型のswitchの時はダブルクォーテーションにする
		switch subcommand{
		case "start":
			worktime = time.Now()
			fmt.Println("Work Start!!!")

		case "end":
			worktime = time.Now()
			fmt.Println("Work End!!!")

		default:
			fmt.Println("'start'か'end'サブコマンドを入力してください")
		}

		fmt.Println(worktime)

	default:
		fmt.Println("workコマンドとサブコマンドstart/endのみです")
		return
	}
}