package repository

import "reflect"

const (
	TagParent = "model"

	TagValueIgnored = "ignore"
)

func ConvertDatabaseEntity(model, entity any) error {
	rEntityType := reflect.TypeOf(entity).Elem()

	rModel := reflect.ValueOf(model).Elem()
	rEntityValue := reflect.ValueOf(entity).Elem()

	length := rEntityType.NumField()
	for i := 0; i < length; i++ {
		field := rEntityType.Field(i)
		if field.Tag.Get(TagParent) == TagValueIgnored {
			continue
		}
		name := field.Name
		value := rEntityValue.FieldByName(name)
		rModel.FieldByName(name).Set(value)
	}

	return nil
}
