# Tips

## Get values from environment
Use os.LookupEnv instead of os.Getenv

```go
func LookupEnv(key string) (string, bool)
```
```go
func Getenv(key string) string
```

os.Getenv("XXX") retrieves the value of the environment variable named by the key. It returns the value, which will be empty if the variable is not present. LookupEnv retrieves the value of the environment variable named by the key. If the variable is present in the environment the value (which may be empty) is returned and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.



Good examples at https://gobyexample.com

# Introduction 
## Start a new project
```
go mod init <project-name>
```
Now you’ll have a go.mod file. If you’re familiar with npm, think of this as your package.json when you are using npm. This is where all of the imports that your module needs, are specified. It’ll be initialized with your new module being declared at the top of it.

## Add dependencies
If you have created your go.mod file, 
```
go get <package-name>
```
will fetch that package and add it to go.mod file with an exact version.
If the package exists in your go.mod file, **it will be updated to the latest version**
Only direct dependencies are recorded in the go.mod file

## Fetch all dependencies
```
go mod download
```
This will fetch all dependencies that are listed in go.mod and go.sum

## List all dependencies (Direct and Indirect)
```
go list -m all
```
This will output a list with go modules installed in your project.
**A better solution would be to look at your go.sum file**

## Remove unused dependencies 
```
go mod tidy
```
This cleans up the go.mod and go.sum with packages that are not used

## Local packages import
```
repo/
├── go.mod      <<<<< Note go.mod is located in repo root
├── pkg1
│   └── pkg1.go
└── pkg2
    └── pkg1.go
```
Then pkg1 would import its peer package as import "repo/pkg2". Note that you cannot use relative import paths like import "../pkg2" or import "./subpkg"

## Test
```go
go test ./ ...
```
Look or tests in current and **all** subdirectories


## Protocol Buffers

Protocol Buffers is an alternative to sending data with JSON and HTTP. But the problem is that JSON is just a string and and the amount of information(and metadata) you can send with it is limited. It has a low information density per KB sent. Also the de/serialization and message validation is left as an exercise to the server. The server needs to make the check that the message it received has the correct format(all required fields exist, the values are of the correct type, etc)

Protocol Buffes utilize binary serialization, which means that they encode the data into a binary stream, which is lightweight and has higher information density.
The de/serialization process uses a predetermined schema to encode/decode the data. In the schema, you can define the data type of a field, if it is optional, etc. So you dont have to bake the validation inside your application logic.
Also due to the fact that the messages are of binary format, it is less intensive for the cpu to de/serialize the message, which can be a good thing for mobile devices for example. 

