# General
The best practice dictated that an error should be returned last from a function

# Create new error
Use error.New to create a new basic error message
```go
func Sqrt(value float64)(float64, error) {
   if(value < 0){
      return 0, errors.New("Math: negative number passed to Sqrt")
   }
   return math.Sqrt(value)
}
```