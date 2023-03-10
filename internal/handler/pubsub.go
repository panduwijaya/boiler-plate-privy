// Package cakes
// Automatic generated
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"cake-store/cake-store/internal/appctx"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var cfg = appctx.NewConfig()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     cfg.Redis.Hosts,
	Password: cfg.Redis.Password, // no password set
	DB:       cfg.Redis.DB,       // use default DB
})

func LISTEN() {
	list_channel := strings.Split(cfg.Redis.Channel, ",")
	for _, value := range list_channel {
		fmt.Println("Redis Subscribe Connected : ", value)
		go subscribe(value)
	}
}

func subscribe(channel ...string) {
	subscriber := redisClient.Subscribe(ctx, channel...)

	var data interface{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Printf("error format")
		}

		if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
			fmt.Printf("error format")
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", data)
	}
}

func Publish(ctx context.Context, param interface{}) {
	fmt.Println(param)
	payload, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	if err := redisClient.Publish(ctx, "listen-callback-payment", payload).Err(); err != nil {
		panic(err)
	}
}
