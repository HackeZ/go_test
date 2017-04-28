package main

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Product struct {
	Stock int32 `redis:"stock"`
	Sales int32 `redis:"sales"`
}

const (
	redisLocation = "127.0.0.1:6379"

	ERRORSTRING = "NULL"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisLocation)
			if err != nil {
				log.Fatal("Dail master redis server error:", err)
			}
			return c, err
		},
	}
}

func main() {
	//var result []interface{}
	//result = append(result, []byte{57, 57, 57}, []byte{48})
	//fmt.Println(result)
	//
	//var p = Product{}
	//_, err := redis.Scan(result, &p.Stock, &p.Sales)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(p)
	//fmt.Println("type =>", reflect.TypeOf(p.Stock))

	conn := pool.Get()
	defer conn.Close()

	// HGETSTR
	hResult := HGetStr("test1", "field1")
	log.Println("HGETSTR  result => ", hResult)

	// HMGET
	hmResult1 := HMGet("test1", "field1", "field2")
	log.Println("hmResult1 =>", hmResult1)

	hmResult2, _ := HGetAll("test1")
	log.Println("hmResult2 =>", hmResult2)

	hmResult3 := redis.

}

// 获取指定key的里面为field对应的值
func HGetStr(key string, field string) string {
	conn := pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		log.Fatal("redis HGetStr失败,key:"+key+", field: "+field, " err:", err)
		return ERRORSTRING
	}
	return value
}

// 获取指定key的里面为field对应的值
func HMGet(key string, fields ...interface{}) map[string]interface{} {
	conn := pool.Get()
	defer conn.Close()
	values, err := redis.Values(conn.Do("HMGET", redis.Args{}.Add(key).Add(fields...)...))
	if err != nil {
		log.Fatal("redis HMGet失败,key:", key, " err:", err)
		return nil
	}
	return values
}

func HGetAll(key string) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}
