package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/*
*
I.下载
go get github.com/go-redis/redis/v8
*/
func main() {
	fmt.Println("开始")
	RedisTestSentinel1()
}

func RedisTestSentinel1() {
	ctx := context.Background()
	// 创建 Redis Sentinel 客户端
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master6380",                                                         // 替换为实际的主节点名称
		SentinelAddrs: []string{"cc.db.test:16380", "aa.db.test:16380", "bb.db.test:16380"}, // 替换为实际的 Sentinel 地址和端口号
		Password:      "密码",
		DB:            12,
	})
	// 测试连接是否成功
	err := client.Ping(ctx).Err()
	if err != nil {
		fmt.Println("无法连接到Redis Sentinel:", err)
		return
	}
	fmt.Println("与Redis Sentinel连接成功")
	err = client.Set(ctx, "name", "Tom", 0).Err()
	if err != nil {
		fmt.Println("设置键值对失败:", err)
		return
	} else {
		fmt.Println("set 成功")

	}

	val, err := client.Get(ctx, "name").Result()
	if err == redis.Nil {
		fmt.Println("键不存在")
	} else if err != nil {
		fmt.Println("无法获取值:", err)
	} else {
		fmt.Println("get 成功 name 的值是:", val)
	}
}
