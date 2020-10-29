package user


import (
  "fmt"

)
var list  []userInfo
var userS userInfo
var userPointer *userInfo

var searchedUser []userInfo
func createUser(fname string, lname string, age int) *userInfo{

     userS := userInfo{
       fname:fname,
       lname :lname,
       age:age,
       address: userAdd{
         hno :12,
         addressLine1: "road 1",
         adressLine2 :"road 2",
         city:  "Delhi",
         country :"India",
       },
     }
    return &userS

}

func SaveUserData(fname string, lname string, age int) {
   addedToSlice(createUser(fname, lname, age))
}

func addedToSlice(pointerToUser *userInfo) *userInfo{
     // all validation code - start

     if pointerToUser.isFNameValid() && pointerToUser.checkAge() && pointerToUser.isLNameValid() {
      //   fmt.Println(userPointer.showAllUser())
         list = append(list, *pointerToUser)
         fmt.Println("\n=====================\n")
         for i, item := range list {
          fmt.Println(" (",i,") : Data: ",  item)
         }
         fmt.Println("\n=====================\n")
         return &userS
     } else {

        if !userPointer.isFNameValid() {
          fmt.Println("First Name out of char's limit")
        } else {
          fmt.Println("First name validation success")
        }

        if !userPointer.isLNameValid() {
          fmt.Println("Last Name out of char's limit")
        }  else {
          fmt.Println("Last name validation success")
        }

        if !userPointer.checkAge() {
            fmt.Println("Age out of range")
        } else {
           fmt.Println("Age name validation success")
        }
        fmt.Println("\nNo user created\n")
        return &userInfo{}
     }

 }


func Update(i int, fname string, lname string) {
  if len(list) <= 0 {
       fmt.Println("No user added!  Please create a user first\n")
  } else {

       if i < len(list) {
         list[i].fname = fname
         list[i].lname = lname
         list[i].age = 35
         list[i].address.hno = 123

         for i, item := range list {
           fmt.Println(i, item)
         }
       } else {
           fmt.Println("Index out of bound\n")
       }
    }
}


func DisplayUser(){
  fmt.Println("\nList of Users\n")
  fmt.Println("\n=====================\n")
  if len(list) <= 0 {
    fmt.Println("No data")
  } else {
    for i, item := range list {
      fmt.Println(" (",i,") : Data: ",  item)
    }
  }

  fmt.Println("\n=====================\n")
}


func SearchUser(fullName string){
   searchedUser = []userInfo{}
   if len(list) <=  0 {
     fmt.Println("\nCan't serach on empty list\n")
   } else {

     for  _ , item := range list {
       fullNameString := item.fname + " " + item.lname

       if fullNameString == fullName {
         searchedUser = append(searchedUser, item)

       }
     }
      fmt.Println("\nSerached result\n")
      fmt.Println(searchedUser)
   }


}



// validation methods

func (u *userInfo) isFNameValid() bool{
    return len((*u).fname) >= 3 && len((*u).fname) <= 10
}

func (u *userInfo) isLNameValid() bool{
   return len((*u).lname) >= 3 && len((*u).lname) <= 10
}

func (u *userInfo) checkAge() bool{
   return (*u).age > 0 && (*u).age <= 100
}
