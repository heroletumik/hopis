package resource

import (
	"fmt"

	"bitbucket.org/atticlab/horizon/db2/core"
	"bitbucket.org/atticlab/horizon/db2/history"
	"bitbucket.org/atticlab/horizon/httpx"
	"bitbucket.org/atticlab/horizon/render/hal"
	"golang.org/x/net/context"
)

// Populate fills out the resource's fields
func (this *Account) Populate(
	ctx context.Context,
	ca core.Account,
	cd []core.AccountData,
	cs []core.Signer,
	ct []core.Trustline,
	ha history.Account,
) (err error) {
	this.ID = ca.Accountid
	this.PT = ha.PagingToken()
	this.AccountID = ca.Accountid
	this.Sequence = ca.Seqnum
	this.SubentryCount = ca.Numsubentries
	this.InflationDestination = ca.Inflationdest.String
	this.HomeDomain = ca.HomeDomain.String
	this.populateType(ca)
	this.Flags.Populate(ca)
	this.Thresholds.Populate(ca)

	// populate balances
	this.Balances = make([]Balance, len(ct)+1)
	for i, tl := range ct {
		err = this.Balances[i].Populate(ctx, tl)
		if err != nil {
			return
		}
	}

	// add native balance
	err = this.Balances[len(this.Balances)-1].PopulateNative(ca.Balance)
	if err != nil {
		return
	}

	// populate data
	this.Data = make(map[string]string)
	for _, d := range cd {
		this.Data[d.Key] = d.Value
	}

	// populate signers
	this.Signers = make([]Signer, len(cs)+1)
	for i, s := range cs {
		this.Signers[i].Populate(ctx, s)
	}

	this.Signers[len(this.Signers)-1].PopulateMaster(ca)

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	self := fmt.Sprintf("/accounts/%s", ha.Address)
	this.Links.Self = lb.Link(self)
	this.Links.Transactions = lb.PagedLink(self, "transactions")
	this.Links.Operations = lb.PagedLink(self, "operations")
	this.Links.Payments = lb.PagedLink(self, "payments")
	this.Links.Effects = lb.PagedLink(self, "effects")
	this.Links.Offers = lb.PagedLink(self, "offers")

	return
}

func (this *Account) populateType(row core.Account) {
	var ok bool
	this.TypeI = int32(row.AccountType)
	this.Type, ok = AccountTypeNames[row.AccountType]

	if !ok {
		this.Type = "unknown"
	}
}
