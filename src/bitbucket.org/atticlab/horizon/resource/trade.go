package resource

import (
	"errors"

	"bitbucket.org/atticlab/horizon/db2/history"
	"bitbucket.org/atticlab/horizon/httpx"
	"bitbucket.org/atticlab/horizon/render/hal"
	"golang.org/x/net/context"
)

// Populate fills out the details
func (res *Trade) Populate(ctx context.Context, row history.Effect) (err error) {
	if row.Type != history.EffectTrade {
		err = errors.New("invalid effect; not a trade")
		return
	}
	row.UnmarshalDetails(res)
	res.ID = row.PagingToken()
	res.PT = row.PagingToken()
	res.Buyer = row.Account

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	res.Links.Self = lb.Link("/accounts", res.Seller)
	res.Links.Seller = lb.Link("/accounts", res.Seller)
	res.Links.Buyer = lb.Link("/accounts", res.Buyer)
	return
}

// PagingToken implementation for hal.Pageable
func (res Trade) PagingToken() string {
	return res.PT
}
