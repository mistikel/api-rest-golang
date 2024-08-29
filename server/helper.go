package server

import (
	"encoding/json"
	"mezink/src/business/entity"
	errors "mezink/stdlib/error"
	"mezink/stdlib/log"
	"net/http"
)

func (e *REST) httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp interface{}) {
	var (
		raw []byte
		err error
	)
	switch data := resp.(type) {
	case []entity.Record:
		resp := &HTTPRecordResp{
			Records: data,
		}
		raw, err = json.Marshal(resp)
	default:
		e.httpRespError(w, r, errors.NewAppError(100, "Undefined response", 404, errors.ErrNotFound))
		return
	}

	if err != nil {
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *REST) httpRespError(w http.ResponseWriter, r *http.Request, err error) {
	appErr, ok := err.(*errors.AppError)
	if !ok {
		appErr = &errors.AppError{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	jsonErrResp := &HTTPCommonResp{
		Code:    appErr.Code,
		Message: appErr.Error(),
	}
	raw, err := json.Marshal(jsonErrResp)
	if err != nil {
		appErr.StatusCode = http.StatusInternalServerError
	}

	log.ErrContext(r.Context(), appErr.Error())
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(appErr.StatusCode)
	_, _ = w.Write(raw)
}
