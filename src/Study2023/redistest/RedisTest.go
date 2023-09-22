package main

import (
"context"
"fmt"
"github.com/go-redis/redis/v8"
)

/**
I.下载
go get github.com/go-redis/redis/v8

 */
func main() {
fmt.Println("开始")
	RedisTest1()
}

func RedisTest1()  {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}


/**
Cluster集群
 */
func RedisTest2()  {
	ctx := context.Background()

	rdb := rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"cc.db.test:16380", "cc.db.test:16380", "cc.db.test:16380", "cc.db.test:16380", "cc.db.test:16380", "cc.db.test:16380"},
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}