package models

type Exist interface {
	existByID(int, ...interface{}) error
	existByName(string, ...interface{}) error
}

type Add interface {
	add(map[string]interface{}) error
}

type Update interface {
	update(int, interface{}) error
}

type Get interface {
	getByID(int, ...interface{}) error
	getByName(string, ...interface{}) error
}

type Delete interface {
	deleteByID(int) error
	deleteByName(string) error
}
