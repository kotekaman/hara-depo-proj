package pelanggan

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/util/customResponse"
	"log"
	"net/http"
	"strconv"
)

func ListPelanggan(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	kodeuser := vars["kodeuser"]
	filter := vars["filter"]
	sort := vars["sort"]
	var qsort string

	if sort == "abjad" {
		qsort = "pelanggan.nama desc"
	} else if sort == "piutang" {
		qsort = "utang desc"
	} else if sort == "transaksi" {
		qsort = "sum(total) desc"
	} else {
		qsort = "pelanggan.time_created desc"
	}

	listPenjualan := []mobile.ListPenjualan{}
	response := mobile.ResponseListPenjualan{}
	jumlahPiutang := mobile.TotalUtang{}

	dbOffset := page * 10

	if filter == "piutang" {
		if err := db.Table("pelanggan").Select("pelanggan.nama,pelanggan.id_pelanggan,sum(hutang.sisa_hutang) as utang").
			Joins("left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan").
			Joins("left join hutang on transaksi_uang.id_transaksi = hutang.id_transaksi").
			Group("pelanggan.nama,pelanggan.id_pelanggan").
			Where("pelanggan.kode_user=? AND utang > 0", kodeuser).
			Offset(dbOffset).
			Limit(10).
			Order(qsort).
			Find(&listPenjualan).Error; err != nil {
			log.Println(err.Error())
			//customResponse.RespondError(w, http.StatusInternalServerError, err.Error())

		}
	} else {
		if err := db.Table("pelanggan").Select(" pelanggan.nama,pelanggan.id_pelanggan,sum(hutang.sisa_hutang) as utang").
			Joins("left join transaksi_uang on pelanggan.id_pelanggan = transaksi_uang.id_pelanggan").
			Joins("left join hutang on transaksi_uang.id_hutang = hutang.id_hutang").
			Group("pelanggan.nama,pelanggan.id_pelanggan").
			Where("pelanggan.kode_user=?", kodeuser).
			Offset(dbOffset).
			Limit(10).
			Order(qsort).
			Find(&listPenjualan).Error; err != nil {
			//var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
			log.Println(err.Error())
			//customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		}
	}

	if err := db.Table("transaksi_uang").Select("sum(hutang.sisa_hutang) as total_piutang").
		Joins("left join hutang on transaksi_uang.id_hutang = hutang.id_hutang").
		Where("transaksi_uang.kode_user=?", kodeuser).
		First(&jumlahPiutang).Error; err != nil {
		log.Println(err.Error())
		//customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	response.Penjualan = listPenjualan
	response.TotalPiutang = jumlahPiutang

	customResponse.RespondJSON(w, 202, response)
}
