package mahasiswa

type Service interface {
	Register(input RegisterMahasiswaInput) (Mahasiswa , error)
}

type service struct {
	repository Repository
}

func NewServices(repository Repository) *service {
	return &service{repository: repository}
}

func(s* service) Register(input RegisterMahasiswaInput) (Mahasiswa , error) {
	mahasiswa := Mahasiswa{}

	mahasiswa.Nama = input.Nama
	mahasiswa.Nim = input.Nim

	newMahasiswa , err := s.repository.Simpan(mahasiswa)
	if err != nil {
		return newMahasiswa , err
	}

	return newMahasiswa, nil
}