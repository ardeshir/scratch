package main
import "fmt"
func main() {

   x := []int{6,7,8,9,5}

    // fmt.Println(x)
   
   /* x[0] = 0
   x[1] = 1
   x[2] = 2
   x[3] = 3
   x[4] = 4
   */
   
     // for i := 0; i < len(x) ; i++  {
     for i, v := range x {
        
      fmt.Printf("%d\t%d\t%v\n", i, v, x[i])
    
    }
    
}
