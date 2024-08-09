package entities

type Todo struct {
    id string
    title string
    description string
}

func NewTodo(id string, title string, description string) Todo {
    return Todo { 
        id: id,
        title: title,
        description: description,
    }
}

func (self *Todo) Get(key string) {}
