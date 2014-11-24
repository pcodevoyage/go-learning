type socket struct {
  io.ReadWriter
  done chan bool
}

func (s socket) Close() error {
    s.done <- true
    return nil
}

func socketHandler(ws *websocket.Conn) {
    s := socket{conn: ws, done: make(chan bool)}
    go match(s)
    <-s.done
}
