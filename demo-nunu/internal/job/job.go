package job

import (
	"demo-nunu/internal/repository"
	"demo-nunu/pkg/jwt"
	"demo-nunu/pkg/log"
	"demo-nunu/pkg/sid"
)

type Job struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
