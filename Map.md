## Declare a map of strings to ints
```go 
    m := map[string]int{
        "James":           32,
        "Miss Moneypenny": 27,
    }
```

## LOOK IF YOU CAN DECLARE A MAP map[string]interface{}

## Get value of key from map
```go
    fmt.Println(m["James"])
> 32
```
## See if property exists
If if the key you ask from the map does not exist, instead of an error you will get an empty value. To see if the key truly exists in the map, do the following
```go
    v, ok := m["Barnabas"]
    fmt.Println(v)
    fmt.Println(ok)
> 0
> false
```
## Only run if property exists. Very mainstream
```go
    if v, ok := m["Barnabas"]; ok {
        fmt.Println(v)
    } else {
        fmt.Println("Not found")
    }
> Not found
```
## Delete from map
```go
delete(m, "James")
```