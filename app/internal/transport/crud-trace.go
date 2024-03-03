// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"crud/internal/interfaces"
	"crud/internal/models"
	"github.com/opentracing/opentracing-go"
)

type traceCrud struct {
	next interfaces.Crud
}

func traceMiddlewareCrud(next interfaces.Crud) interfaces.Crud {
	return &traceCrud{next: next}
}

func (svc traceCrud) Create(ctx context.Context, req models.UserRequest) (resp models.UserResponse, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "Create")
	return svc.next.Create(ctx, req)
}

func (svc traceCrud) GetUser(ctx context.Context) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "GetUser")
	return svc.next.GetUser(ctx)
}

func (svc traceCrud) GetUsers(ctx context.Context) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "GetUsers")
	return svc.next.GetUsers(ctx)
}

func (svc traceCrud) Update(ctx context.Context) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "Update")
	return svc.next.Update(ctx)
}

func (svc traceCrud) Delete(ctx context.Context) (err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("method", "Delete")
	return svc.next.Delete(ctx)
}
