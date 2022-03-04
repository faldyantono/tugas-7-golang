package main
import ("fmt"
"runtime"
"reflect"
"time"
)
func main(){
  runtime.GOMAXPROCS(6)
  go reading(5, "Order")
  reading(5, "Refund")
}
func reading(Nomor int, Order string){
  for i:=0;i<=5;i++{
    fmt.Printf("%s\n", reflect.TypeOf(Nomor))
    fmt.Printf("%s\n", reflect.TypeOf(Order))
    time.Sleep(time.Millisecond * 100)
  }
}
