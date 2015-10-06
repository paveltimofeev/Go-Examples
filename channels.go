package main
import ("fmt"; "time")

func main(){

  fmt.Println("CHANNELS EXAMPLE")

  // создание канала
  var c chan string = make(chan string)

  go Sender(c)
  go Reciever(c)

  time.Sleep(time.Second * 3)
}

func Sender(c chan string){

  c <- "message"
}

func Reciever(c chan string){
  msg := <- c
  fmt.Println("Recieved: ", msg)
}
