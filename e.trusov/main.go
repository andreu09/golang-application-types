package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	host string = "localhost"
	port int    = 8080

	html string = `<html>
		<body>
			<h1>
				Wow, you are genius!!!
			</h1>
			<img src="%v">
			</img>
		</body>
	</html>`

	cat = []string{
		0: "https://storage.azh.kz/archive/595b/23/9a/9e4fd9/1b/3c/8b4/fee/photos/1708144016.jpg/full.jpg",
		1: "https://legkovmeste.ru/wp-content/uploads/2019/02/post_5aa37458c8593.jpg",
	}
)

func main() {

	serverAddres := fmt.Sprintf("%v:%v", host, port)

	fmt.Println("Start Kitty server")
	http.HandleFunc("/pers", ShowKittyPers)
	http.HandleFunc("/siam", ShowKittySiam)
	http.HandleFunc("/random", ShowKittyRandom)
	fmt.Println(serverAddres)

	log.Fatal(http.ListenAndServe(serverAddres, nil))
}

func ShowKittyPers(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, cat[0])))
}
func ShowKittySiam(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, cat[1])))
}
func ShowKittyRandom(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, cat[rand.Intn(2)])))
}
