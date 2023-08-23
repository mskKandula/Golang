package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/mskKandula/websockOpt/websock"
)

func main() {

	// Commented out since running in docker,By default docker container has max file descriptor limit
	// Increase resources limitations
	// var rLimit syscall.Rlimit
	// if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
	// 	panic(err)
	// }

	// rLimit.Cur = rLimit.Max
	// if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Starting Server...")

	r := gin.Default()

	pprof.Register(r)

	pool := websock.NewPool()

	go pool.Start()

	r.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c.Writer, c.Request)
	})
	r.Run(":8000")
}

func serveWs(pool *websock.Pool, w http.ResponseWriter, r *http.Request) {

	conn, err := websock.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websock.Client{
		Conn: conn,
		Pool: pool,
	}

	// go client.Read()

	pool.Register <- client
}
