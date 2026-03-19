package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("Monday, January 2, 2006 — 15:04")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Welcome</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            background: #fff;
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
            display: flex;
            align-items: center;
            justify-content: center;
            min-height: 100vh;
            text-align: center;
            color: #333;
        }
        .welcome { font-size: 2.5rem; font-weight: 300; }
        .date { font-size: 1rem; color: #999; margin-top: 0.5rem; }
    </style>
</head>
<body>
    <div>
        <div class="welcome">Welcome :)</div>
        <div class="date">%s</div>
    </div>
</body>
</html>`, now)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Println("listening on :3000")
	http.ListenAndServe(":3000", nil)
}
