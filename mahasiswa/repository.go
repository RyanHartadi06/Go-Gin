package mahasiswa

import "gorm.io/gorm"

type Repository interface {
	Simpan(mahasiswa Mahasiswa) (Mahasiswa , error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r* repository) Simpan(mahasiswa Mahasiswa) (Mahasiswa , error) {
	err := r.db.Create(&mahasiswa).Error

	if err != nil {
		return 	mahasiswa , err
	}
	return mahasiswa , nil
}