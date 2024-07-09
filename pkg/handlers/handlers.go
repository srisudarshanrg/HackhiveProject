package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

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
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
		CustomErrors: nil,
	})
}

func (a *HandlerAccess) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	detail_entered := r.Form.Get("username_entered")
	password_entered := r.Form.Get("password_entered")

	searchDetailQuery := `select * from login_details where username=$1 or email=$2 or phone=$3`
	result, err := db.Exec(searchDetailQuery, detail_entered, detail_entered, detail_entered)

	if err != nil {
		log.Println(err)
	}

	rowsAffected, _ := result.RowsAffected()

	errorMap := map[string]interface{}{}
	var errorString string

	if rowsAffected == 0 {
		errorString = "Username or Password incorrect"
		errorMap["notFound"] = errorString
		render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
			CustomErrors: errorMap,
		})
	} else {
		var hashed_password string

		confirmPasswordQuery := `select password from login_details where username=$1 or email=$2 or phone=$3`
		row, err := db.Query(confirmPasswordQuery, detail_entered, detail_entered, detail_entered)
		if err != nil {
			log.Println(err)
		}

		defer row.Close()
		for row.Next() {
			err := row.Scan(&hashed_password)
			if err != nil {
				panic(err)
			}
		}

		check := GetPasswordFromHash(password_entered, hashed_password)

		if check {
			log.Println("Login successful.")
		} else {
			errorString = "Username or Password incorrect"
			errorMap["notFound"] = errorString
			render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{
				CustomErrors: errorMap,
			})
		}
	}
}

func (a *HandlerAccess) SignUp(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "sign-up.page.tmpl", &models.TemplateData{
		CustomErrors: nil,
	})
}

func (a *HandlerAccess) PostSignUp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse form")
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")

	searchUniqueQueryUsername := `select username from login_details where username = $1 or email = $2 or phone = $3`
	result, err := db.Exec(searchUniqueQueryUsername, username, email, phone)

	if err != nil {
		log.Println(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		hashed_password, err := HashPassword(password)

		if err != nil {
			log.Fatal(err)
		}

		addRowQuery := `insert into login_details (username, password, email, phone) values($1, $2, $3, $4)`
		_, err = db.Exec(addRowQuery, username, hashed_password, email, phone)
		if err != nil {
			log.Println(err)
		}

		driver.DisplayRows(db)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		errorMap := map[string]interface{}{}

		errorText := "This user already exists. Choose another one."
		errorMap["uniqueDetail"] = errorText

		render.RenderTemplate(w, r, "sign-up.page.tmpl", &models.TemplateData{
			CustomErrors: errorMap,
		})
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hashedPassword), nil
}

func GetPasswordFromHash(entered_password string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(entered_password))
	return err == nil
}
