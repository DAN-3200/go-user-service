package adapters

// Necessita do redis-server ativo
import (
	"app/internal/domain/entity"
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

type SessionCache struct{}

func (it SessionCache) SetUserSession(info entity.UserSession) error {
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

func (it SessionCache) GetUserSession(Id string) (*entity.UserSession, error) {
	var obj entity.UserSession
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

func (it SessionCache) LogoutUserSession(Id string) error {
	err := useRedis.Del(ctx, Id).Err()
	if err != nil {
		fmt.Printf("Erro ao remover chave: %v", err)
		return err
	}

	return nil
}

func (it SessionCache) GetInfoSession(ctx *gin.Context, key string) (*entity.UserSession, error) {
	var value_format *entity.UserSession

	value_base, ok := ctx.Get(key)
	if !ok {
		return value_format, fmt.Errorf("Key não existe")
	}

	value_format, ok = value_base.(*entity.UserSession)
	if !ok {
		return value_format, fmt.Errorf("Erro ao formatar a informação")
	}
	return value_format, nil
}
