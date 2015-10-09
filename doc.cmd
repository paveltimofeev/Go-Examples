@echo off
cls

echo -----------------------------------------------------------

REM - Get documentation for examplePackage
godoc github.com/ptimofeev/Go-Examples/examplePackage

echo -----------------------------------------------------------

REM - Get documentation for PublicFoo() in examplePackage
godoc github.com/ptimofeev/Go-Examples/examplePackage PublicFoo

echo -----------------------------------------------------------

REM - Run documentation server at 8080 port
echo Open 'http://localhost:8080/pkg/github.com/ptimofeev/Go-Examples/'
godoc -http=":8080"

echo -----------------------------------------------------------
