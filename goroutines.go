package main
import ("fmt"; "time")

func main()  {

  fmt.Println("GOROUTINES EXAMPLE")

  // запуск функции как гоурутины
  go foo1();
  // запуск функции как гоурутины
  go foo2();
  // запуск функции как гоурутины
  go foo3();

  time.Sleep(time.Second * 10)
}

func foo1()  {
  for{
    time.Sleep(time.Millisecond * 1900)
    fmt.Print("як-")
  }
}

func foo2()  {
  for{
    time.Sleep(time.Millisecond * 2000)
    fmt.Print("цуп-")
  }
}

func foo3()  {
  for{
    time.Sleep(time.Millisecond * 3000)
    fmt.Println("цоп")
  }
}
