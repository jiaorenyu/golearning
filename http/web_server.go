package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, str)
	value := r.FormValue("test")
	fmt.Println(value)
}

func (st Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, st.Greeting)
	fmt.Fprint(w, st.Punct)
	fmt.Fprint(w, st.Who)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Error")
	}

	defer c.Close()

	reply, err := redis.String(c.Do("get", "a"))

	//fmt.Println(reply)

	//fmt.Println(r.FormValue("key"))
	//fmt.Println(r.UserAgent())
	fmt.Fprint(w, reply)
}

func main() {
	http.HandleFunc("/get", getHandler)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
