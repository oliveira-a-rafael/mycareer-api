package main

import (
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/oliveira-a-rafael/my-career-api/app"
	"github.com/oliveira-a-rafael/my-career-api/config"
	"github.com/oliveira-a-rafael/my-career-api/controllers"
	"github.com/oliveira-a-rafael/my-career-api/database"
	"github.com/oliveira-a-rafael/my-career-api/domains"
)

func main() {

	log.Info("initializing app at: %s", time.Now())

	port := config.APP_PORT
	if port == "" {
		err := errors.New("app port not defined")
		log.Fatal(err)
		panic(err)
	}

	if config.APP_RESET {
		if err := initBase(); err != nil {
			log.Error(err.Error())
		}
	} else {
		err := createBase()
		if err != nil {
			panic("Error on createBase :: " + err.Error())
		}
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	router.HandleFunc("/careersTest", controllers.ListaCareersToTest).Methods("GET")

	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST", "OPTIONS")

	router.HandleFunc("/career/new", controllers.CreateCareer).Methods("POST", "OPTIONS")
	router.HandleFunc("/career/{id}", controllers.UpdateCareer).Methods("PUT", "OPTIONS")
	router.HandleFunc("/career/{id}", controllers.DeleteCareer).Methods("DELETE")
	router.HandleFunc("/careers", controllers.ListCareers).Methods("GET")
	router.HandleFunc("/career/{id}", controllers.GetCareer).Methods("GET")
	router.HandleFunc("/career/{id}/players", controllers.CareerPlayers).Methods("GET")

	router.HandleFunc("/player", controllers.CreatePlayer).Methods("POST", "OPTIONS")
	router.HandleFunc("/player", controllers.ListPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", controllers.UpdatePlayer).Methods("PUT", "OPTIONS")
	router.HandleFunc("/player/{id}", controllers.GetPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", controllers.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/player/points", controllers.CreatePlayerPoints).Methods("POST", "OPTIONS")

	router.HandleFunc("/positions", controllers.GetPosition).Methods("GET")

	router.Use(app.JwtAuthentication)

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Accept-Endcoding", "Content-Length"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins(config.ALLOWED_HOSTS))(router)))

}

func createBase() (err error) {
	db := database.GetInstance()
	exists := db.HasTable(&domains.Account{})
	defer db.Close()
	if !exists {
		err = migrateDataBase()
		if err != nil {
			log.Error(err)
			return
		}
	}
	return err
}

func initBase() (err error) {

	err = purgeDataBase()
	if err != nil {
		log.Error(err)
		return
	}

	err = createBase()
	if err != nil {
		log.Error(err)
		return
	}

	return err
}

func migrateDataBase() (err error) {
	db := database.GetInstance()

	log.Info("Trying to mirate db")
	err = db.AutoMigrate(
		&domains.Account{},
		&domains.Career{},
		&domains.Player{},
		&domains.PlayerSkill{},
		&domains.PlayerHistory{},
		&domains.PlayerPoints{},
	).Error

	if err != nil {
		log.Error(err)
		defer db.Close()
		return err
	}

	defer db.Close()
	return err
}

func purgeDataBase() (err error) {
	log.Info("Trying to purge db")
	if config.APP_ENV == "dev" {
		db := database.GetInstance()
		err := db.DropTableIfExists(
			&domains.Account{},
			&domains.Career{},
			&domains.Player{},
			&domains.PlayerSkill{},
			&domains.PlayerHistory{},
			&domains.PlayerPoints{},
		).Error
		if err != nil {
			db.Close()
			return err
		}
		db.Close()
	} else {
		log.Info("Action delete table is not permited in this enviroment")
	}

	return err
}
