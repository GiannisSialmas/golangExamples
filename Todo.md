# Channels and Routine
To make a function a routine, simply add `go` in front of it
```
go waitForGracefulShutdown 
```

```
myChannel := make(chan int) - creates myChannel which is a channel of type int

channel <- value - sends a value to a channel

value := <- channel - receives a value from a channel

<- channel - This is a deadlock. This will block the code until the channel receives a value

var done chan bool or
done := make(chan bool)
 // var initializes a channel as a nil channel. make initalizes the channel and sets up the backing data for the specified type.


By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.


```


Flag module