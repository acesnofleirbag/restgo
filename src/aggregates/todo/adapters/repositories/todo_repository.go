package repositories

import (
	"api/src/aggregates/todo/domain/entities"
	"api/src/core/adapters/uow"
	"fmt"
)

func Store(uow *uow.Postgres, todo entities.Todo) {
    query := fmt.Sprintf(`
        INSERT INTO 
            todos 
        VALUES (
            '%v', 
            '%v'
        )`, 
        1, 2,
    )

    uow.Run(query)
}
