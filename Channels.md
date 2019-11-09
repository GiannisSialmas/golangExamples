## Prevent go program from exiting
If ou want to execute go routines and don't have anything in main, put the below to block it
```go
fmt.Scanln()
```

## Go Routines
To make a function a routine, simply add `go` in front of it. It will detach and start to run on it's own.
```go
go waitForGracefulShutdown 
```
## Go channels
var initializes a channel as a nil channel. make initalizes the channel and sets up the backing data for the specified type.
```go
var done chan bool or
done := make(chan bool)
```

### Buffered or unbuffered?
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.

For unbuffered channel, the sender will block on the channel until the receiver receives the data from the channel, whilst the receiver will also block on the channel until sender sends data into the channel.

Compared with unbuffered counterpart, the sender of buffered channel will block when there is no empty slot of the channel(you define the slots when you create the channel), while the receiver will block on the channel when it is empty.


## Useful commands
```go
myChannel := make(chan int) - creates myChannel which is a unbuffered channel of type int
channel <- value - send a value to a channel
value := <- channel - receive a value from a channel

<- channel - This is a deadlock. This will block the code until the channel receives a value

```