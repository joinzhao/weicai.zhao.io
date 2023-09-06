package repox

import (
	"gorm.io/gorm/schema"
	"weicai.zhao.io/gormx"
)

type Repo interface {
	RepoName() string
}

type Tabler[T schema.Tabler] interface {
	New() T
}

type ReaderWriterRepo[T schema.Tabler] interface {
	Repo
	Tabler[T]
	Creator[T]
	Updater[T]
	Deleter[T]
	First[T]
	Find[T]
	Counter[T]
}

type WriterRepo[T schema.Tabler] interface {
	Repo
	Tabler[T]
	Creator[T]
	Updater[T]
	Deleter[T]
}

type ReaderRepo[T schema.Tabler] interface {
	Repo
	Tabler[T]
	First[T]
	Find[T]
	Counter[T]
}

type Creator[T schema.Tabler] interface {
	Create(T) error
	CreateInBatch([]T) (int64, error)
}

type Updater[T schema.Tabler] interface {
	Updates(any, ...gormx.Option) (int64, error)
	Update(T, ...gormx.Option) error
}

type Deleter[T schema.Tabler] interface {
	Delete(...gormx.Option) (int64, error)
}

type First[T schema.Tabler] interface {
	First(T, ...gormx.Option) error
}

type Find[T schema.Tabler] interface {
	Find(...gormx.Option) ([]T, error)
}

type Counter[T schema.Tabler] interface {
	Count(...gormx.Option) (int64, error)
}
