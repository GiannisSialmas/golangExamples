# General
```go
type person struct {
    first string
    last  string
    age   int
}
 
p1 := person{
   first: "James",
   last:  "Bond",
   age:   32,
}
fmt.Println(p1.first, p1.last, p1.age)
> James Bond 32
```

# Print Key along with value
```go
fmt.Printf("%+v\n", struct)
```

# Embedded
```go
type person struct {
    first string
    last  string
    age   int
}
type secretAgent struct {
    person
    ltk bool
}
    sa1 := secretAgent{
        person: person{
            first: "James",
            last:  "Bond",
            age:   32,
        },
        ltk: true,
    }
    fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
    fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)
> James Bond 32 true
> James Bond 32 true

```
**All properties of the embedded type become first level citizens of the top level object**

# Anonymous Structs
```go
    p1 := struct {
        first string
        last  string
        age   int
    }{
        first: "James",
        last:  "Bond",
        age:   32,
    }

```