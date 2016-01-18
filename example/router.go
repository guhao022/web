package main
import (
	. "web"
	"log"
	"net/http"
	"github.com/num5/web"
)

func Erro(w http.ResponseWriter, r *http.Request) {
	e := &Error{
		ID: "001",
		Links: &ErrLinks{
			About: "http://baidu.com",
		},
		Status: 500,
		Code:   0001,
		Title:  "Title must not be empty",
		Detail: "Never occures in real life",
		Source: &ErrSource{
			Pointer: "#err",
		},
		Meta: map[string]interface{}{
			"creator": "api2go",
		},
	}

	err := NewError(w, e)

	if err != nil {
		log.Println(err)
	}

}

func main() {
	SetTrac(true)

	r := web.New()

	r.Get("/err", Erro)

	log.Printf("Server start listen on %d", 9900)

	err := http.ListenAndServe(":9900", r)

	if err != nil {
		panic(err)
	}
}
