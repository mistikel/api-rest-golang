package server

import (
	"database/sql"
	"mezink/src/business/entity"
	"mezink/src/business/usecase"
	errors "mezink/stdlib/error"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
)

var IsShuttingDown = false

type Handler interface {
	CreateRouter() *mux.Router
	HealthCheck(w http.ResponseWriter, r *http.Request)
	GetRecord(w http.ResponseWriter, r *http.Request)
}

func NewHandler(uc *usecase.Usecase, db *sql.DB) Handler {
	return &REST{
		uc:       uc,
		param:    schema.NewDecoder(),
		db:       db,
		validate: validator.New(),
	}
}

type REST struct {
	uc       *usecase.Usecase
	param    *schema.Decoder
	db       *sql.DB
	validate *validator.Validate
}

func (c *REST) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if IsShuttingDown {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Preparing to shutdown"))
		return
	}

	response := "OK"
	status := http.StatusOK
	if err := c.db.Ping(); err != nil {
		response = "Not OK"
		status = http.StatusBadGateway
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}

func (c *REST) CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", c.HealthCheck).Methods("GET")
	r.HandleFunc("/record", c.GetRecord).Methods("GET")
	return r
}

// @Summary Get all records
// @Description Retrieve a list of records
// @Tags records
// @Accept  json
// @Produce  json
// @Param startDate query string false "The start date for filtering records"
// @Param endDate query string false "The end date for filtering records"
// @Param minCount query integer false "Minimum count for filtering records"
// @Param maxCount query integer false "Maximum count for filtering records"
// @Success 200 {array} HTTPRecordResp
// @Failure 400 {object} HTTPCommonResp
// @Failure 500 {object} HTTPCommonResp
// @Router /records [get]
func (c *REST) GetRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var param entity.RecordParam
	if err := c.param.Decode(&param, r.URL.Query()); err != nil {
		c.httpRespError(w, r, errors.NewAppError(errors.CodeHTTPParamDecode, "Failed to decode param", http.StatusBadRequest, err))
		return
	}

	if err := c.validate.Struct(param); err != nil {
		c.httpRespError(w, r, errors.NewAppError(errors.CodeHTTPParamValidate, "Failed to validate param", http.StatusBadRequest, err))
		return
	}

	result, err := c.uc.Record.GetRecords(ctx, param)
	if err != nil {
		c.httpRespError(w, r, err)
		return
	}

	c.httpRespSuccess(w, r, http.StatusOK, result)
}
