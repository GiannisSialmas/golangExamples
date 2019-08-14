Good examples at https://gobyexample.com

# Introduction 
## Protocol Buffers

Protocol Buffers is an alternative to sending data with JSON and HTTP. But the problem is that JSON is just a string and and the amount of information(and metadata) you can send with it is limited. It has a low information density per KB sent. Also the de/serialization and message validation is left as an exercise to the server. The server needs to make the check that the message it received has the correct format(all required fields exist, the values are of the correct type, etc)

Protocol Buffes utilize binary serialization, which means that they encode the data into a binary stream, which is lightweight and has higher information density.
The de/serialization process uses a predetermined schema to encode/decode the data. In the schema, you can define the data type of a field, if it is optional, etc. So you dont have to bake the validation inside your application logic
