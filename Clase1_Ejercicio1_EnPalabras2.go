//Adriana Bonilla Ugalde 206060468
//Sistemas Distribuidos  II Ciclo 2020
//Nombre de un número dado en palabras

package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
  //  "strings"
)

func handler(writer http.ResponseWriter, request *http.Request) {

  num, err := strconv.Atoi(request.URL.Path[1:])
  if err == nil {
      cadena := enPalabras(num)
      fmt.Fprintf(writer, "La cantidad %s en palabras es: %s ", request.URL.Path[1:],cadena)
   }
}

func enPalabras(num int) string {
      var unidades = []string{"cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
  	  var decenas =  []string{"cero","diez", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
      var cadena string

      if num< 10{
         cadena = unidades[num]
      }else{
          cadena = decenas[num/10]
          if num%10 > 0 {
             cadena = cadena + " y " + unidades[num%10]
          }
      }
      return cadena

}


func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
