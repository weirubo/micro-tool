package template

var (
	DomainSRV = `package domain

import (
	"context"
)

// {{title .Alias}} is representing the Article data struct
type {{title .Alias}} struct {
	Id        int64 
	Name      string 
	Created   int
	Updated   int
}

// {{title .Alias}}Usecase represent the article's usecases
type {{title .Alias}}Usecase interface {
	Get(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	GetList(ctx context.Context, {{.Alias}} *{{title .Alias}}) ([]*{{title .Alias}}, error)
	Create(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	Update(ctx context.Context, {{.Alias}} *{{title .Alias}}, condition *{{title .Alias}}) (int64, error)
}

// {{title .Alias}}Repository represent the article's repository contract
type {{title .Alias}}Repository interface {
	Get(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	GetList(ctx context.Context, {{.Alias}} *{{title .Alias}}) ([]*{{title .Alias}}, error)
	Create(ctx context.Context, {{.Alias}} *{{title .Alias}}) error
	Update(ctx context.Context, {{.Alias}} *{{title .Alias}}, condition *{{title .Alias}}) (int64, error)
}`
)
