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

	img = []string{
		"https://avatars.mds.yandex.net/i?id=3d1ed2534042761b8be4e786a3ce876d-5593341-images-thumbs&n=13",
		"https://99px.ru/sstorage/56/2013/05/image_563005130022185879681.jpg",
	}
)

func main() {

	serverAddres := fmt.Sprintf("%v:%v", host, port)

	fmt.Println("Start Kitty server")
	http.HandleFunc("/siam", ShowSiamKitty)
	http.HandleFunc("/pers", ShowPersKitty)
	http.HandleFunc("/random", ShowRandomKitty)
	fmt.Println(serverAddres)

	log.Fatal(http.ListenAndServe(serverAddres, nil))
}

// Siam Cats
func ShowSiamKitty(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, img[0])))
}

// Pers Cats
func ShowPersKitty(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(fmt.Sprintf(html, img[1])))
}

// Random Cats
func ShowRandomKitty(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	rand.Seed(time.Now().UnixNano())
	w.Write([]byte(fmt.Sprintf(html, img[rand.Intn(len(img))])))
}
