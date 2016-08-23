package horizon

import (
	"encoding/json"
	"bitbucket.org/atticlab/horizon/db2"
	"bitbucket.org/atticlab/horizon/db2/history"
	"bitbucket.org/atticlab/horizon/render/hal"
	"bitbucket.org/atticlab/horizon/render/sse"
	"bitbucket.org/atticlab/horizon/resource"
)

// PaymentsIndexAction returns a paged slice of payments based upon the provided
// filters
type PaymentsIndexAction struct {
	Action
	LedgerFilter      int32
	AccountFilter     string
	MultiAccountFilter	[]string
	TransactionFilter string
	PagingParams      db2.PageQuery
	Records           []history.Operation
	Page              hal.Page
}

// JSON is a method for actions.JSON
func (action *PaymentsIndexAction) JSON() {
	action.Do(action.loadParams, action.loadRecords, action.loadPage)
	action.Do(func() {
		hal.Render(action.W, action.Page)
	})
}

// SSE is a method for actions.SSE
func (action *PaymentsIndexAction) SSE(stream sse.Stream) {
	action.Setup(action.loadParams)
	action.Do(
		action.loadRecords,
		func() {
			stream.SetLimit(int(action.PagingParams.Limit))
			records := action.Records[stream.SentCount():]

			for _, record := range records {
				res, err := resource.NewOperation(action.Ctx, record)

				if err != nil {
					stream.Err(action.Err)
					return
				}

				stream.Send(sse.Event{
					ID:   res.PagingToken(),
					Data: res,
				})
			}
		})
}

func (action *PaymentsIndexAction) loadParams() {
	action.ValidateCursorAsDefault()
	action.AccountFilter = action.GetString("account_id")
	addressesStr := action.GetString("multi_accounts")
	if action.AccountFilter == "" && addressesStr != "" {
		var addresses []string 
		err := json.Unmarshal([]byte(addressesStr),&addresses)
		if err == nil{
			action.MultiAccountFilter = addresses
		}
	}
	action.LedgerFilter = action.GetInt32("ledger_id")
	action.TransactionFilter = action.GetString("tx_id")
	action.PagingParams = action.GetPageQuery()
}

func (action *PaymentsIndexAction) loadRecords() {
	q := action.HistoryQ()
	ops := q.Operations().OnlyPayments()

	switch {
	case action.AccountFilter != "":
		ops.ForAccount(action.AccountFilter)
	case action.MultiAccountFilter != nil:
		ops.ForAccounts(action.MultiAccountFilter)
	case action.LedgerFilter > 0:
		ops.ForLedger(action.LedgerFilter)
	case action.TransactionFilter != "":
		ops.ForTransaction(action.TransactionFilter)
	}

	action.Err = ops.Page(action.PagingParams).Select(&action.Records)
}

func (action *PaymentsIndexAction) loadPage() {
	for _, record := range action.Records {
		var res hal.Pageable
		res, action.Err = resource.NewOperation(action.Ctx, record)
		if action.Err != nil {
			return
		}
		action.Page.Add(res)
	}

	action.Page.BaseURL = action.BaseURL()
	action.Page.BasePath = action.Path()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}
