package main

import "net/http"

// Routeur pour l'ensemble de nos routes

func router() http.Handler {
	mux := http.NewServeMux()
	// Routes de test
	mux.HandleFunc("GET /health", HealthHandle)

	// Routes techniques
	mux.HandleFunc("GET /cpu", CPUHandler)
	mux.HandleFunc("GET /ps", PSHandler)
	mux.HandleFunc("GET /ps/{user}", PSUserHandler)
	mux.HandleFunc("GET /net", NetHandler)
	mux.HandleFunc("GET /net/{card}", NetNameHandler)
	mux.HandleFunc("GET /mem", MemHandler)
	mux.HandleFunc("GET /disk", DiskHandler)
	mux.HandleFunc("GET /load", LoadHandler)

	// Autres cas : fichiers statiques
	fs := http.FileServer(http.Dir("www"))
	mux.Handle("/", fs)

	return mux
}
