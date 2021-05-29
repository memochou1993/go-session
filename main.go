package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
)

var (
	key        = []byte(os.Getenv("SESSION_KEY"))
	store      = sessions.NewCookieStore(key)
	cookieName = "auth"
)

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "Secret")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, cookieName)

	session.Values["authenticated"] = false
	session.Save(r, w)
}
