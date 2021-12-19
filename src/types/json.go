package types

type Response struct {
	Message string `json:"message"`
}

type Person struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}