//go:generate wire

package main

func main() {
	s := InitServer()
	s.ListenAndServe()
}
