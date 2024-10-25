package model

import (
	"gitlab.com/digeon-inc/japan-association-for-clinical-engineers/e-privado/api/domain/entconst"
	"gitlab.com/digeon-inc/japan-association-for-clinical-engineers/e-privado/api/domain/entity"
)

type model[T any] interface {
	Entity() T
}

// ToEntities converts slice of models to slice of entities.
func ToEntities[Entity any, Model model[Entity]](models []Model) []Entity {
	ret := make([]Entity, len(models))
	for i, model := range models {
		ret[i] = model.Entity()
	}
	return ret
}

func (u User) Entity() entity.User {
	return entity.User{
		ID:         u.ID,
		StudentID:      u.StudentID,
		IdmUniv:        u.IdmUniv,
		IdmBus:         u.IdmBus,
	}
}

func (f File) Entity() entity.File {
	return entity.File{
		FileID:      f.FileID,
		FileKind:    entconst.FileKind(f.FileKind),
		FileName:    f.FileName,
		FileURL:     "",
		ContentType: f.ContentType,
		FileSize:    f.FileSize,
		CreatedUser: f.CreatedUser.Entity(),
		UpdatedUser: f.UpdatedUser.Entity(),
	}
}
