package configs

import "fmt"

func MessageLogConcatenate(prefix, sufix string) string {
	return fmt.Sprintf("[\"%s\"] %s", prefix, sufix)
}
