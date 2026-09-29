package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories/master"
)

var (
	ErrPicIDTidakDitemukan = errors.New("pic_id tidak ditemukan")
	ErrPicWajibDiisi       = errors.New("pic_id (pilih dari data yang sudah ada) atau pic (isi manual, minimal nama_pic) wajib diisi")
)

func ResolvePicID(
	picRepo *master.PICRepository,
	instansiID uint,
	picID uint,
	manual *models.ManualPICRequest,
) (uint, error) {
	if picID > 0 {
		existing, err := picRepo.FindByID(picID)
		if err != nil || existing == nil {
			return 0, ErrPicIDTidakDitemukan
		}
		return picID, nil
	}

	if manual == nil || strings.TrimSpace(manual.NamaPIC) == "" {
		return 0, ErrPicWajibDiisi
	}

	newPic := &models.MasterPIC{
		InstansiID: &instansiID,
		NamaPIC:    strings.TrimSpace(manual.NamaPIC),
		NoHP:       manual.NoHP,
		NIK:        manual.NIK,
		Email:      manual.Email,
	}

	if err := picRepo.Create(newPic); err != nil {
		return 0, err
	}

	return newPic.ID, nil
}