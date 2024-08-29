package server

import "mezink/src/business/entity"

type HTTPRecordResp struct {
	HTTPCommonResp
	Records []entity.Record `json:"records"`
}

type HTTPCommonResp struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
