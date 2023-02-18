
# GoKer

| Bug ID|  Ref | Patch | Type | SubType | SubsubType |
| ----  | ---- | ----  | ---- | ---- | ---- |
|[grpc#1275]|[pull request]|[patch]| Blocking | Communication Deadlock | Channel |

[grpc#1275]:(grpc1275_test.go)
[patch]:https://github.com/grpc/grpc-go/pull/1275/files
[pull request]:https://github.com/grpc/grpc-go/pull/1275
 
## Description

Some description from developers or pervious reseachers

> Two goroutines are invovled in this deadlock. The first goroutine
  is the main goroutine. It is blocked at case <- donec, and it is
   waiting for the second goroutine to close the channel.
   The second goroutine is created by the main goroutine. It is blocked
   when calling stream.Read(). stream.Read() invokes recvBufferRead.Read().
   The second goroutine is blocked at case i := r.recv.get(), and it is
   waiting for someone to send a message to this channel.
   It is the client.CloseSream() method called by the main goroutine that
   should send the message, but it is not. The patch is to send out this message.

Possible intervening

```
///
/// G1 									G2
/// testInflightStreamClosing()
/// 									stream.Read()
/// 									io.ReadFull()
/// 									<- r.recv.get()
/// CloseStream()
/// <- donec
/// ------------G1 timeout, G2 leak---------------------
///
```

