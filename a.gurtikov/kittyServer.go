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

	page      string
	imagesMap = map[int]string{
		0: "https://img5.goodfon.ru/wallpaper/nbig/0/d3/koshka-kot-portret-lezhit-siamskaia-goluboglazaia-fon-morda.jpg",
		1: "https://murkoshka.ru/wp-content/uploads/kak-opredelit-porodu-koshek-i-kotov-19.jpg",
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
	kittyKey := rand.Intn(2)
	ShowKitty(w, r, kittyKey)
}

func ShowKitty(w http.ResponseWriter, r *http.Request, key int) {
	page = fmt.Sprintf(html, imagesMap[key])
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(page))
}
