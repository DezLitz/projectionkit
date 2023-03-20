package dynamoprojection

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// HandlerOption is used to alter the behavior of AWS DynamoDB projection
// handler.
type HandlerOption interface {
	applyOptionToAdaptor(*decorators)
}

// ResourceRepositoryOption is used to alter the behavior of a ResourceRepository.
type ResourceRepositoryOption interface {
	applyResourceRepositoryOption(*decorators)
}

// TableOption is used to alter the behavior of oeprations that manipulate
// DynamoDB tables.
type TableOption interface {
	applyTableOption(*decorators)
}

type decorators struct {
	decorateGetItem            func(*dynamodb.GetItemInput) []func(*dynamodb.Options)
	decoratePutItem            func(*dynamodb.PutItemInput) []func(*dynamodb.Options)
	decorateDeleteItem         func(*dynamodb.DeleteItemInput) []func(*dynamodb.Options)
	decorateTransactWriteItems func(*dynamodb.TransactWriteItemsInput) []func(*dynamodb.Options)
	decorateCreateTableItem    func(*dynamodb.CreateTableInput) []func(*dynamodb.Options)
	decorateDeleteTableItem    func(*dynamodb.DeleteTableInput) []func(*dynamodb.Options)
}

type options struct {
	applyOptionToAdaptorFunc          func(*decorators)
	applyResourceRepositoryOptionFunc func(*decorators)
	applyTableOptionFunc              func(*decorators)
}

func (o *options) applyOptionToAdaptor(decorators *decorators) {
	if o.applyOptionToAdaptorFunc != nil {
		o.applyOptionToAdaptorFunc(decorators)
	}
}

func (o *options) applyResourceRepositoryOption(decorators *decorators) {
	if o.applyResourceRepositoryOptionFunc != nil {
		o.applyResourceRepositoryOptionFunc(decorators)
	}
}

func (o *options) applyTableOption(decorators *decorators) {
	if o.applyTableOptionFunc != nil {
		o.applyTableOptionFunc(decorators)
	}
}

// WithDecorateGetItem adds a decorator for DynamoDB GetItem operations.
//
// The decorator function may modify the input structure in-place. It returns a
// slice of DynamoDB request.Option values that are applied to the API request.
func WithDecorateGetItem(
	dec func(*dynamodb.GetItemInput) []func(*dynamodb.Options),
) interface {
	HandlerOption
	ResourceRepositoryOption
} {
	return &options{
		applyOptionToAdaptorFunc: func(d *decorators) {},
		applyResourceRepositoryOptionFunc: func(d *decorators) {
			d.decorateGetItem = dec
		},
	}
}

// WithDecoratePutItem adds a decorator for PutItem operation. The decorator can
// modify the passed PutItemInput structure and return a slice of request.Option
// to alter the request prior to its execution.
func WithDecoratePutItem(
	dec func(*dynamodb.PutItemInput) []func(*dynamodb.Options),
) interface {
	HandlerOption
	ResourceRepositoryOption
} {
	return &options{
		applyOptionToAdaptorFunc: func(d *decorators) {},
		applyResourceRepositoryOptionFunc: func(d *decorators) {
			d.decoratePutItem = dec
		},
	}
}

// WithDecorateDeleteItem adds a decorator for DeleteItem operation. The
// decorator can modify the passed DeleteItemInput structure and return a slice
// of request.Option to alter the request prior to its execution.
func WithDecorateDeleteItem(
	dec func(*dynamodb.DeleteItemInput) []func(*dynamodb.Options),
) interface {
	HandlerOption
	ResourceRepositoryOption
} {
	return &options{
		applyOptionToAdaptorFunc: func(d *decorators) {},
		applyResourceRepositoryOptionFunc: func(d *decorators) {
			d.decorateDeleteItem = dec
		},
	}
}

// WithDecorateTransactWriteItems adds a decorator for TransactWriteItems
// operation. The decorator can modify the passed TransactWriteItemsInput
// structure and return a slice of request.Option to alter the request prior to
// its execution.
//
// Warning! The order of the TransactWriteItems in the input structure is
// meaningful to both DynamoDB and this package. Specifically, the first item is
// used to update the projection's resource versions; it MUST NOT be modified or
// reordered.
func WithDecorateTransactWriteItems(
	dec func(*dynamodb.TransactWriteItemsInput) []func(*dynamodb.Options),
) interface {
	HandlerOption
	ResourceRepositoryOption
} {
	return &options{
		applyOptionToAdaptorFunc: func(d *decorators) {},
		applyResourceRepositoryOptionFunc: func(d *decorators) {
			d.decorateTransactWriteItems = dec
		},
	}
}

// WithDecorateCreateTable adds a decorator for CreateTable operation. The
// decorator can modify the passed CreateTableInput structure and return
// a slice of request.Option to alter the request prior to its execution.
func WithDecorateCreateTable(
	dec func(*dynamodb.CreateTableInput) []func(*dynamodb.Options),
) TableOption {
	return &options{
		applyTableOptionFunc: func(d *decorators) {
			d.decorateCreateTableItem = dec
		},
	}
}

// WithDecorateDeleteTable adds a decorator for DeleteTable operation. The
// decorator can modify the passed DeleteTableInput structure and return a slice
// of request.Option to alter the request prior to its execution.
func WithDecorateDeleteTable(
	dec func(*dynamodb.DeleteTableInput) []func(*dynamodb.Options),
) TableOption {
	return &options{
		applyTableOptionFunc: func(d *decorators) {
			d.decorateDeleteTableItem = dec
		},
	}
}
