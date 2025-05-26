package lib

import "main/lib/routes"

func init() {
	Server.
		WithRequestHandler("GET /todos", routes.GetTodos).
		WithRequestHandler("POST /todos", routes.PostTodos).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /events", routes.GetEvents).
		WithRequestHandler("GET /welcome", routes.GetWelcome).
		WithRequestHandler("GET /", routes.GetAny)
}
