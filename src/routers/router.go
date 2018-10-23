package routers

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func init() {
	fmt.Println("Start Init func: ")
}

func main() {
	fmt.Println("Main config run: ")
	// Routing
	R := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

}

// r.Get("/", ha.StartPage)
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
