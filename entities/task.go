package entities

type Task struct {
	Id         int64
	TaskDetail string `validate:"required" label:"Detail Tugas"`
	Assigment  string `validate:"required" label:"Di Tugaskan Kepada"`
	Deadline   string `validate:"required" label:"Tanggal Tugas Berakhir"`
	Status     string `validate:"required" label:"Status tugas (Belum Selesai/Selesai)"`
}
