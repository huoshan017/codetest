package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        []string{":6379", ":6380", ":6381", ":6382", ":6383", ":6384"},
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	})

	ctx := context.TODO()

	key := "aaa"
	value := 111
	status := rdb.Set(ctx, key, value, 10*time.Second)
	if status.Err() != nil {
		log.Fatalf("redis set key %v with value %v error: %v", key, value, status.Err())
		return
	}

	strCmd := rdb.Get(ctx, key)
	if strCmd.Err() != nil {
		log.Fatalf("redis get key %v error: %v", key, strCmd.Err())
		return
	}

	var err error
	var val string
	val, err = strCmd.Result()
	if err != nil {
		log.Fatalf("string cmd result error: %v", err)
		return
	}

	log.Printf("redis get key %v value %v", key, val)

	var set bool
	set, err = rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("set ", set)
	}

	set, err = rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("set ", set)
	}

	var vals []string
	vals, err = rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("vals ", vals)
	}

	var z []redis.Z
	z, err = rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: 0, Count: 2}).Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("z ", z)
	}

	var v int64
	v, err = rdb.ZInterStore(ctx, "out", &redis.ZStore{Keys: []string{"zset1", "zset2"}, Weights: []float64{2, 3}}).Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("v ", v)
	}

	var i interface{}
	i, err = rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("i ", i)
	}
}
