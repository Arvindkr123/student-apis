package storage

type storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}
