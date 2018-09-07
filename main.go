package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
	"github.com/liuzl/fmr"
	"github.com/robertkrimen/otto"
)

var (
	grammar = flag.String("g", "math.grammar", "grammar file")
	js      = flag.String("js", "math.js", "javascript file")
	input   = flag.String("i", "", "file of original text to read")
	start   = flag.String("start", "result", "start rule")
	eval    = flag.Bool("eval", false, "execute flag")
)

func main() {
	flag.Parse()
	//fmr.Debug = true
	g, err := fmr.GrammarFromFile(*grammar)
	if err != nil {
		glog.Fatal(err)
	}
	script, err := ioutil.ReadFile(*js)
	if err != nil {
		glog.Fatal(err)
	}
	vm := otto.New()
	if _, err = vm.Run(script); err != nil {
		glog.Fatal(err)
	}

	var in *os.File
	if *input == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(*input)
		if err != nil {
			glog.Fatal(err)
		}
		defer in.Close()
	}
	br := bufio.NewReader(in)

	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			glog.Fatal(c)
		}
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}

		ps, err := g.EarleyParseMaxAll(line, *start)
		if err != nil {
			glog.Fatal(err)
		}
		for i, p := range ps {
			for _, f := range p.GetFinalStates() {
				trees := p.GetTrees(f)
				//fmt.Printf("%+v\n", p)
				fmt.Printf("p%d tree number:%d\n", i, len(trees))
				for _, tree := range trees {
					sem, err := tree.Semantic()
					if err != nil {
						glog.Fatal(err)
					}

					if !*eval {
						fmt.Println(sem)
					} else {
						result, err := vm.Run(sem)
						if err != nil {
							glog.Error(err)
						}
						rs, _ := result.Export()
						fmt.Printf("%s = %+v\n", sem, rs)
					}
					//eval, err := tree.Eval()
					//fmt.Printf("Eval: %s, Err: %+v\n", eval, err)
				}
			}
		}
		fmt.Println()
	}
}
