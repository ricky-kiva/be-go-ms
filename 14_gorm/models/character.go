package models

type Character struct {
	Id     uint64 `json:"character_id"`
	Name   string `json:"character_name" binding:"required"`
	Level  uint64 `json:"character_level" binding:"required"`
	Gender string `json:"character_race" binding:"required"`
}
