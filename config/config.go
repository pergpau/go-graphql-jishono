package config

import (
	"os"
	"fmt"
)

func DBConnectionString() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s", 
	"reader", 
	os.Getenv("DB_PASSWORD"), 
	"165.22.81.143", 
	"3306", 
	"jisho")
}