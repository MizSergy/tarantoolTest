package database

import (
"encoding/json"
"fmt"
"github.com/go-redis/redis"
"log"
"os"
"time"
)

var ListRedisConnect []RedisClient
var ListRedis2Connect []RedisClient

type RedisClient struct {
	client *redis.Client
	Host string
	Port string
	Password string
}

type Redis2Client struct {
	client *redis.Client
	Host string
	Port string
	Password string
}

func RedisInit() {
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("REDIS_URL_%d", index))
		if len(host) == 0 {
			break
		}
		port := os.Getenv(fmt.Sprintf("REDIS_PORT_%d", index))
		pass := os.Getenv(fmt.Sprintf("REDIS_PASS_%d", index))

		ListRedisConnect = append(ListRedisConnect, RedisClient{
			Host: host,
			Port: port,
			Password:pass,
		})
		fmt.Printf("Connect redis %s:%s\n", host, port)
		index ++
	}
	fmt.Printf("Redis count connect %d\n", len(ListRedisConnect))
}

func Redis2Init() {
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("REDIS_URL_ALT_%d", index))
		if len(host) == 0 {
			break
		}
		port := os.Getenv(fmt.Sprintf("REDIS_PORT_ALT_%d", index))
		pass := os.Getenv(fmt.Sprintf("REDIS_PASS_ALT_%d", index))

		ListRedis2Connect = append(ListRedis2Connect, RedisClient{
			Host: host,
			Port: port,
			Password:pass,
		})
		fmt.Printf("Connect redis %s:%s\n", host, port)
		index ++
	}
	fmt.Printf("Redis count connect %d\n", len(ListRedisConnect))
}

func (r *RedisClient) GetQueueCollections(key string) string{
	val, _ := r.client.LPop(key).Result()
	return val
}

func (r *RedisClient) GetAllCollections(key string) []string{
	val, _ := r.client.LRange(key, 0 , -1).Result()
	return val
}

func (r *RedisClient) RemoveCollections(key string, number int){
	for i:=0; i < number; i++{
		r.client.LPop(key)
	}
}

func (r *RedisClient) SetSessionData(key string, value string) {
	r.client.Set("s-"+key, value, time.Hour)
}

func (r *RedisClient) GetSessionData(key string) string {
	result := r.client.Get("s-" + key)
	return result.Val()
}

func (r *RedisClient) SendCollectedData(item interface{}, flag string) {
	data, _ := json.Marshal(item)
	if err := r.client.RPush(flag, data).Err(); err != nil {
		fmt.Println(err.Error())
	}
}

func (r *RedisClient) Connect() *redis.Client {
	if r.client == nil{
		r.client = redis.NewClient(&redis.Options{
			Addr:     r.Host + ":" + r.Port,
			Password: r.Password,
			DB:      0,  // use default DB
		})
		_, err := r.client.Ping().Result()

		if err != nil {
			fmt.Printf("%s:%s - no connect\n", r.Host, r.Port)
			return nil
		}
	}
	return r.client
}

func (r *RedisClient) SetData(key string, value string) {
	r.client.Set("vc-"+key, value, time.Hour)
}

func (r *RedisClient) GetData(key string) string {
	result := r.client.Get("vc-" + key)
	return result.Val()
}

func (r *RedisClient) Disconnect()  {
	if r.client == nil{
		return
	}
	err := r.client.Close()
	if err != nil{
		fmt.Printf("%s:%s - %s\n", r.Host, r.Port, err.Error())
		return
	}
}

//==================================================================
var redis_client *redis.Client

type QueryData struct {
	Table string
	Fields map[string] string
}

func RedisConnect() *redis.Client {
	if redis_client == nil{
		redis_client = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASS"),
			DB:      0,  // use default DB
		})
		_, err := redis_client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}	}
	return redis_client
}

func GetQueueCollections(redis_cli *redis.Client, key string) string{
	val, _ := redis_cli.LPop(key).Result()
	return val
}




func RedisClose(){
	for _, client := range ListRedisConnect{
		client.Disconnect()
	}
	if redis_client != nil{
		_ = redis_client.Close()
	}
}

