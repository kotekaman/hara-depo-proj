package app

func (app *App) setRouters() {
	// Routing for handling the projects
	app.Post("/v1/otlet/otp", app.createNewUser)
	app.Post("/v1/otlet/otpconfirm", app.ConfirmOtp)
	app.Post("/v1/otlet/addbarang", app.AddBarang)
	app.GetQueryBarangSort("/v1/otlet/listbarang", app.ListBarang)
	app.GetQueryBarangSort("/v1/otlet/listtransaksi", app.ListTransaksi)
	app.GetQueryBarang("/v1/otlet/listkategori", app.ListKategori)
	app.GetQueryBarang("/v1/otlet/listpersediaan", app.ListPersediaanBarang)
	app.GetBarangInformation("/v1/otlet/baranginfo", app.BarangInformations)
	app.Post("/v1/otlet/registrasi", app.RegisterUser)
	app.Post("/v1/otlet/addkategori", app.AddKategori)
	app.Post("/v1/otlet/addsuplier", app.AddSuplier)
	app.Post("/v1/otlet/addpelanggan", app.AddPelanggan)
	app.Post("/v1/otlet/pembelianbarang", app.PembelianBarang)
	app.Post("/v1/otlet/penjualanbarang", app.JualanBarang)
	app.DeleteBarang("/v1/otlet/deletebarang", app.DeleteBarangI)
	app.DeleteKategory("/v1/otlet/deletekategori", app.DeleteKategoryI)
	app.UpdateBarang("/v1/otlet/updatebarang", app.UpdateBarangI)
	app.UpdateBarang("/v1/otlet/updatekategori", app.UpdateKategoryI)
}
