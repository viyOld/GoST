package routers

import (
	"fmt"
	"net/http"

	"../../src/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var R chi.Router

func init() {
	fmt.Println("Start Init Route: ")
	R = chi.NewRouter()

	R.Use(middleware.RequestID)
	R.Use(middleware.RealIP)
	R.Use(middleware.DefaultCompress)

	R.Get("/", handlers.StartPage)
	R.Get("/login", handlers.LoginPage)
	R.Post("/login", handlers.PostLogin)

	// static files
	R.Handle("/assets/images/*", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("./assets/images/"))))
	R.Handle("/assets/css/*", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./assets/css/"))))
	R.Handle("/assets/js/*", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("./assets/js/"))))
}

func main() {
	fmt.Println("Main config run: ")
	// Routing
	// R = chi.NewRouter()

	// R.Use(middleware.RequestID)
	// R.Use(middleware.RealIP)

	// R.Get("/", handlers.StartPage)
	// R.Get("/start", handlers.StartPage)

}

func Status() {
	fmt.Println("Status OK: ")
}

// //r.Get("/favicon.ico", ha.FaviconHandler)
// r.Get("/auth/login", ha.GetLogin)
// r.Post("/auth/login", ha.PostLogin)
// r.Get("/auth/sign-up", ha.GetSignUp)
// r.Post("/auth/sign-up", ha.PostSignUp)
// r.Get("/auth/refresh", ha.Refresh)
// r.Post("/profile", ha.SetProfile)
// r.Get("/profile", ha.GetProfile)
// r.Post("/profile/photo", ha.SetAvatar)
// r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

// r.Route("/lobby", func(r chi.Router) {
// 	//r.With(paginate).Get("/", listLobby)
// 	r.Get("/", ha.GetLobbyList)
// 	r.Post("/", ha.NewLobby)
// 	r.Get("/new_messages", ha.NewMessages)
// 	r.Route("/{lobbyID}", func(r chi.Router) {
// 		//r.Use(lobbyCtx)
// 		r.Get("/", ha.GetInvite)
// 		r.Get("/join", ha.JoinLobby)
// 		r.Post("/members", ha.AddMembers)
// 		r.Delete("/members/{mmbersID}", ha.DeleteMembers)
// 		r.Delete("/", ha.DeleteLobby)
// 		r.Get("/members", ha.GetMembers)
// 		r.Post("/logo", ha.SetLogo)
// 		r.Get("/messages", ha.GetMessages)
// 		r.Post("/message", ha.NewMessage)
// 		r.Post("/doc", ha.SendAttachment)
// 	})
// })
