# golang-chat
A creation from phase 3 of DevBootcamp: new technology "discovery project": golang + angular.js. 

Walk through of the app: 
[Part1] (https://youtu.be/UEucu9PjzC0)
[Part2] (https://youtu.be/dV4SC6KJQC8)

## GO / Angular project

### Setup: 

1. Install GO

2. Create a workspace = directory hierarchy that contains go source code and the package objects and command binaries that the compilers produces from that source code. 

Where? 
Any directory. 

Creating a WORKSPACE in a directory called “gocode”
```
$ mkdir gocode
```
To tell the GO TOOL about that workspace, set the go path environment variable. 
```
$ export GOPATH=$HOME/gocode
```
Source code resides in the src directory of your workspace

making a directory called src. inside creating a NAMESPACE - a unique base path inside which all of your gocode will reside  
```
$ cd gocode 
```
Using the base path of my github account: 
```
$ mkdir -p src/github.com/erinbeitel
```
note: -p creates any missing intermediate pathname components.

### Creating Hello World program: 

creating the directory hello inside the namespace. 
```
$ cd src/github.com/erinbeitel
$ mkdir hello
```
 making a file hello.go inside the hello directory
```
$ vim hello.go
```
inside hello.go we will put the standard hello.go program. 

note: this code belongs to package main. this is a convention that tells the go tool to produce an executable command instead of a package object that would be imported by other code. 

We can use the go tool to build and install that binary. 
:wq to write and quit the vim editor. 

$ go install

the go tool has installed gocode.go to the bin directory inside the workspace. 
```
$ ls ~/gocode/bin
```
```
$ ~/gocode/bin/hello  -> prints Hello, new gopher! 
```

Just wrote, built, and executed a go program!

### Creating a simple function that reverses a string:
Adding the bin directory to our system path bc we will be installing and running a bunch of go programs in the future. 
```
$ export PATH=$HOME/gocode/bin:$PATH
```
now we can just run hello to run the hello world program. 

Creating a package. Packages are just like commands, except that they can be imported by other packages and have a package name other than main. 

Starting with a string handling package. 

create a directory inside the namespace called “string”. 
```
$ mkdir string
```
```
$ cd string
```
inside of string, create a string.go file. 
```
$ vim string.go
```
#### Creating the string reverse function

Note: the function starts with a capital “R”. This means it can be exported and used by other packages. 

Exiting the editor, we can check that the function compiles by running
```
$ go build
```
Where did the object file go?
When used on go packages, go build just builds the code and then throws away the output. To make the string package available to other code, we need to install the package object to our workspace with: 
```
$ go install 
```

Now, we can find the package object inside the pkg directory of our workspace. 

Now we can use this string package from our hello program. 
```
$ cd ../hello
$ vim hello.go
```
update hello.go

compile and run hello command
```
$ go install 
$ hello
```
the go tool has located the github.com/erinbeitel/string package and used that when compiling and linking the hello program.

running hello prints the string “Hello new gopher!" backwards.

### TESTS: 
The tests for a package live in that package’s directory, in files that end in _test 

creating a test string_test.go inside test directory. 
```
$ vim string_test.go 
```

Using the testing package, we will write a quick table driven test for the reverse function. 

note: the string_test.go file is also part of package string but it won’t be compiled into the package when you run go install or go build. However, when we run
```
$go test
```
the go tool will compile all of the package files, including the tests, and synthesize a test harness that runs any functions beginning with Test (with a capital T). 

### Creating a Chat App

Check that the GOPATH is setup correctly. 
$ echo $GOPATH
(if it returns a path, that means you have it configured correctly)

Check that we have GO on our path.
$ go
(should return something other than command not found)

Package declaration statement: 
package main
This package happens to be “main”. Main is a special package for GO in that it can create executables. 

import statement

function declaration 
```
func main() {}
```
to run the file using local host: 
```
$ go run main.go
```
(note: port = 8000)

#### Making a chat app: 

1. started a new OS repo with a GO gitignore and a MIT license and a read me. 

2. creating a file called main.go

3. creating a new directory called web
```
$ mkdir web (use this later)
```
4. creating a file in web called about.html
```
$ vim about.html
```
5. creating a main.go file
```
$ vim main.go
```
(added this to github repo)

6. adding an external dependency: github.com/gorilla/websocket (putting that in main.go)
'''
$ go get github.com/gorilla/websocket
'''
This downloads the source and puts it in the github.com folder, on the go path. 
(note: in the demo, the gorilla folder he downloaded had a .git folder in it. might need to look into that. I think he had downloaded it previously.)  

7. Adding a handler that will accept a web socket connection. into main.go, adding: 
'''
http.HandleFunc("/ws", wsHandler)
'''
8. Also wring a wsHandler function inside of main.go.
There is a writer (w) so anything we write will be returned to the client. 
```
 func wsHandler(w ResponseWriter, r *Request){
}
```
The second part is for the client’s request.

Once our web socket hits the “/ws" we setup above, it will hit this function and then we can do stuff with it. 
First: we will upgrade the connection. It comes in as a http and that gets upgraded to a web socket connection. 

Calls upgrade on the variables <strong>conn</strong> and <strong>err</strong>. if it’s successful you get the connection object back (conn).
```
  11 func wsHandler(w http.ResponseWriter, r *http.Request) {
  12   //from gorilla
  13   conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
  14   if _, ok := err.(websocket.HandshakeError); ok {
  15     http.Error(w, "Not a websocket handshake", 400)
  16     return
  17   } else if err != nil {
  18     log.Println(err)
  19     return
  20   }
  21   defer conn.Close()
  22   for {
  23     _, msg, err := conn.ReadMessage()
  24     if err != nil {
  25       return
  26     }
  27     log.Println(string(msg))
  28     if err = conn.WriteMessage(websocket.TextMessage, msg); err != nil {
  29       return
  30     }
  31   }
  32 }
```
We want to close that connection object so we do defer conn.Close at the end. 

We don’t want to leave the function unless we are closing it or the client is intentionally closing it so we put the for infinite loop at the bottom. 

now we are going to work on the client: 

9.  Making a second html file inside web called chat.go.
naming the app 
```
<html ng-app="golang-chat">
```
creating our first controller: <strong>MainCt1</strong>

Go server was created.

### Making a chat server

Need to store the connections somewhere. 

Making a function called <strong>sendAll</strong> that will take a msg that we want to send to every connection and loop over all of the connections and write whatever message that we were told to send all. 

If there is an error, we are going to remove from the connections the connection that had the error. 

We are also going to add the connection if we are getting a new one (in the conn.Close, err section of main.go
delete(connections, conn)

also need to add the connection if we are getting a new one… also in conn.Close

<strong> connections[conn] = true </strong>

Finally, add send all at the end of that same section.

Opening localhost:8000 in the browswer and localhost:8000 in a private browing window will allow chat between the two windows.
