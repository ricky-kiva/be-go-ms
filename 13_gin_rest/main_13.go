package main

type Student struct {
	Student_id      uint64 `json:"student_id" binding:"required"`
	Student_name    string `json:"student_name" binding:"required"`
	Student_age     uint64 `json:"student_age" binding:"required"`
	Student_address string `json:"student_address" binding:"required"`
	Student_phone   string `json:"student_phone" binding:"required"`
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
