package conf

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func init(){
	c := cron.New()

	c.Start()
	_, err := c.AddFunc("*/5 * * * * *", func() {

	})
	if err != nil {
		fmt.Println("错误")
		return
	}
}