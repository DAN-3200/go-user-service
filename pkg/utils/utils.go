package utils

import (
	"fmt"
	"reflect"
)

func Ternary[Anyone any](conditional bool, value1 Anyone, value2 Anyone) Anyone {
	if conditional {
		return value1
	}
	return value2
}

func MapSQLInsertFields(sqlFields map[string]string, payload any) ([]string, []any, error) {
	v := reflect.ValueOf(payload)
	t := reflect.TypeOf(payload)

	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("payload precisa ser struct")
	}

	var (
		cols []string
		args []any
		i    = 0
	)

	for j := 0; j < v.NumField(); j++ {
		field := v.Field(j)
		fieldType := t.Field(j)
		sqlCol, ok := sqlFields[fieldType.Name]

		if !ok {
			continue // ignora campos sem mapeamento
		}

		if field.IsNil() {
			continue // pula campos nil (não informados)
		}

		deref := field.Elem() // pega valor apontado

		i++
		cols = append(cols, fmt.Sprintf("%s=$%d", sqlCol, i))
		args = append(args, deref.Interface())
	}

	if len(cols) == 0 {
		return nil, nil, fmt.Errorf("não há campos para mapear")
	}

	return cols, args, nil
}
