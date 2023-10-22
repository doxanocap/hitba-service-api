package api

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type BillingAPI struct {
	manager interfaces.IManager
	config  *model.Config
}

func NewBillingAPI(manager interfaces.IManager, config *model.Config) *BillingAPI {
	return &BillingAPI{
		manager: manager,
		config:  config,
	}
}

func (b *BillingAPI) Pay(ctx context.Context, payment *model.PaymentRequest) error {
	//data := &StudentsV2{}
	//
	//ids, err := json.Marshal(requestBodyV2{
	//	Ids: cfg.CensusIDs,
	//})
	//if err != nil {
	//	return nil, errs.Wrap("marshal request body: %v", err)
	//}
	//
	//body := bytes.NewReader(ids)
	//res, err := gohttp.NewRequest().
	//	SetURL(cfg.CensusStudentsUrl).
	//	SetMethod(gohttp.MethodPost).
	//	SetRequestBody(body).
	//	SetHeader("X-Auth-Token", cfg.CensusStudentsAuthToken).
	//	Execute(ctx)
	//if err != nil {
	//	return nil, errs.Wrap("census request: %v", err)
	//}
	//log.Info(fmt.Sprintf("CENSUS-CRON-JOB: request status: %s", res.Status))
	//
	//responseBody, err := io.ReadAll(res.Body)
	//if err != nil {
	//	return nil, errs.Wrap("read response body: %v", err)
	//}
	//if err = json.Unmarshal(responseBody, data); err != nil {
	//	return nil, errs.Wrap("unmarshal request: %v", err)
	//}
	//
	//return data, nil
	return nil
}
