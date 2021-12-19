package types

type Response struct {
	Message string `json:"message"`
}

type Person struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}

type DbConnection struct {
	DbName string `json:"db_name"`
	Username string `json:"name"`
	Password string `json:"password"`
}