package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"script-import/db"
	"script-import/domain"
	"script-import/script-import/repository"
	"script-import/script-import/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pandeptwidyaop/golog"
	"github.com/sirupsen/logrus"
	"github.com/thatisuday/commando"
)

func Init() {
	e := godotenv.Load("../.env")
	if e != nil {
		log.Println(".env not found, using global variable")
	}
	golog.New()
	db.NewGormClient()

}

func main() {
	commando.
		SetExecutableName("samapta-cli").
		SetVersion("0.0.1").
		SetDescription("This CLI tool for smartbtw exam services")
	commando.
		Register("seed:samapta").
		SetDescription("seeding for testing purpose only").
		SetShortDescription("seeding for testing purpose only").
		SetAction(func(m1 map[string]commando.ArgValue, m2 map[string]commando.FlagValue) {
			Init()
			content, err := os.ReadFile("data.json")
			if err != nil {
				logrus.Fatal("Error when opening file: ", err)
			}

			if err != nil {
				logrus.Fatal(errors.New("error on marshaling json " + err.Error()))
			}

			cgRepo := repository.NewPostgrePelanggan(db.GormClient.DB)
			userRepo := repository.NewPostgreUser(db.GormClient.DB)
			timeoutContext := fiber.Config{}.ReadTimeout
			ucScript := usecase.NewPelangganUseCase(cgRepo, userRepo, timeoutContext)

			var bd []domain.InsertData
			err = json.Unmarshal(content, &bd)
			if err != nil {
				logrus.Fatal("Error when unmarshaling json: ", err)
			}

			// fmt.Println("ini bd: ", bd)
			id := 20270
			pelangganID := 20209
			for index, val := range bd {
				val.UserID = uint(id)
				val.PelangganID = uint(pelangganID)
				id += 1
				pelangganID += 1
				err := ucScript.CreateImportPelanggan(&val)
				if err != nil {
					logrus.Fatal("Error when inserting data: ", err)
				}
				logrus.Info("Nomor: ", index, " ", "Name: ", val.NamaPelanggan, " ", "SUCCESSFUL")
			}
			logrus.Info("Successfully running script")
		})

	commando.Parse(nil)

}
