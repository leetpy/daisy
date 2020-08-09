package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetEventsIndexName() string {
	year := GetCurrentYear()
	return fmt.Sprintf("%s-%d", viper.GetString("es.events_index_prefix"), year)
}
