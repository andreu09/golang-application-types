package main

import (
	"fmt"
	"log"
	"net/http"
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

	page string
	img  string = "https://previews.123rf.com/images/linelly/linelly1803/linelly180300124/98380361-lovely-kitten-portrait-cute-kitty-kitten.jpg"
)

func main() {
	page = fmt.Sprintf(html, img)
	serverAddres := fmt.Sprintf("%v:%v", host, port)

	fmt.Println("Start Kitty server")
	http.HandleFunc("/show", ShowKitty)
	fmt.Println(serverAddres)

	log.Fatal(http.ListenAndServe(serverAddres, nil))
}

func ShowKitty(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	fmt.Println(user)
	w.Write([]byte(page))
}
