package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/norlight/zimuzu-go/client"
	"github.com/norlight/zimuzu-go/resource"
	"github.com/norlight/zimuzu-go/search"
)

//填入你自己的 api key
const (
	cid       = ""
	accesskey = ""
)

func init() {
	log.SetPrefix("ZIMUZU: ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(os.Stdout)
}

func main() {
	if len(cid) == 0 || len(accesskey) == 0 {
		log.Println("configure your api key first.")
		os.Exit(1)
	}

	c := client.New(cid, accesskey)
	r := resource.New(&c)
	s := search.New(&c)

	var k string
	var chID string

	flag.StringVar(&k, "s", "", "search")
	flag.StringVar(&chID, "r", "", "fetch resource links")
	flag.Parse()
	if n := flag.NFlag(); n != 0 {
		if n > 1 {
			flag.Usage()
			os.Exit(1)
		}
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "s":
				result, err := s.AlfSearch(k)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(result))
				os.Exit(0)

			case "r":
				result, err := r.AlfList(chID)
				if err != nil {
					panic(err)
				}
				fmt.Print(string(result))
				os.Exit(0)
			}
		})
	}

	if len(os.Args) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "top":
		result, err := r.AlfTop()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(result))
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
