package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/GeertJohan/go.rice"
	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/liuzl/goutil/rest"
	"github.com/robertkrimen/otto"
)

var (
	addr    = flag.String("addr", ":8080", "bind address")
	grammar = flag.String("g", "math.grammar", "grammar file")
	jsfile  = flag.String("js", "math.js", "javascript file")
	start   = flag.String("start", "number", "start rule")
	eval    = flag.Bool("eval", false, "execute flag")
)

var (
	g      *fmr.Grammar
	onceG  sync.Once
	js     *otto.Otto
	onceJs sync.Once
	mux    sync.Mutex
)

func G() *fmr.Grammar {
	onceG.Do(func() {
		var err error
		g, err = fmr.GrammarFromFile(*grammar)
		if err != nil {
			panic(err)
		}
	})
	return g
}

func JS() *otto.Otto {
	onceJs.Do(func() {
		var err error

		script, err := ioutil.ReadFile(*jsfile)
		if err != nil {
			panic(err)
		}
		js = otto.New()
		if _, err = js.Run(script); err != nil {
			panic(err)
		}
	})
	return js
}

type Result struct {
	FMR string      `json:"fmr"`
	V   interface{} `json:"v"`
}

func errMsg(w http.ResponseWriter, err string) {
	rest.MustEncode(w, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{Status: "error", Message: err})
}

func MathHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	text := strings.TrimSpace(r.FormValue("text"))
	if text == "" {
		errMsg(w, "text is empty")
		return
	}

	start := strings.TrimSpace(r.FormValue("s"))
	if start == "" {
		start = "number"
	}
	p, err := G().EarleyParse(text, start)
	if err != nil {
		errMsg(w, err.Error())
		return
	}
	var results []*Result
	for _, f := range p.GetFinalStates() {
		trees := p.GetTrees(f)
		for _, tree := range trees {
			sem, err := tree.Semantic()
			if err != nil {
				errMsg(w, err.Error())
				return
			}

			result, err := JS().Run(sem)
			if err != nil {
				errMsg(w, err.Error())
				return
			}
			rs, _ := result.Export()
			results = append(results, &Result{sem, rs})
		}
	}
	rest.MustEncode(w, results)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	G()
	JS()
	defer glog.Info("server exit")
	http.Handle("/api/", rest.WithLog(MathHandler))
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
