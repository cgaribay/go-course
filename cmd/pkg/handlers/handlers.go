package handlers

import (
	"net/http"

	"github.com/cgaribay/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
	/*_, err := fmt.Fprintf(w, "This is the home page")
	if err != nil {
		fmt.Println(err)
	}*/
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
	/*_, err := fmt.Fprintf(w, "This is the about page")
	if err != nil {
		fmt.Println(err)
	}*/
}
