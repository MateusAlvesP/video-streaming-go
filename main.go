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
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

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

	q, err := ch.QueueDeclare("video_frames", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Device closed")
			return
		}

		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(img.ToBytes()); err != nil {
			log.Println("Failed to encode frame:", err)
			continue
		}

		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        buf.Bytes(),
		})
		if err != nil {
			log.Println("Failed to publish frame:", err)
		}
	}
}
