package main

import (
	"gorm-gen-peoject/db"
	"gorm-gen-peoject/model"
	"gorm.io/gen"
)

func table2struct() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
	})

	g.UseDB(db.DB)

	g.ApplyBasic(
		g.GenerateModel("users"),
	)

	g.ApplyInterface(
		func(method model.TUserMethod) {},
		g.GenerateModel("users"),
	)

	// execute the action of code generation
	g.Execute()
}