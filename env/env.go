package env

import (
	"log"
	"os"
	"sync"

	env "github.com/joho/godotenv"

)

var once sync.Once

//LoadEnv carga el archivo .env si aun no se ha cargado
func LoadEnv()  {
	once.Do(func(){
		if err:= env.Load();err != nil {
			log.Println("No se encontró el archivo .env, procediendo con variables de entorno alternas")
		}
	})
}

//GetEnv devuelve el valor de una variable de entorno
func GetEnv(key, fallback string) string {
	LoadEnv()//Asegurar que el entorno ya esté cargado

	value := os.Getenv(key)
	if value == ""{
		return fallback
	}
	
	return value
}