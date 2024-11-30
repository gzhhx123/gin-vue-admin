package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	bagiqueDb := global.GetGlobalDBByDBName("bagique")
	bagiqueDb.AutoMigrate(bagique.Brand{}, bagique.Product{}, bagique.Company{}, bagique.Seller{}, bagique.Evaluate{}, bagique.EvaluatePrice{}, bagique.Purchase{}, bagique.TrackCompany{})
	return nil
}
