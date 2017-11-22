package main
import "fmt"
func main() {

   x := []int{6,7,8,9,5}

    fmt.Println(len(x))
   
   /* x[0] = 5
   x[1] = 6
   x[2] = 7
   x[3] = 8
   x[4] = 9
   */
   
     // for i := 0; i < len(x) ; i++  {
     for i, v := range x {
        
      fmt.Printf("%d\t%d\t = %v\n", i, v, x[i])
    
    }
    
    fmt.Println("Before the append() \n")
    fmt.Println(x)
    // slice the slice
    fmt.Println(x[1:len(x)])
    
    x = append(x, 87, 80, 90, 100, 182)
    
    
     fmt.Println("After append()\n")
     fmt.Println(x)
    // slice the slice
    fmt.Println(x[1:len(x)])
    
    y := []int{1212, 12312, 346, 66}
    x = append(x, y...)
    
    fmt.Println("After append(x, y...)\n")
     fmt.Println(x)
    // slice the slice
    fmt.Println(x[1:len(x)])
    
    
}
