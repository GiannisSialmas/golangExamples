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

## Fetch dependencies
```
go get
```
This will fetch all dependencies for a project
When it encounters an import of a package not provided by any module in go.mod, the go command automatically looks up the module containing that package and adds it to go.mod, using the latest version.
Only direct dependencies are recorded in the go.mod file

## List all dependencies (Direct and Indirect)
```
go list -m all
```
This will output a list with go modules installed in your project.
**A better solution would be to look at your go.sum file**



## Build the project
```
go build
```















## Protocol Buffers

Protocol Buffers is an alternative to sending data with JSON and HTTP. But the problem is that JSON is just a string and and the amount of information(and metadata) you can send with it is limited. It has a low information density per KB sent. Also the de/serialization and message validation is left as an exercise to the server. The server needs to make the check that the message it received has the correct format(all required fields exist, the values are of the correct type, etc)

Protocol Buffes utilize binary serialization, which means that they encode the data into a binary stream, which is lightweight and has higher information density.
The de/serialization process uses a predetermined schema to encode/decode the data. In the schema, you can define the data type of a field, if it is optional, etc. So you dont have to bake the validation inside your application logic.
Also due to the fact that the messages are of binary format, it is less intensive for the cpu to de/serialize the message, which can be a good thing for mobile devices for example. 


