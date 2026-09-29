package services

import (
	"iot-backend/internal/models"
	"iot-backend/internal/repositories/master"
)

func ApplyAddressOverride(
	instansiRepo *master.InstansiRepository,
	instansiID uint,
	override *models.AddressOverride,
) error {
	if override == nil {
		return nil
	}

	hasChange := (override.Alamat != nil && *override.Alamat != "") ||
		(override.KotaKab != nil && *override.KotaKab != "") ||
		(override.Provinsi != nil && *override.Provinsi != "")

	if !hasChange {
		return nil
	}

	instansi, err := instansiRepo.FindByID(instansiID)
	if err != nil {
		return err
	}

	if override.Alamat != nil && *override.Alamat != "" {
		instansi.Alamat = override.Alamat
	}
	if override.KotaKab != nil && *override.KotaKab != "" {
		instansi.KotaKab = override.KotaKab
	}
	if override.Provinsi != nil && *override.Provinsi != "" {
		instansi.Provinsi = override.Provinsi
	}

	return instansiRepo.Update(instansi)
}

func ApplyPicOverride(
	picRepo *master.PICRepository,
	picID uint,
	override *models.PicOverride,
) error {
	if override == nil || picID == 0 {
		return nil
	}

	hasChange := (override.NamaPIC != nil && *override.NamaPIC != "") ||
		(override.NoHP != nil && *override.NoHP != "")

	if !hasChange {
		return nil
	}

	pic, err := picRepo.FindByID(picID)
	if err != nil {
		return err
	}

	if override.NamaPIC != nil && *override.NamaPIC != "" {
		pic.NamaPIC = *override.NamaPIC
	}
	if override.NoHP != nil && *override.NoHP != "" {
		pic.NoHP = override.NoHP
	}

	return picRepo.Update(pic)
}