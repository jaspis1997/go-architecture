package main

import (
	"log"
	"playground"
	"playground/internal/app"
	_ "playground/internal/init"
	"playground/internal/repository"
	_ "playground/internal/repository/init"
)

func main() {
	e := playground.NewWebEngine()
	{
		config := repository.Config{
			Main: repository.SQLiteConfig{
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
	e.Run("", 8080, true)
}
