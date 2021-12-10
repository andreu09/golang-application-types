package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
)

var (
	host string = "localhost"
	port int    = 8080
	html string = `<html>
	                    <body>
						    <h1> Wow, you are genius!!!</h1>
							<img src="%v"></img>
						</body>
					</html>`
	page string
	imgMap = map[int]string{
		0: "https://cdn.nur.kz/images/1120/b73ff3edc8a079fc.jpeg?version=1", 
		1: "https://crosti.ru/patterns/00/1d/fb/d2fc671697/picture.jpg",
	}	
)

func main() {
	serverAddres := fmt.Sprintf("%v:%v", host, port)
	fmt.Println("Start Kitty server")

	http.HandleFunc("/siam", SiamKitty)
	http.HandleFunc("/pers", PersKitty)
	http.HandleFunc("/random", RandomKitty)	

	fmt.Println(serverAddres)

	log.Fatal(http.ListenAndServe(serverAddres, nil))
}

func SiamKitty(w http.ResponseWriter, r *http.Request) {
	ShowKitty(w, r, int(0))
}

func PersKitty(w http.ResponseWriter, r *http.Request) {	
	ShowKitty(w, r, int(1))
}

func RandomKitty(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	rnKey := rand.Intn(2)
	ShowKitty(w, r, rnKey)
}

func ShowKitty(w http.ResponseWriter, r *http.Request, key int){
	page = fmt.Sprintf(html, imgMap[key])
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(page))
}
