package main

import (
   "flag"
   "log"
)

func main(){
   var name string
   flag.StringVar(&name,"name","go tour", "help")
   flag.StringVar(&name,"n","go tour", "help")
   flag.Parse()
   log.Printf("name :%s",name)
}
