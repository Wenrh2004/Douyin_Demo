package main

import (
	action "Douyin_Demo/kitex_gen/douyin/publish/action/douyinpublishactionservice"
	"log"
)

func main() {
	svr := action.NewServer(new(DouyinPublishActionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
