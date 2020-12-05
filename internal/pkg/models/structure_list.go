package models

type StructureList struct {
	Paths []string
}

func NewStructureList(paths []string) *StructureList {
	return &StructureList{
		Paths: paths,
	}
}
