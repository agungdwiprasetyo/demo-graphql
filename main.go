package main

func main() {
	service := NewService()
	service.ServeHTTP(8080)
}
