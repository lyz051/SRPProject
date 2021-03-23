package models

type Exist interface {
	ExistByID(int, ...interface{}) error
	ExistByName(string, ...interface{}) error
}

type Add interface {
	Add(map[string]interface{}) error
}

type Update interface {
	Update(int, interface{}) error
}

type Get interface {
	getByID(int, ...interface{}) error
	getByName(string, ...interface{}) error
}

type Delete interface {
	deleteByID(int) error
	deleteByName(string) error
}
