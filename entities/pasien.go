// package entities

// type Pasien struct {
// 	Id           int64
// 	NamaLengkap  string `validate:"required"`
// 	NIK          string `validate:"required"`
// 	JenisKelamin string `validate:"required"`
// 	TempatLahir  string `validate:"required"`
// 	TanggalLahir string `validate:"required"`
// 	Alamat       string `validate:"required"`
// 	NoHp         string `validate:"required"`
// }

package entities

type Pasien struct {
	Id           int64
	NamaLengkap  string `validate:"required" label:"Nama Lengkap"`
	NIK          string `validate:"required"`
	JenisKelamin string `validate:"required" label:"Jenis Kelamin"`
	TempatLahir  string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir string `validate:"required" label:"Tanggal Lahir "`
	Alamat       string `validate:"required"`
	NoHp         string `validate:"required" label:"no.Hp"`
}
