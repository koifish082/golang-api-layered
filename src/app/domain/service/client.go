package service

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/koifish082/golang-api-layered/src/app/config"
	serviceLog "github.com/koifish082/golang-api-layered/src/app/domain/log"
	"github.com/koifish082/golang-api-layered/src/app/domain/model"
	"github.com/koifish082/golang-api-layered/src/app/library/log"
	"github.com/spf13/viper"
)

var (
	client *resty.Client
	once   sync.Once
)

func getAPIClient() *resty.Client {
	timeout := viper.GetInt(config.EnvAPIRequestTimeout)
	retryCount := viper.GetInt(config.EnvAPIRequestRetryCount)
	logger := serviceLog.ClientLogger()
	once.Do(func() {
		client = resty.New().
			SetLogger(logger).
			SetTimeout(time.Duration(timeout) * time.Millisecond).
			AddRetryCondition(func(r *resty.Response, e error) bool {
				return r.StatusCode() >= http.StatusInternalServerError
			}).
			SetRetryWaitTime(500 * time.Millisecond).
			SetRetryCount(retryCount)
	})
	return client
}

func createCommonHeaders() map[string]string {
	return map[string]string{
		"Accept": "application/vnd.github.v3+json",
	}
}

func loggingRequest(endpoint string, headers, query map[string]string) {
	h, _ := json.Marshal(headers)
	qry, _ := json.Marshal(query)
	log.Info("Request URL:", endpoint,
		", header:", string(h),
		", query:", string(qry),
	)
}

func loggingResponse(res *resty.Response) {
	body := string(res.Body())
	log.Infof("Response: Method=[%s] Time=[%s] Status=[%s] PATH[%s] Body=[%s]",
		res.Request.Method, res.Time(), res.Status(), res.Request.RawRequest.URL.Path, body)
}

func checkResponseWithOptions(res *resty.Response) *model.ServiceError {
	loggingResponse(res)
	if res.StatusCode() >= 200 && res.StatusCode() < 400 {
		return nil
	}
	serverErrResp, err := model.UnmarshalServerErrorResponse(res.Body())
	if err != nil {
		return model.NewServiceError(err)
	}

	return model.NewServiceErrorByServerError(serverErrResp, res.StatusCode())
}

func get(endpoint string, headers, queries map[string]string) (*resty.Response, *model.ServiceError) {
	loggingRequest(endpoint, headers, queries)

	resp, err := getAPIClient().R().
		SetHeaders(headers).
		SetQueryParams(queries).
		Get(endpoint)
	if err != nil {
		return nil, model.NewServiceError(err)
	}

	e := checkResponseWithOptions(resp)
	if e != nil {
		return nil, e
	}

	return resp, nil
}
