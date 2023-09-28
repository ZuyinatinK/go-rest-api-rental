package tenantcontroller

import (
	"encoding/json"
	"errors"
	"go-rest-api-rental/config"
	"go-rest-api-rental/helper"
	"go-rest-api-rental/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tenants []models.Tenant
	var tenantResponse []models.TenantResponse

	if err := config.DB.Joins("Property").Find(&tenants).Find(&tenantResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Tenant's", tenantResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var tenant models.Tenant

	if err := json.NewDecoder(r.Body).Decode(&tenant); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// Cek Author
	var property models.Property
	if err := config.DB.First(&property, tenant.PropertyID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Property Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&tenant).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create Tenant", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var tenant models.Tenant
	var tenantResponse models.TenantResponse

	if err := config.DB.Joins("Property").First(&tenant, id).First(&tenantResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Tenant Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Tenant", tenantResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var tenant models.Tenant

	if err := config.DB.First(&tenant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Tenant Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var tenantPayload models.Tenant
	if err := json.NewDecoder(r.Body).Decode(&tenantPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var property models.Property
	if tenantPayload.PropertyID != 0 {
		if err := config.DB.First(&property, tenantPayload.PropertyID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Property Not Found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&tenantPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Update Tenant", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var tenant models.Tenant
	res := config.DB.Delete(&tenant, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Tenant Not Found", nil)
		return
	}
	helper.Response(w, 200, "Success Delete Tenant", nil)
}
