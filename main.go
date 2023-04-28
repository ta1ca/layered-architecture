package main

import (
	"awesomeProject/config"
	"awesomeProject/infrastructure/persistence"
	"awesomeProject/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	handler "awesomeProject/interface/handler"
)

func main() {
	// 依存関係を注入
	itemPersistence := persistence.NewItemPersistence(config.Connect())
	itemUseCase := usecase.NewItemUseCase(itemPersistence)
	itemHandler := handler.NewItemHandler(itemUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/items", itemHandler.Index)
	router.POST("/api/items", itemHandler.Create)
	router.PUT("/api/items", itemHandler.Update)

	// サーバ起動
	fmt.Println("------------------------")
	fmt.Println("サーバ起動 http://localhost:8080")
	fmt.Println("------------------------")

	http.ListenAndServe(":8080", &Server{router})
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET POST PUT")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}
