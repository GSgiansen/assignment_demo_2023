package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisClient struct {
	cli *redis.Client
}

func (c *RedisClient) InitClient(ctx context.Context, address, password string) error {
	//Initialise a new redis client, no password set
	red := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set for ease
		DB:       0,
	})

	// test connection to the redis server
	if err := red.Ping(ctx).Err(); err != nil {
		return err
	}

	c.cli = red
	return nil
}

func (c *RedisClient) SaveMessage(ctx context.Context, roomID string, message *Message) error {
	// converts the wanted message into json format for sending
	text, e := json.Marshal(message)
	if e != nil {
		msg := fmt.Sprintf("error saving message, err: %v", e)
		log.Fatal(msg)
		return e
	}

	//setting member to be part of redis datastore set
	member := &redis.Z{
		Score:  float64(message.Timestamp),
		Member: text,
	}

	_, e = c.cli.ZAdd(ctx, roomID, *member).Result()
	if e != nil {
		return e
	}

	return nil
}

func (c *RedisClient) RoomIDMessages(ctx context.Context, roomID string, start, end int64, reverse bool) ([]*Message, error) {
	var (
		stringMessages []string
		//messages obtained from the redis server are in String format
		messages []*Message
		err      error
	)

	if reverse {
		stringMessages, err = c.cli.ZRevRange(ctx, roomID, start, end).Result()
		if err != nil {
			return nil, err
		}
	} else {
		// Asc order with time -> first message is the earliest message
		stringMessages, err = c.cli.ZRange(ctx, roomID, start, end).Result()
		if err != nil {
			return nil, err
		}
	}

	//unpacks the messages into JSON before converting them to Messages struct
	for index, msg := range stringMessages {
		print(index)
		temp := &Message{}
		err := json.Unmarshal([]byte(msg), temp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, temp)
	}

	return messages, nil
}
