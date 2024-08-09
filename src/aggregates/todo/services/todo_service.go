package services

import (
	TodoRepo "api/src/aggregates/todo/adapters/repositories"
	Todo "api/src/aggregates/todo/domain/entities"
)

func Store() {
    todo := Todo.NewTodo("1", "Lorem ipsum", "Description...")

    TodoRepo.Store(todo)
}
