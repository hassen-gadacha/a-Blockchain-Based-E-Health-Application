package main

import (
	"Encryptionlogic"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Proxy server starting on 8084 port ...")

	db, err := sql.Open("mysql", "root:houssem@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/saveUser", Encryptionlogic.SaveUser(db)).Methods("POST")
	r.HandleFunc("/DecryptCipher", Encryptionlogic.Decrypt(db)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8084", r))
}
