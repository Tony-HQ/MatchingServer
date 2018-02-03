package main

import pb "./models"
import "fmt"

func test() bool {
	gamer := pb.Gamer{"aaa", "bbb", 99, 1}
	fmt.Println(gamer.String())
	return true
}

func main() {
	test()
}
