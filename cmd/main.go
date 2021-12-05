package main

import (
	"message-board/api"
	"message-board/dao"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		return
	}
	api.InitEngine()
}
