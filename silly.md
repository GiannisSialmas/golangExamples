
# Pointers
## Get adress
Use &var to get the adress of the var
```go
a := "hello"
fmt.Println(a)  //a
fmt.Println(&a) // 0x1040a124
```
## Dereference adres
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