package asynq

import (
	"github.com/hibiken/asynq/internal/base"
	"github.com/hibiken/asynq/internal/errors"
)

type Broker = base.Broker

var ErrNoProcessableTask = errors.ErrNoProcessableTask
