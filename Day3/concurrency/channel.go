package main

import (
  "fmt"
  "time"
)

func firstTask( ch chan string ) {
  var msg string
  for i := 0; ; i++ {
	  //Writing some message onto the channel - Send
	  msg = fmt.Sprintf("From Task 1 %d", i )
	  ch <- msg 
	  fmt.Println ("Task 1: sent a new message")
  }

}

func secondTask( ch chan string ) {
  for {
	//Retrieving the message from the channel - Receive
	message := <- ch
	fmt.Println ( "Task 2: received message ", message )
	time.Sleep( time.Millisecond * 3 )
  }
}

func main() {
  //create a channel for first and second task to communicate with each other
  //This channel will accept string type messages
  var ch chan string = make(chan string)

  go firstTask(ch)
  go secondTask(ch)

  var waitForKeyboardPress string
  fmt.Println("Press any key to quit ...")
  fmt.Scanln(&waitForKeyboardPress)
}
