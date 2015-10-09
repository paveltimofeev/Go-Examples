package examplePackage

// private functions starts with 'lower' case character
func privateFoo() string {
    return "private function"
}

// public functions starts with UPPER case character
func PublicFoo() string {
  return "public function"
}

// unexported (private) scruct starts with 'lower' case character
type privateStruct struct {
  isPrivate bool
  IsPublic bool
}

// exported (public) scruct starts with UPPER case character
type PublicStruct struct {
  // private fields starts with 'lower' case character
  isPrivate bool
  // public fields starts with UPPER case character
  IsPublic bool
}

// exported (public) intrerfaces starts with UPPER case character
type PublicInterface interface {

  IsPublic() bool
}
