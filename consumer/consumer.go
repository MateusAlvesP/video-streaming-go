package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"gocv.io/x/gocv"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	window := gocv.NewWindow("Video Stream")
	defer window.Close()

	msgs, err := ch.Consume("video_frames", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		var imgBytes []byte
		buf := bytes.NewBuffer(msg.Body)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(&imgBytes); err != nil {
			fmt.Println("Failed to decode frame:", err)
			continue
		}

		img, err := gocv.NewMatFromBytes(480, 640, gocv.MatTypeCV8UC3, imgBytes)
		if err != nil {
			log.Println("Failed to create mat from bytes:", err)
			continue
		}

		window.IMShow(img)
		if gocv.WaitKey(1) >= 0 {
			break
		}
		img.Close()
	}
}
