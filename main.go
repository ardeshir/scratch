package main
import "fmt"
func main() {

   x := []int{0,1,2,3,4}

    fmt.Println(x)
   
   /* x[0] = 0
   x[1] = 1
   x[2] = 2
   x[3] = 3
   x[4] = 4
   */
   
     //for i := 0; i < len(x) ; i++  {
    for k, v := range x {
        
      fmt.Printf("%d\t%v\n", x[k], x[v] )
    
    }
    
}
