package main

import (
	"github.com/garyburd/redigo/redis"
	l "log"
)

func main() {
	l.Println("golang with redis")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
		l.Panic("cannot dial")
	}
	defer c.Close()
	l.Println("connection open")

	out, err := redis.Values(c.Do("GET", "key01"))
	if err != nil {
		l.Println(err)
		return
	}

	l.Println("out", out)

}
