// pprof 서버 활성
// 이 패키지를 포함 하려면 
// import _ "github.com/hyhecor/go-debug/pprof"
// +build debug

package pprof

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

import _ "net/http/pprof"

const (
	defaultPort int = 16060
)

func init() {
	url := fmt.Sprintf("%s:%d", "localhost", defaultPort)

	text := fmt.Sprintf(`[Desktop Entry]
Encoding=UTF-8
Name=pprof
Type=Link
URL=http://%s/debug/pprof/
Icon=text-html
`, url)

	ioutil.WriteFile("pprof.desktop", []byte(text), 0755)

	go func () {
		println("pprof:", http.ListenAndServe(url, nil))
	}()
}
