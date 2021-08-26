package main

import (
	"fmt"
	"math"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"strconv"

)

func main(){
	r := chi.NewRouter()

	r.Use(middleware.Logger)
  	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ingrese un número"))
  	})
  	r.Route("/{num}", func(r chi.Router) {
	  r.Get("/", func(w http.ResponseWriter, r *http.Request){
	  numero := chi.URLParam(r, "num")
	  fmt.Println(numero)
	  num, _ := strconv.Atoi(numero)
		//le quita los decimales al numero
		rdo :=extract(int(getResult(num)))
	  w.Write([]byte(fmt.Sprintf("Respuesta: %s\n", rdo)))

  	})
  })
  http.ListenAndServe(":3000", r)
}

//convierte a string y concatena
func extract(n int) string{

	strN := strconv.Itoa(n)
	c:=len(strN)
	fmt.Println(c)
	var rest string = ""
	if len(strN) == 2 {
		rest = "0"+strN
	}else if len(strN) == 1{
		rest = "00"+strN
	}else if len(strN) > 3{
		rest = strN[len(strN)-3:]
	}else{
		rest = strN
	}
	
	return rest
}

//saca la cuenta
func getResult(n int) float64 {
	return math.Pow((3 + math.Sqrt(5)),float64(n))
}