package main
import "fmt"

func main() {

  structsExample()
  interfacesExample()
}


// ---------------------------------------------------------------------------

func structsExample()  {

  slice := make([]int, 5)

  // инициализация структуры с передачей значения полям
  printer1 := SlicePrinter { Header: "SlicePrinter1", Item: " printing item value:" }
  printer1.print(slice)

  // инициализация структуры
  var printer2 SlicePrinter
  printer2.Header = "SlicePrinter2"
  printer2.Item = " print item val:"
  printer2.print(slice)

  // инициализация структуры с возвратом указателя
  var printer3 = new(SlicePrinter)
  printer3.Header = "SlicePrinter3"
  printer3.Item = " slice item value:"
  printer3.print(slice)

  // инициализация структуры с вложенным типом
  var wrapper = new(SlicePrinterWrapper)
  wrapper.Printer.Header = "SlicePrinterWrapper"
  wrapper.Printer.print( slice )

  // инициализация структуры с вложенным типом, объявленным как базовый
  var wrapper2 = new(SlicePrinterWrapper2)
  wrapper2.Header = "SlicePrinterWrapper2"
  wrapper2.print( slice )
}

// объявление структуры
type SlicePrinter struct {
  Header string
  Item string
}

// добавление метода printSlice к структуре Printer
func (p *SlicePrinter) print(arr []int) {

  fmt.Println(p.Header)
  for _, val := range arr {
    fmt.Println(p.Item, val)
  }
}


// вложенный тип
// wrapper := SlicePrinterWrapper
// wrapper.Printer.print( ... )
type SlicePrinterWrapper struct {
  Printer SlicePrinter
}

// вложенный тип, как базовый
// wrapper := SlicePrinterWrapper2
// wrapper.print( ... )
type SlicePrinterWrapper2 struct {
  SlicePrinter
}

// ---------------------------------------------------------------------------

func interfacesExample(){

    array := [5]int{1,2,3,4,5}

    var printer1 = new(SlicePrinter)
    printer1.Header = "SlicePrinter implements interface IPrinter"

    var printer2 = new(SlicePrinterWrapper2)
    printer2.Header ="SlicePrinterWrapper2 implements interface IPrinter"

    PrintWithIPrinter(printer1, array[1:3])
    PrintWithIPrinter(printer2, array[1:3])
}

// объявление интерфейса
type IPrinter interface {
  print(arr []int)
}

// использование интерфейcа
func PrintWithIPrinter(printer IPrinter, arr []int) {
  printer.print(arr)
}

// ---------------------------------------------------------------------------
