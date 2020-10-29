package user

type userInfo struct {
  fname string
  lname string
  age int
  address userAdd
}
type userAdd struct {
  hno int
  addressLine1 string
  adressLine2 string
  city string
  country string
}


type errorEl struct{
  errorMsg string
}
