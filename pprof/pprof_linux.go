// pprof 서버 활성
// 이 패키지를 포함 하려면
// import _ "github.com/hyhecor/go-debug/pprof"
//// +build debug

package pprof

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"syscall"

	_ "net/http/pprof"
)

func init() {
	const defaultPort int = 16060

	hosturl := fmt.Sprintf("%s:%d", "localhost", defaultPort)
	url := fmt.Sprintf("http://%s:%d/debug/pprof/", "localhost", defaultPort)

	bin, err := exec.LookPath("xdg-open")
	if err != nil {
		println("Not found binary. xdg-open")
	}
	args := []string{url}
	env := os.Environ()
	syscall.Exec(bin, args, env)

	go func() {
		println("pprof:", http.ListenAndServe(hosturl, nil))
	}()
}
