// pprof 서버 활성
// 이 패키지를 포함 하려면
// import _ "github.com/hyhecor/go-debug/pprof"
//// +build debug

package pprof

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	_ "net/http/pprof"
)

func init() {
	const defaultPort int = 16060

	hosturl := fmt.Sprintf("localhost:%d", defaultPort)
	url := fmt.Sprintf("http://localhost:%d/debug/pprof/", defaultPort)

	// bin, err := exec.LookPath("start")
	// if err != nil {
	// 	println("Not found binary. start")
	// }
	bin := "start"
	args := []string{url}
	env := os.Environ()
	syscall.Exec(bin, args, env)

	go func() {
		println("pprof:", http.ListenAndServe(hosturl, nil))
	}()
}
