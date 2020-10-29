package main

import (
  user "./user"
 "fmt"
 "bufio"
  "os"

)

func main(){
//  fmt.Println( user.CreateUser())


  gothere:  fmt.Println("\n\n\n\nOptions to select\n\n\n1 Create user\n\n2 Update user\n\n3 Find user\n\n4 Show users")
  fmt.Println("\nEnter options: ")

  var indexValue int
  //var fullName  string
    var fname string
    var lname string
    var age int
    var selectOption int
    fmt.Scanln(&selectOption)
    switch selectOption {
    case 1:
        fmt.Println("Create User\n\nEnter First name : ")
        fmt.Scanln(&fname)
        fmt.Println("Enter Last name : ")
        fmt.Scanln(&lname)
        fmt.Println("Enter age : ")
        fmt.Scanln(&age)
        user.SaveUserData(fname,lname, age)

        goto gothere
    case 2:

        fmt.Println("Enter index: ")
        fmt.Scanln(&indexValue)

        fmt.Println("Update F name: ")
        fmt.Scanln(&fname)


        fmt.Println("Update L name: ")
        fmt.Scanln(&lname)

        user.Update(indexValue, fname, lname)
        goto gothere
    case 3:
        fmt.Println("Search with User's full name: ")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan() // use `for scanner.Scan()` to keep reading
        fullName := scanner.Text()
        user.SearchUser(fullName)
        goto gothere
    case 4:
        user.DisplayUser()
        goto gothere
    default:
        fmt.Println("Do nothing")
        goto gothere
    }

}
