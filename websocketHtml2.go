package main

import "html/template"

func rootHandler(w http.ResponseWriter, r *http.Request) {
    rootTemplate.Execute(w, listenAddr)
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>
    websocket = new WebSocket("ws://{{.}}/socket");
    websocket.onmessage = onMessage;
    websocket.onclose = onClose;
</html>
`))
