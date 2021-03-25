package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	ERROR_LEVEL_FATAL = "fatal"
	ERROR_LEVEL_WARN  = "warn"
	ERROR_LEVEL_ERROR = "error"
	ERROR_LEVEL_PANIC = "panic"
)

const (
	LOG_LEVEL_TRACE = "trace"
	LOG_LEVEL_DEBUG = "debug"
	LOG_LEVEL_INFO  = "info"
)

func MessageNew(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(err.Error())
	}
}

func RespondNew(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(err.Error())
	}
}

func RespondError(w http.ResponseWriter, err error) {
	ErrorHandler(err, ERROR_LEVEL_ERROR)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(MessageNew(err.Error())); err != nil {
		log.Error(err.Error())
	}
}

func ErrorHandler(err error, errorType string) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	switch errorType {
	case "fatal":
		log.Fatal(err.Error())
	case "warn":
		log.Warn(err.Error())
	case "error":
		log.Error(err.Error())
	case "panic":
		log.Panic(err.Error())
	default:
		log.Info(err.Error())
	}
}

func BuildHashedPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ConvertToUint(val string) (uint, error) {
	u64, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		ErrorHandler(err, ERROR_LEVEL_ERROR)
		return 0, err
	}
	return uint(u64), nil
}

var currentUser uint

var GetCurrentUser = func() (uint, error) {
	if currentUser < 0 {
		return currentUser, errors.New("can not access user logged")
	}
	return currentUser, nil
}

func SetCurrentUser(id uint) {
	currentUser = id
}
