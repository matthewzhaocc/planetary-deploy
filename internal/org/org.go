package org

import (
	"encoding/json"
	"net/http"

	"github.com/matthewzhaocc/planetary-deploy/internal/types"

	"gorm.io/gorm"
)

type Org struct {
	gorm.Model
	name string
}

type NewOrg struct {
	Name string `json:"name"`
}

func AddOrg(serv types.Server, newOrg NewOrg) {
	org := Org{name: newOrg.Name}
	serv.DB.Create(&org)
}

func AddOrgRoute(serv types.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var org NewOrg
		json.NewDecoder(r.Body).Decode(&org)
		AddOrg(serv, org)
	}
}
