package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/srisudarshanrg/HackhiveProject/pkg/config"
	"github.com/srisudarshanrg/HackhiveProject/pkg/driver"
	"github.com/srisudarshanrg/HackhiveProject/pkg/models"
	"github.com/srisudarshanrg/HackhiveProject/pkg/render"
)

var Repo HandlerAccess
var db *sql.DB

type HandlerAccess struct {
	App *config.AppConfig
}

func SetAppConfigHandler(a *config.AppConfig) *HandlerAccess {
	return &HandlerAccess{
		App: a,
	}
}

func NewHandlers(r *HandlerAccess) {
	Repo = *r
}

func DatabaseAccess(database *sql.DB) {
	db = database
}

func (a *HandlerAccess) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) PostLogin(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func (a *HandlerAccess) SignUp(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "sign-up.page.tmpl", &models.TemplateData{})
}

func (a *HandlerAccess) PostSignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse form")
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	addRowQuery := `insert into login_details (username, password) values($1, $2)`
	_, err = db.Exec(addRowQuery, username, password)
	if err != nil {
		log.Println(err)
	}

	driver.DisplayRows(db)
}
