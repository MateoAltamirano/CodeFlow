package main

func main() {
	mux := Route()
	server := CreateServer(mux)
	server.Run()
}