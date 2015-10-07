package main
import ("fmt"; "time"; "math/rand"; "strconv")

func main(){

  fmt.Println("CHANNELS EXAMPLE")

  // создание канала
  var c chan string = make(chan string)

  go Sender(c)
  go Reciever(c)

  time.Sleep(time.Second * 3)

  fmt.Println("CHANNELS EXAMPLE")

  SelectChannels()

  // time.Sleep(time.Second * 3)
}

func Sender(c chan string){

  for {

    c <- "message"               // пишет в канал
    time.Sleep(time.Second * 1 ) // ждёт 1 сек
  }
}

func Reciever(c chan string){

  for {

    msg := <- c // считывает с канала сообщение как только оно там появиться
    fmt.Println("Recieved: ", msg)
    // сразу, без паузы переходит к след. циклу
  }
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

func SelectChannels() {

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
