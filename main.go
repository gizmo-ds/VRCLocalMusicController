package main

import (
	"github.com/gin-gonic/gin"
	kbe "github.com/micmonay/keybd_event"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	kb, err := kbe.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	g := gin.Default()

	g.GET("/event", func(c *gin.Context) {
		e := c.Query("event")
		switch e {
		case "prev":
			kb.SetKeys(kbe.VK_MEDIA_PREV_TRACK)
		case "play":
			kb.SetKeys(kbe.VK_MEDIA_PLAY_PAUSE)
		case "next":
			kb.SetKeys(kbe.VK_MEDIA_NEXT_TRACK)
		}
		kb.Launching()
		c.Status(404)
	})

	g.Run("0.0.0.0:58787")
}
