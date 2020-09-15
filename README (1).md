# MP1
#Daniel Flores and Nicholas Hillis



----
# To Run
To run this program you first need to type the following commands into terminal


```bash
go run tcpC.go 1234
go run tcpS.go 1234
``` 
We know this is not how you were supposed to run the program, but we wanted to fix our TCP before we could fix any other issue.
However you can get multiple Clients to one server by opening multiple terminals, so we used goroutines to handle the connection

The program will open and begin to listen for inputs

Here is where you activate the tcp client:

```bash
Send message:
```

Where 2 is the process number and hello is the message. We used pointers to assign the user inputs, so it does not yet work with multiple words.
```bash
Example Code:
Send message: hello
Sent hello to process system time is "xxx"
```
The time function does work here, but we could not find a way to easily declare the process variable and left a placeholder value

```bash
Example Code:
Received hello from process system time is "xxx"
```
----
### Processes

We took the time to try and get a concurrent TCP server which took us much longer than we anticipated but we wanted to fix our problems from MP0 first so we did not have the time to implement config and a proper timeout. We hope to revise this code for resubmission so we can fix these issues.

### Sources

The following sites were used to help create our code:
```text
https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/#test-the-concurrent-tcp-server
https://opensource.com/article/18/5/building-concurrent-tcp-server-go
https://blog.golang.org/concurrency-timeouts 
https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/applevelprotocols/simple_example.html
http://zetcode.com/golang/readfile/ 
https://github.com/ashe7kh/TCP-Email-Program 
https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go 
https://medium.com/@onexlab.io/golang-config-file-best-practise-d27d6a97a65a 
https://github.com/tkanos/gonfig 
https://golang.org/pkg/encoding/gob/ 
https://github.com/matthieuberger/go-tcp-serverhttps://github.com/matthieuberger/go-tcp-server 
```