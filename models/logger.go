package models

import "log"

type Logger struct {
    infoLogger    *log.Logger
    debugLogger   *log.Logger
    errorLogger   *log.Logger
    warningLogger *log.Logger
}
