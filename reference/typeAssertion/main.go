var i interface{}
i = int(42)

a, ok := i.(int)
// a == 42 and ok == true

b, ok := i.(string)
// b == "" (default value) and ok == false

// Example with custom type

if uid, ok := sess.Values["user"].(bson.ObjectId); ok {
  ...
}

// sess.Values["user"] is an interface{}, and what is between parenthesis is called a type assertion. It checks that the value of sess.Values["user"] is of type bson.ObjectId. If it is, then ok will be true. Otherwise, it will be false.
