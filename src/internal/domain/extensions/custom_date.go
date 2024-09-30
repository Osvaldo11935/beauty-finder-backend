// /pkg/extensios/custom_date_time.go
package extensios

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

const (
	customDateLayout       = "2006-01-02"
	CustomDateLayout       = "2006-01-02 15:04:05"
	fullDateTimeLayout     = "2006-01-02 15:04:05 -0700 MST"
	extendedDateTimeLayout = "2006-01-02 15:04:05 +0000 UTC"
)

func (cdt *CustomDate) UnmarshalJSON(b []byte) error {
	dateString := string(b)
	if len(dateString) > 2 {
		dateString = dateString[1 : len(dateString)-1] // Remove aspas
	}

	// Tentar analisar o formato de data e hora completo
	parsedTime, err := time.Parse(CustomDateLayout, dateString)
	if err == nil {
		cdt.Time = parsedTime
		return nil
	}

	// Tentar analisar o formato de data apenas
	parsedTime, err = time.Parse(customDateLayout, dateString)
	if err == nil {
		cdt.Time = parsedTime
		return nil
	}

	// Tentar analisar o formato de data e hora completo com fuso horário
	parsedTime, err = time.Parse(fullDateTimeLayout, dateString)
	if err == nil {
		cdt.Time = parsedTime
		return nil
	}

	// Tentar analisar o formato de data e hora completo com fuso horário extendido
	parsedTime, err = time.Parse(extendedDateTimeLayout, dateString)
	if err == nil {
		cdt.Time = parsedTime
		return nil
	}

	return fmt.Errorf("não foi possível analisar a data: %v", dateString)
}

func (cdt *CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cdt.Time.Format(CustomDateLayout))
}

func (cdt *CustomDate) Value() (driver.Value, error) {
	return cdt.Time.Format(CustomDateLayout), nil
}

func (cdt *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*cdt = CustomDate{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*cdt = CustomDate{Time: v}
	case string:
		var t time.Time
		var err error
		formats := []string{
			CustomDateLayout,
			customDateLayout,
			fullDateTimeLayout,
			extendedDateTimeLayout,
		}
		for _, format := range formats {
			t, err = time.Parse(format, v)
			if err == nil {
				*cdt = CustomDate{Time: t}
				return nil
			}
		}
		return err
	case []byte:
		var t time.Time
		var err error
		formats := []string{
			CustomDateLayout,
			customDateLayout,
			fullDateTimeLayout,
			extendedDateTimeLayout,
		}
		for _, format := range formats {
			t, err = time.Parse(format, string(v))
			if err == nil {
				*cdt = CustomDate{Time: t}
				return nil
			}
		}
		return err
	default:
		return fmt.Errorf("cannot scan type %T into CustomDate: %v", value, value)
	}
	return nil
}
