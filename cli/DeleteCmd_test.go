package main

import (
        "net/http"
        "net/http/httptest"
	"fmt"	
//	"strings"
        "testing"
	 "github.com/codegangsta/cli"
	"flag"
)

//delete with valid input
func TestDelete(T *testing.T) {

        ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "A-OK")
        }))

        defer ts.Close()
        url := ts.URL

        _,err := httpDelete(url)

        if err != nil {
                //no error should occur
                T.Fail()
        }
}

//delete command  with valid input
func TestDeleteCmd(T *testing.T) {
	 ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintln(w, "A-OK")
        }))

        defer ts.Close()
	
	MrRedisFW = ts.URL
	set := flag.NewFlagSet("name", 0)
	set.String("name", "TestInstance","doc")
	c := cli.NewContext(nil, set, nil)

        DeleteCmd(c)

}

//delete command  with server error 
func TestDeleteCmdWithBadServer(T *testing.T) {

	 url := "https://10.11.12.13:5656"

        MrRedisFW = url 
	set := flag.NewFlagSet("name", 0)
	set.String("name", "TestInstance","doc")
	c := cli.NewContext(nil, set, nil)

        DeleteCmd(c)

}

//delete command  with empty input
func TestDeleteCmdWithEmptyName(T *testing.T) {
	  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
                fmt.Fprintln(w, "A-OK")
        }))
	
	 MrRedisFW = ts.URL

        set := flag.NewFlagSet("name", 0)
        set.String("name", "","doc")
        c := cli.NewContext(nil, set, nil)

        DeleteCmd(c)

}
