package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/binding"
	"github.com/nanobox-io/golang-scribble"
	"log"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email string `form:"email"`
}

func main() {

	m := martini.Classic()

	db, err := scribble.New("./users", nil)

	if err != nil {
		log.Println("Error in db")
	}

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset: "UTF-8",
	}))

	m.Use(martini.Static("static"))

	m.Get("/", func(r render.Render) {
		content, err := db.ReadAll("users")
		if (err != nil) {
			log.Println("Can not read from db")
		}
		fmt.Println(content)
		r.HTML(http.StatusOK, "index", nil)
	})
	m.Post("/", binding.Bind(User{}), func (r render.Render, u User){
		db.Write("users", u.Username, u)
		r.Redirect("/", 302)
	})


	m.Run()

}
