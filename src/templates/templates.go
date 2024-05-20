package templates

import (
	"html/template"
)

var PublicHome *template.Template
var PublicResume *template.Template
var PublicContact *template.Template
var AuthLogin *template.Template
var PrivateHome *template.Template

func init() {
	PublicHome = template.Must(template.ParseFiles("templates/util/base.tmpl", "templates/util/navbar.tmpl", "templates/public/home.tmpl"))
	PublicResume = template.Must(template.ParseFiles("templates/util/base.tmpl", "templates/util/navbar.tmpl", "templates/public/resume.tmpl"))
	PublicContact = template.Must(template.ParseFiles("templates/util/base.tmpl", "templates/util/navbar.tmpl", "templates/public/contact.tmpl"))
	AuthLogin = template.Must(template.ParseFiles("templates/util/base.tmpl", "templates/util/navbar.tmpl", "templates/auth/login.tmpl"))
	PrivateHome = template.Must(template.ParseFiles("templates/util/base.tmpl", "templates/util/navbar.tmpl", "templates/private/home.tmpl"))
}