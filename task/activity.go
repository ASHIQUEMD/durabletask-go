package task

import (
	"context"

	"github.com/microsoft/durabletask-go/internal/protos"
)

// ActivityContext is the context parameter type for activity implementations.
type ActivityContext interface {
	GetInput(resultPtr any) error
	Context() context.Context
}

type activityContext struct {
	TaskID int32
	Name   string

	rawInput []byte
	ctx      context.Context
}

// Activity is the functional interface for activity implementations.
type Activity func(ctx ActivityContext) (any, error)

func newTaskActivityContext(ctx context.Context, taskID int32, ts *protos.TaskScheduledEvent) *activityContext {
	return &activityContext{
		TaskID:   taskID,
		Name:     ts.Name,
		rawInput: []byte(ts.Input.GetValue()),
		ctx:      ctx,
	}
}

// GetInput unmarshals the serialized activity input and saves the result into [v].
func (actx *activityContext) GetInput(v any) error {
	return unmarshalData(actx.rawInput, v)
}

func (actx *activityContext) Context() context.Context {
	return actx.ctx
}
