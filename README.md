bytecounter
---

A byte-counting io.Writer.

Will take an action after a certain amount of bytes passed through the underlying writer.

### Usage

```go
func main() {
  // ...

  // lets assume you have a writer w
  // and you want to track how many bytes pass through it
  // and take an action when the number of bytes reaches a certain threshold
  bc := NewByteCounter(w, 1024, func() {
    fmt.Println("One KB has passed through w")
  })

  // lets assume you have a reader r
  io.Copy(bc, r)

  // You can also replace w with bc
  w = bc

  // ...
}
```
