package models
/* would be deleted */

type GeneratorItf interface {
	GenerateDirectory(path string) error
	GenerateFile(path string, content []byte) error
}
