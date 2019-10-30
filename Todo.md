# Channels and Routine
To make a function a routine, simply add `go` in front of it
```
go waitForGracefulShutdown 
```

```
myChannel := make(chan int) - creates myChannel which is a channel of type int

channel <- value - sends a value to a channel

value := <- channel - receives a value from a channel
```


Flag module