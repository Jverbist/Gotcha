package main

import (
  "fmt"
  "log"
  "net/http"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  
  
)


func helloHandler(w http.ResponseWriter, r *http.Request)  {
  if r.URL.Path != "/hello" {
    http.Error(w, "404 Dumbass", http.StatusNotFound)
    return
  }
  
  if r.Method != "GET" {
    http.Error(w, "Use GET to get something usefull", http.StatusNotFound)
    return
  }


  fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request)  {
  if err := r.ParseForm(); err != nil {
    fmt.Fprintf(w, "ParseForm() err: %v", err)
    return 
  }
  fmt.Fprintf(w, "POST request successful\n")
  firstname := r.FormValue("fname")
  lastname := r.FormValue("lname")
  email := r.FormValue("email")

  fmt.Fprintf(w, firstname)
  fmt.Fprintf(w, "Last Name = %s\n", lastname)
  fmt.Fprintf(w, "email = %s\n", email)



}


func addUsers(db *sql.DB, first_name string, last_name string, email string) {
  records := `INSERT INTO people(first_name, lastname, email) VALUES (?, ?, ?)`
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec(first_name, last_name, email)
  if err != nil {
    log.Fatal(err)
  }
}











func main()  {

//SQLlite3 information and variables
  // db, err := sql.Open("sqlite3", "./names.db")
  // checkErr(err)






//webserver part
  fileserver := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileserver) // handler dat de variable fileserver gebruikt om de paginas uit de folder /static aan te bieden
  http.HandleFunc("/hello", helloHandler) // hier komt de functie die reageert op /hello. inclusief de error handling wanner het niet over GET of de juiste URL gaat
  http.HandleFunc("/form", formHandler)






  fmt.Printf("Starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
} 
