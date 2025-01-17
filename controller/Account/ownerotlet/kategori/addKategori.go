package kategori

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/model/mobile"
	"hara-depo-proj/otp"
	"hara-depo-proj/util/customResponse"
	"io/ioutil"
	"net/http"
)

func AddKategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	kategory := mobile.Kategory{}
	kategoryUndifined := mobile.Kategory{}
	stok := mobile.Stok{}

	body, err1 := ioutil.ReadAll(r.Body)

	err1 = json.Unmarshal(body, &kategory)
	err2 := json.Unmarshal(body, &stok)
	_ = err2
	_ = err1

	defer r.Body.Close()

	kategoryUndifined.KodeKategory = 1
	kategoryUndifined.KodeUser = kategory.KodeUser
	kategoryUndifined.NamaKategory = "Undifined"

	if err := db.Save(&kategoryUndifined).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Save(&kategory).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	stok.IdKategori = kategory.KodeKategory
	stok.IdStok = otp.RandomString(5)

	if err := db.Save(&stok).Error; err != nil {
		customResponse.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	customResponse.RespondJSON(w, http.StatusOK, kategory)

}
