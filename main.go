package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
	"regexp"
	"zdamtosam.pl/src/backend"
	"zdamtosam.pl/src/db"
	"zdamtosam.pl/src/frontend"
)

var dbClient = db.NewDatabaseClient()
var authClient = db.NewAuthClient()
var api = backend.NewHandler(dbClient)
var front = frontend.NewHandler(dbClient, authClient)

func init() {
	functions.HTTP("Entrypoint", mainHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	firstPath := regexp.MustCompile("/?([^/]*)/?")
	group := firstPath.FindStringSubmatch(r.URL.Path)
	if len(group) > 1 {
		switch group[1] {
		case "search":
			front.Search(w, r)
			break
		case "login":
			front.Login(w, r)
			break
		case "index.html":
			front.Handle(w, r)
			break
		case "privacy-policy":
			front.PrivacyPolicy(w, r)
			break
		case "terms-of-service":
			front.TermsOfService(w, r)
			break
		case "ads.txt":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "manifest.json":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "robots.txt":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "img":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "css":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "js":
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			break
		case "sitemap.xml":
			api.Sitemap(w, r)
			break
		case "api":
			api.Handle(w, r)
			break
		default:
			front.Handle(w, r)
			break
		}
		return
	}

	front.Handle(w, r)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":80", nil)
}
