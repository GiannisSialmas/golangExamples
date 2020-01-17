
# Pointers
## Get adress
Use &var to get the adress of the var
```go
a := "hello"
fmt.Println(a)  //a
fmt.Println(&a) // 0x1040a124
```
## Dereference adress
Use *var to get the value of a pointer
```go
a := "hello"
fmt.Println(a)  //hello
fmt.Println(&a) // 0x1040a124
fmt.Println(*&a) // hello

```


# Pretty print
```go
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
			fmt.Println(string(b))
	}
	return
}
```

# Print type of variable
```go
fmt.Printf("%T\n", a)
```

# Appending to slices
```go

var s []int
printSlice(s)
// append works on nil slices.
s = append(s, 0)
printSlice(s)
// The slice grows as needed.
s = append(s, 1)
printSlice(s)

// We can add more than one element at a time.
s = append(s, 2, 3, 4)
printSlice(s)

```
