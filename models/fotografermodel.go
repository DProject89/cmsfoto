package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DProject89/gocrud/config"
	"github.com/DProject89/gocrud/entities"
)

type FotograferModel struct {
	conn *sql.DB
}

func NewFotograferModel() *FotograferModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &FotograferModel{
		conn: conn,
	}
}

func (p *FotograferModel) FindAll() ([]entities.Fotografer, error) {

	rows, err := p.conn.Query("select * from fotografer")
	if err != nil {
		return []entities.Fotografer{}, err
	}
	defer rows.Close()

	var dataFotografer []entities.Fotografer
	for rows.Next() {
		var fotografer entities.Fotografer
		rows.Scan(
			&fotografer.Id,
			&fotografer.NamaLengkap,
			&fotografer.NIK,
			&fotografer.JenisKelamin,
			&fotografer.TempatLahir,
			&fotografer.TanggalLahir,
			&fotografer.Alamat,
			&fotografer.NoHp)

		if fotografer.JenisKelamin == "1" {
			fotografer.JenisKelamin = "Laki-laki"
		} else {
			fotografer.JenisKelamin = "Perempuan"
		}
		// 2006-01-02 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2006-01-02", fotografer.TanggalLahir)
		// 02-01-2006 => dd-mm-yyyy
		fotografer.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataFotografer = append(dataFotografer, fotografer)
	}

	return dataFotografer, nil

}

func (p *FotograferModel) Create(fotografer entities.Fotografer) bool {

	result, err := p.conn.Exec("insert into fotografer (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)",
		fotografer.NamaLengkap, fotografer.NIK, fotografer.JenisKelamin, fotografer.TempatLahir, fotografer.TanggalLahir, fotografer.Alamat, fotografer.NoHp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *FotograferModel) Find(id int64, fotografer *entities.Fotografer) error {

	return p.conn.QueryRow("select * from fotografer where id = ?", id).Scan(
		&fotografer.Id,
		&fotografer.NamaLengkap,
		&fotografer.NIK,
		&fotografer.JenisKelamin,
		&fotografer.TempatLahir,
		&fotografer.TanggalLahir,
		&fotografer.Alamat,
		&fotografer.NoHp)
}

func (p *FotograferModel) Update(fotografer entities.Fotografer) error {

	_, err := p.conn.Exec(
		"update fotografer set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
		fotografer.NamaLengkap, fotografer.NIK, fotografer.JenisKelamin, fotografer.TempatLahir, fotografer.TanggalLahir, fotografer.Alamat, fotografer.NoHp, fotografer.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *FotograferModel) Delete(id int64) {
	p.conn.Exec("delete from fotografer where id = ?", id)
}
