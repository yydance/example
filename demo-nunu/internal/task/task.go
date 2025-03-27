package task

import (
	"demo-nunu/internal/repository"
	"demo-nunu/pkg/jwt"
	"demo-nunu/pkg/log"
	"demo-nunu/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
