package main

type Student struct {
	Id      uint64 `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Age     uint64 `json:"age" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
