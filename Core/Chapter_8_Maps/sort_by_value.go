
// Go program to sort the map by Values
package main
  
import (
    "fmt"
    "sort"
)
  
func main() {
    basket := map[string]int{"orange": 5, "apple": 7, 
                             "mango": 3, "strawberry": 9}
  
    keys := make([]string, 0, len(basket))
  
    for key := range basket {
        keys = append(keys, key)
    }
  
    fmt.Println(basket)
    fmt.Println(keys)
  
    sort.SliceStable(keys, func(i, j int) bool{
        return basket[keys[i]] < basket[keys[j]]
    })
  
    fmt.Println(keys)
}