//go:generate go run github.com/google/wire/cmd/wire

package main

func main() {
	s := InitServer()
	s.ListenAndServe()
}
