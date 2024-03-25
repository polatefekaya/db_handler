package internals

import (
	"reflect"
	"time"
)

// playerE
func FillTime[a any](T a) a {
	c, valid := getValidValue(T, "CreatedAt")
	if valid {
		c.Set(reflect.ValueOf(time.Now().UTC()))
	}
	u, valid := getValidValue(T, "UpdatedAt")
	if valid {
		u.Set(reflect.ValueOf(time.Now().UTC()))
	}
	return T
}

func getValidValue[a any](T a, name string) (reflect.Value, bool) {
	v := reflect.ValueOf(T).Elem()
	f := v.FieldByName(name)

	notvalid := !f.IsValid() || !f.CanSet()

	if notvalid {
		return f, false
	}

	return f, true
}

func FillStatId[a any](T a, id string) a {
	i, valid := getValidValue(T, "StatisticId")
	if valid {
		i.SetString(id)
	}
	return T
}
