package repositories

import (
	"context"
	"fmt"
	"itsware/internal/models"
	"itsware/pkg/db"
	"reflect"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
)

func CreateCabinet(ctx context.Context, cabinet models.Cabinet) (*models.Cabinet, error) {
	cabinetType := pgtype.Record{
		Fields: []pgtype.Value{
			&pgtype.Text{String: cabinet.Name, Status: pgtype.Present},
			&pgtype.Text{String: cabinet.Location, Status: pgtype.Present},
		},
	}

	query := "SELECT * FROM create_cabinet($1);"

	var createdCabinet models.Cabinet

	rows, err := db.Pool.Query(ctx, query, cabinetType)
	if err != nil {
		return nil, fmt.Errorf("error in query: %w", err)
	}

	createdCabinet, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.Cabinet])
	if err != nil {
		return nil, fmt.Errorf("error executing function: %w", err)
	}

	return &createdCabinet, nil
}

func generateRecord(input any) (*pgtype.Record, error) {
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct")
	}

	fields := make([]pgtype.Value, 0)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type()
		fieldName := t.Field(i).Name

		status := pgtype.Present
		if field.Kind() == reflect.Ptr && field.IsNil() {
			status = pgtype.Null
		}

		switch fieldType.Kind() {
		case reflect.String:
			fields = append(fields, &pgtype.Text{
				String: field.String(),
				Status: status,
			})
		case reflect.Ptr:
			if field.Elem().Kind() == reflect.String {
				fields = append(fields, &pgtype.Text{
					String: field.Elem().String(),
					Status: status,
				})
			}
		default:
			panic(fmt.Sprintf("Unsupported field type for field '%s'", fieldName))
		}
	}

	return &pgtype.Record{Fields: fields}, nil
}
