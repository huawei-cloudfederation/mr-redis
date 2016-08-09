package main

import (
        "net/http"
        "net/http/httptest"
        "fmt"
        "testing"
	"flag"
	"github.com/codegangsta/cli"
)

func TestIsRunning(T *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "A-OK")
        }))

        defer ts.Close()

	MrRedisFW = ts.URL

	IsRunning("TestInstance")

}

func TeststatusOf(T *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "A-OK")
        }))

        defer ts.Close()

        MrRedisFW = ts.URL

        statusOf("TestInstance")

}


func TestStatusCmd(T *testing.T){

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "valid json")
        }))

        defer ts.Close()

        MrRedisFW = ts.URL


	set := flag.NewFlagSet("test", 0)
        set.String("name", "TestInstance","doc")
        c := cli.NewContext(nil, set, nil)

	StatusCmd(c)
}

func TestStatusCmdEmptyName(T *testing.T){

	  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "A-ok")
        }))

        defer ts.Close()

        MrRedisFW = ts.URL


	set := flag.NewFlagSet("test", 0)
        set.String("name", "","doc")
        c := cli.NewContext(nil, set, nil)

	StatusCmd(c)
}

func TestStatusCmdInvalidJson(T *testing.T){

	  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "wrong json")
        }))

        defer ts.Close()

        MrRedisFW = ts.URL


	set := flag.NewFlagSet("test", 0)
        set.String("name", "TestInstance","doc")
        c := cli.NewContext(nil, set, nil)

	StatusCmd(c)
}
