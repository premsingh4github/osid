package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/zserge/webview"
	"log"
	"net/url"
)

func main() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("OS Id:" + id)
	id_new, err_new := machineid.ProtectedID("myAppName")
	if err_new != nil {
		log.Fatal(err)
	}
	fmt.Println("Secure Id: " + id_new)

	const myHTML =  "<!doctype html><html><body>test html"+id_new+"</body></html>"
	//`<!doctype html><html><body>test html</body></html>`
	w := webview.New(webview.Settings{
		URL: `data:text/html,` + url.PathEscape(myHTML),
		Resizable:true,
	})
	w.Run();

}
