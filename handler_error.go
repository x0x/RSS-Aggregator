package main

import "net/http"


func handlerError(w http.ResponseWriter, r* http.Request) {
  
  respondWithError(w, 500, "Interval Server Error")

}

