package main

import "time"

//Todo 個別のTodoを定義したstruct
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos Todoの配列
type Todos []Todo
