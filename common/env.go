package common 


import (
	"os"
)


func Get_env (key string) string {

	return os.Getenv(key)
}
