package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

const maxRequestBodySize = 20 * 1024 * 1024 // 20 MB

// handlePostRequest handles POST requests on the "/post" path
func handlePostRequest(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Header.ContentType()) != "application/octet-stream" {
		ctx.Error("Invalid Content-Type", fasthttp.StatusUnsupportedMediaType)
		return
	}

	body := ctx.PostBody()
	hash := sha256.Sum256(body)
	fmt.Printf("SHA256 Hash: %x %d bytes\n", hash, len(body))

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(fmt.Sprintf("%x", hash))
}

// handleIndexRequest handles requests to the root path and returns index.html
func handleIndexRequest(ctx *fasthttp.RequestCtx) {
	file, err := os.ReadFile("index.html")
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(file)
}

func main() {
	server := fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {

			switch string(ctx.Path()) {
			case "/post":
				handlePostRequest(ctx)
			case "/":
				handleIndexRequest(ctx)
			default:
				ctx.Error("Not Found", fasthttp.StatusNotFound)
			}
		},
	}

	server.MaxRequestBodySize = maxRequestBodySize
	// Launch the server on port 8080
	log.Fatal(server.ListenAndServe(":8080"))
}
