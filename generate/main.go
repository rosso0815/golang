package main

//go:generate echo Hello, Go Generate!
//go:generate go run internal/generate/generatefoobar.go

func main() {
	Output()
}
