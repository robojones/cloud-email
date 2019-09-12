//go:generate bash build.sh
//go:generate wire

package main

func main() {
	s := InitServer()
	s.ListenAndServe()
}
