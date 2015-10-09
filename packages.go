package main
import "fmt"
import "github.com/ptimofeev/Go-examples/examplePackage"

func main(){

  // call of exported (public) fonction
  result := examplePackage.PublicFoo()
  fmt.Println("public function =", result)

  // creation of exported (public) struct
  var pStruct examplePackage.PublicStruct
  fmt.Println("false =", pStruct.IsPublic)

  // implementation of exported (public) interface
  test := new(Test)
  fmt.Println("true =", Check(test) )
}

// usage of exported (public) interface
func Check(obj examplePackage.PublicInterface) bool {

  return obj.IsPublic()
}


type Test struct { }

func (t *Test) IsPublic() bool {
  return true
}
