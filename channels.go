package main
import ("fmt"; "time"; "math/rand"; "strconv")

func main(){

  SimpleChannelsExample()
  SelectChannelsExample()
  BufferedChannelsExample()
}

func SimpleChannelsExample(){

  fmt.Println("CHANNELS EXAMPLE")

  // создание канала
  var c chan string = make(chan string)

  go Sender(c) // запуск гоурутины отправляющей сообщение по каналу каждую секунду
  go Reciever(c) // запуск гоурутны считывающей сообщения с канала

  time.Sleep(time.Second * 3)
}

func Sender(c chan string){

  for {

    c <- "message"               // пишет в канал
    time.Sleep(time.Second * 1 ) // ждёт 1 сек
  }
}

func Reciever(c chan string){

  for {

    // считывает с канала сообщение как только оно появиться
    // пока сообщения нет блокирует поток
    msg := <- c
    fmt.Println("Recieved: ", msg)

    // сразу, без паузы переходит к след. циклу
  }
}

//-------------------------------------------------------------------

func SelectChannelsExample() {

  fmt.Println("CHANNELS EXAMPLE")

  var c1 chan string = make(chan string)
  var c2 chan string = make(chan string)

  go SendingOnlyChan(c1, "message from channel 1")
  go SendingOnlyChan(c2, "message from channel 2")

  go RecievingOnlyChan(c1)
  go RecievingOnlyChan(c1)

  go func(){

    for {

      select {
       case message := <- c1:
          fmt.Println("Recieving from Chan1: ", message)
        case message := <- c2:
          fmt.Println("Recieving from Chan2: ", message)
      }
    }

  }()


  time.Sleep(time.Second * 3)
}

// канал с доступом только на запись (только для отправки)
func SendingOnlyChan(c chan<- string, message string){

  for {

    rand := rand.New(rand.NewSource(99))
    var duration = strconv.Itoa(rand.Intn(1000))
    d,_ := time.ParseDuration( duration + "ms" )

    c <- "SendingOnly " + message + ", delay:  " + duration

    time.Sleep(time.Second * d )
  }
}

// канал только на чтение (только для получения)
func RecievingOnlyChan(c <-chan string){

  for {
    recieved := <- c
    fmt.Println("RecievingOnlyChan: ", recieved)
  }
}

//-------------------------------------------------------------------

func BufferedChannelsExample() {

  fmt.Println("BUFFERED CHANNELS EXAMPLE")

  var bc chan string = make(chan string, 3)

  go MessageSender(bc, "one", 100)
  go MessageSender(bc, "two", 100)
  go MessageSender(bc, "three", 100)
  go MessageSender(bc, "FOUR", 100)

  go MessageReciever(bc, 2000)

  time.Sleep(time.Second * 5)
}

func MessageSender(c chan string, msg string, d time.Duration){

  for {

    fmt.Println(" Sending: ", msg)
    c <- msg
    fmt.Println("    Sent: ", msg)
    time.Sleep(time.Millisecond * d )
  }
}

func MessageReciever(c chan string, d time.Duration){

  for {
    fmt.Println("     Got: ", <- c)
    time.Sleep(time.Millisecond * d )
  }
}
