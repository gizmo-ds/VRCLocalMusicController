package main

import (
	"github.com/gin-gonic/gin"
	kbe "github.com/micmonay/keybd_event"
)

var firstTrigger = make(map[string]bool)

func init() {
	gin.SetMode(gin.ReleaseMode)

	firstTrigger = map[string]bool{
		"prev": true,
		"play": true,
		"next": true,
	}
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
			if !firstTrigger["prev"] {
				kb.SetKeys(kbe.VK_MEDIA_PREV_TRACK)
			} else {
				firstTrigger["prev"] = false
			}
		case "play":
			if !firstTrigger["play"] {
				kb.SetKeys(kbe.VK_MEDIA_PLAY_PAUSE)
			} else {
				firstTrigger["play"] = false
			}
		case "next":
			if !firstTrigger["next"] {
				kb.SetKeys(kbe.VK_MEDIA_NEXT_TRACK)
			} else {
				firstTrigger["next"] = false
			}
		case "reset_first":
			firstTrigger = map[string]bool{
				"prev": true,
				"play": true,
				"next": true,
			}
		}
		kb.Launching()
		c.Status(404)
	})

	g.Run("0.0.0.0:58787")
}
