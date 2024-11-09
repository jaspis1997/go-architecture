package main

import (
	"log"
	"playground/internal/app"
	_ "playground/internal/init"
	"playground/internal/repository"
	"playground/internal/repository/database"
	"playground/internal/web"
)

func main() {
	e := web.New()
	{
		config := repository.Config{
			Main: database.SQLiteConfig{
				Filename: "test.db",
			},
		}
		repo, err := repository.New(config)
		if err != nil {
			log.Fatal(err)
		}
		err = repository.Migrate(repo)
		if err != nil {
			log.Fatal(err)
		}
		err = app.Init(repo)
		if err != nil {
			log.Fatal(err)
		}
	}
	e = routes(e)
	e.Run("localhost", 8080)
}
