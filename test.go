package main
import "fmt"
import "os"
//import "time"

// объявление типа
var text3 uint = 1
// объявление константы
const pi int = 3
// объявление неск. переменных
var (
  t1 = "text"
  t2 = "message"
  t3 = "commentary"
)
const (
  G = 9.8
  G1 = "10"
)

// ---------------------------------------------------------------------------

func main() {

  var text string = "zzz"
  var count int = len(text)
  fmt.Println(text[0], count * pi)

  text2 := "xxx"
  fmt.Println(text2)

  arrayExample()
  sumExample(1,2,3)
  sumExample(2,4,8,16,32)
  panicExample()
  pointersExample()
}

// ---------------------------------------------------------------------------

func arrayExample(){

  // create array
  var a [5]float64

  for i:=0; i< len(a) ;i++ {
    a[i] = float64(i)
  }

  for _, val := range a {
    fmt.Println("item:", val)
  }

  // create array
  a2 := [5]int{ 1, 2, 3, 5, 8 }
  printArray(a2)

  // создание среза
  cut1 := make([]int, 5)
  printSrez(cut1)

  // создание среза
  cut2 := make([]int, 5, 10)
  printSrez(cut2)

  // создание среза из массива
  cut3 := a2[1:3]
  printSrez(cut3)

  cut4 := a2[:3]
  printSrez(cut4)

  cut5 := a2[3:]
  printSrez(cut5)

  cut6 := a2[:]
  //cut6 := a2[:len(a2)] // the same
  printSrez(cut6)

  // создание словаря
  map1 := make(map[string]float32)
  map1["USD"] = 50.0
  map1["EUR"] = 60.0
  map1["EURO"] = 60.0

  // создание словаря с ининциализацией значений
  map2 := map[string]float32{
    "EUR/USD": 1.33,
    "USD/EUR": 0.85,
  }

  // удаление элемента
  delete(map1, "EURO")
  // получение элемента с проверкой на наличие
  val, success := map1["EURO"]
  if !success {
    fmt.Println("Element EURO was deleted successfully")
  }else{
    fmt.Println("Element EURO =", val)
  }

  if v, ok := map1["EUR"]; ok {
    fmt.Println("Element EUR =", v)
  }

  printMap(map1)
  printMap(map2)


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
}

func printArray(a [5]int){
  fmt.Println("-array:")
  for _, val := range a {
    fmt.Println("item:", val)
  }
  fmt.Println("")
}

func printSrez(a []int){
  fmt.Println("-srez:")
  for _, val := range a {
    fmt.Println("item:", val)
  }
  fmt.Println("")
}

func printMap(m map[string]float32){
  for key, val := range m {
    fmt.Println(key, val)
  }
}

// ---------------------------------------------------------------------------


// функция, принимающяя несколько аргументов
func sumExample(args ...int) int{

  var result int = 0
  for _,a := range args {
    result += a
    fmt.Println(a, "=>", result)
  }
  fmt.Println("Total =", result)
  return result
}

// функция возвращяющая несколько значений
func readInput(msg string) (inpt string, success bool){
  fmt.Println(msg)
  fmt.Scanf("%s", &inpt)
  success = inpt != ""
  return inpt, success
}

// отложенные функции
func deferExample(){

  file,_ := os.Open("test.go")

  defer file.Close()
  // файл закроется при выходе из deferExample(),
  // даже если произойдёт исключение
  // даже если в функции несколько return'ов

  // ниже могут идти операции с file
}


// ---------------------------------------------------------------------------


func pointersExample()  {
  i := 0;
  incr(&i) // передача значения по ссылке (передача указателя)
  fmt.Println("Incremented i=", i )
  incr(&i) // передача значения по ссылке (передача указателя)
  fmt.Println("Incremented i=", i )
}

// функция принимающая указатель на переменную (передача по ссылке)
// используется так: incr( &i)
func incr( val *int)  {
  *val++ // обращение к указателю всегда делается через *.
}


// ---------------------------------------------------------------------------


// panic & recover, исключение и его перехват
func panicExample(){

  // отложенный вызов анонимной функции
  defer func()  {
    str := recover() // восстановление, возвращает сообщение panic
    fmt.Println("Panic:", str)
  }()

  // выброс исключения, прерывание программы, но отложенная функция
  // всё равно должна отработать
  panic("Exception!")
}


// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
