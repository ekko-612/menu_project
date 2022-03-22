package main
  
import ( "fmt")
 
func main() {
        var inout string
        fmt.Print("go menu\n")
        for {
                fmt.Print("this is a menu\n>")
                fmt.Scanln(&inout)
                if inout == "quit" {
                        break
                }
                switch inout {
                case "hello":
                        fmt.Println("hello page get in")
                case "edit":
                        fmt.Println("set get in")
                case "quit":
                        break
                default:
                        fmt.Println("wrong order")
                }
        }
        
}
