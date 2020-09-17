package main

import "fmt"

func thirdPartyConnectDB() {
	// 自分で例外を投げる
	panic("unable to connect database!")
}

func save() {
	// 上にいるとpanicで落ちる
	// thirdPartyConnectDB()

	defer func() {
		// panicをrecoverでキャッチ（強制終了しないようにする）
		s := recover()
		fmt.Println(s)
	}()

	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
