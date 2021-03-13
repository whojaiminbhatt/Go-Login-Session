package accountcontroller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("uname")
	password := request.Form.Get("psw")
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)
	if username == "root" && password == "12345" {
		session, _ := store.Get(request, "mysession")
		session.Values["username"] = username
		session.Save(request, response)
		http.Redirect(response, request, "/account/welcome", http.StatusSeeOther)

	} else {
		data := map[string]interface{}{
			"err": "Invalid Username Or Password",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
		tmp.Execute(response, data)

	}

}

func Welcome(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	username := session.Values["username"]
	fmt.Println("username", username)
	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("views/accountcontroller/welcome.html")
	tmp.Execute(response, data)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/account/index", http.StatusSeeOther)
}
