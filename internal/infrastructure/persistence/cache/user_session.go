package cache

// Necessita do redis-server ativo
import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var useRedis *redis.Client

// Primeiro a ser executado
func InitCoreRedis(core *redis.Client) {
	useRedis = core
}

var Session = LayerUserSession{}

type LayerUserSession struct{}

type UserSession struct {
	Id    string
	Name  string
	Email string
	Role  string
	JWT   string
}

func (it LayerUserSession) SetUserSession(Id string, Name string, Email string, Role string, JWT string) error {
	info := UserSession{
		Id:    Id,
		Name:  Name,
		Email: Email,
		Role:  Role,
		JWT:   JWT,
	}
	infoJSON, err := json.Marshal(info)
	if err != nil {
		fmt.Print(err)
		return err
	}
	err = useRedis.Set(ctx, info.Id, infoJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Não foi possível salvar no Redis \n", err)
		return err
	}

	return nil
}

func (it LayerUserSession) GetUserSession(Id string) (*UserSession, error) {
	var obj UserSession
	query, err := useRedis.Get(ctx, Id).Result()
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

func (it LayerUserSession) LogoutUserSession(Id string) error {
	err := useRedis.Del(ctx, Id).Err()
	if err != nil {
		fmt.Printf("Erro ao remover chave: %v", err)
		return err
	}

	return nil
}

func (it LayerUserSession) GetInfoSession(ctx *gin.Context, key string) (*UserSession, error) {
	var value_format *UserSession

	value_base, ok := ctx.Get(key)
	if !ok {
		return value_format, fmt.Errorf("Key não existe")
	}

	value_format, ok = value_base.(*UserSession)
	if !ok {
		return value_format, fmt.Errorf("Erro ao formatar a informação")
	}
	return value_format, nil
}
