package fotografercontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/DProject89/cmsfoto/libraries"

	"github.com/DProject89/cmsfoto/models"

	"github.com/DProject89/cmsfoto/entities"
)

var validation = libraries.NewValidation()
var fotograferModel = models.NewFotograferModel()

func Index(response http.ResponseWriter, request *http.Request) {

	fotografer, _ := fotograferModel.FindAll()

	data := map[string]interface{}{
		"fotografer": fotografer,
	}

	temp, err := template.ParseFiles("views/fotografer/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/fotografer/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var fotografer entities.Fotografer
		fotografer.NamaLengkap = request.Form.Get("nama_lengkap")
		fotografer.NIK = request.Form.Get("nik")
		fotografer.JenisKelamin = request.Form.Get("jenis_kelamin")
		fotografer.TempatLahir = request.Form.Get("tempat_lahir")
		fotografer.TanggalLahir = request.Form.Get("tanggal_lahir")
		fotografer.Alamat = request.Form.Get("alamat")
		fotografer.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(fotografer)

		if vErrors != nil {
			data["fotografer"] = fotografer
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data fotografer berhasil disimpan"
			fotograferModel.Create(fotografer)
		}

		temp, _ := template.ParseFiles("views/fotografer/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var fotografer entities.Fotografer
		fotograferModel.Find(id, &fotografer)

		data := map[string]interface{}{
			"fotografer": fotografer,
		}

		temp, err := template.ParseFiles("views/fotografer/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var fotografer entities.Fotografer
		fotografer.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		fotografer.NamaLengkap = request.Form.Get("nama_lengkap")
		fotografer.NIK = request.Form.Get("nik")
		fotografer.JenisKelamin = request.Form.Get("jenis_kelamin")
		fotografer.TempatLahir = request.Form.Get("tempat_lahir")
		fotografer.TanggalLahir = request.Form.Get("tanggal_lahir")
		fotografer.Alamat = request.Form.Get("alamat")
		fotografer.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(fotografer)

		if vErrors != nil {
			data["fotografer"] = fotografer
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data fotografer berhasil diperbarui"
			fotograferModel.Update(fotografer)
		}

		temp, _ := template.ParseFiles("views/fotografer/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	fotograferModel.Delete(id)

	http.Redirect(response, request, "/fotografer", http.StatusSeeOther)
}
