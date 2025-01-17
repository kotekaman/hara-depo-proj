package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hara-depo-proj/config"
	"hara-depo-proj/model/mobile"
	"log"
	"net/http"
)

func (app *App) Initialize() {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./dist/")))

	db, err := gorm.Open("postgres", "host="+baseConfig.Postgres.Url+" port="+baseConfig.Postgres.Port+""+
		" user="+baseConfig.Postgres.User+" dbname="+baseConfig.Postgres.Db+" password="+baseConfig.Postgres.Password+
		" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	app.TbUserHara = mobile.DBMigrationAccount(db, &mobile.Userhara{})
	app.TbBarang = mobile.DBMigrationAccount(db, &mobile.BarangOtlet{})
	app.TbUserOtlet = mobile.DBMigrationAccount(db, mobile.UserOtlet{})
	app.TbTokoOtlet = mobile.DBMigrationAccount(db, mobile.Toko{})
	app.TbKategori = mobile.DBMigrationAccount(db, mobile.Kategory{})
	app.TbStok = mobile.DBMigrationAccount(db, mobile.Stok{})
	app.TbPelanggan = mobile.DBMigrationAccount(db, mobile.Pelanggan{})
	app.TbSuplier = mobile.DBMigrationAccount(db, mobile.Suplier{})
	app.TbTransaksiBarang = mobile.DBMigrationAccount(db, mobile.TransaksiBarang{})
	app.TbTransaksiUang = mobile.DBMigrationAccount(db, mobile.TransaksiUang{})
	app.TbTransaksiUang = mobile.DBMigrationAccount(db, mobile.Hutang{})
	app.Router = mux.NewRouter()

	app.setRouters()
	app.Router.PathPrefix("/swaggerui/").Handler(sh)
}
