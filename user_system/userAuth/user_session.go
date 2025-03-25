package userAuth

import (
	"app/db"
	u "app/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type UserSession struct {
	Id    int
	Name  string
	Email string
	Role  string
	JWT   string
}

var ctx = context.Background()
var useRedis = db.Conn_Redis()

func SetUserSession(info UserSession) error {
	infoJSON, err := json.Marshal(info)
	if err != nil {
		fmt.Print(err)
		return err
	}

	err = useRedis.Set(ctx, u.ToString(info.Id), infoJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Não foi possível salvar no Redis \n", err)
		return err
	}

	return nil
}

func GetUserSession(Id string) (*UserSession, error) {
	var obj UserSession
	var query, err = useRedis.Get(ctx, Id).Result()
	if err != nil {
		fmt.Println("Erro de Consultar Redis: ", err)
		return &obj, err
	}

	err = json.Unmarshal([]byte(query), &obj)
	if err != nil {
		fmt.Println("Erro json.Unmarshal: ", err)
		return &obj, err
	}

	return &obj, nil
}

func LogoutUserSession(Id string) error {
	var err = useRedis.Del(ctx, Id).Err()
	if err != nil {
		fmt.Printf("Erro ao remover chave: %v", err)
		return err
	}

	return nil
}
