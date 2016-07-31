package transactions

import (
	"bitbucket.org/atticlab/go-smart-base/build"
	"bitbucket.org/atticlab/go-smart-base/keypair"
	"bitbucket.org/atticlab/go-smart-base/xdr"
	"bitbucket.org/atticlab/horizon/config"
	"bitbucket.org/atticlab/horizon/db2/core"
	"bitbucket.org/atticlab/horizon/db2/history"
	"bitbucket.org/atticlab/horizon/log"
	"bitbucket.org/atticlab/horizon/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreatePassiveOfferOpFrame(t *testing.T) {
	tt := test.Start(t).Scenario("base")
	defer tt.Finish()

	log.DefaultLogger.Entry.Logger.Level = log.DebugLevel

	historyQ := &history.Q{
		tt.HorizonRepo(),
	}
	coreQ := &core.Q{
		tt.CoreRepo(),
	}
	config := test.NewTestConfig()

	root := test.BankMasterSeed()

	validAsset := build.Asset{
		Code:   "UAH",
		Issuer: root.Address(),
	}
	invalidAsset := build.Asset{
		Code:   "USD",
		Issuer: root.Address(),
	}

	Convey("Invalid asset", t, func() {
		Convey("Invalid selling", func() {
			createPassiveOffer := build.CreatePassiveOffer(build.Rate{
				Price:   build.Price("10"),
				Selling: invalidAsset,
				Buying:  validAsset,
			}, build.Amount("1000"))
			checkInvalidAsset(createPassiveOffer, root, historyQ, coreQ, config)
		})
		Convey("Invalid buying", func() {
			createPassiveOffer := build.CreatePassiveOffer(build.Rate{
				Price:   build.Price("10"),
				Selling: validAsset,
				Buying:  invalidAsset,
			}, build.Amount("1000"))
			checkInvalidAsset(createPassiveOffer, root, historyQ, coreQ, config)
		})
	})
	Convey("Success", t, func() {
		createPassiveOffer := build.CreatePassiveOffer(build.Rate{
			Price:   build.Price("10"),
			Selling: validAsset,
			Buying:  validAsset,
		}, build.Amount("1000"))
		tx := build.Transaction(createPassiveOffer, build.Sequence{1}, build.SourceAccount{root.Address()})
		txE := tx.Sign(root.Seed()).E
		opFrame := NewOperationFrame(&txE.Tx.Operations[0], txE)
		isValid, err := opFrame.CheckValid(historyQ, coreQ, &config)
		So(err, ShouldBeNil)
		So(isValid, ShouldBeTrue)
		So(opFrame.GetResult().Result.MustTr().MustCreatePassiveOfferResult().Code, ShouldEqual, xdr.ManageOfferResultCodeManageOfferSuccess)
	})
}

func checkInvalidAsset(createPassiveOffer build.ManageOfferBuilder, root *keypair.Full, historyQ history.QInterface, coreQ core.QInterface, config config.Config) {
	tx := build.Transaction(createPassiveOffer, build.Sequence{1}, build.SourceAccount{root.Address()})
	txE := tx.Sign(root.Seed()).E
	opFrame := NewOperationFrame(&txE.Tx.Operations[0], txE)
	isValid, err := opFrame.CheckValid(historyQ, coreQ, &config)
	So(err, ShouldBeNil)
	So(isValid, ShouldBeFalse)
	So(opFrame.GetResult().Result.MustTr().MustCreatePassiveOfferResult().Code, ShouldEqual, xdr.ManageOfferResultCodeManageOfferMalformed)
	So(opFrame.GetResult().Info.GetError(), ShouldEqual, ASSET_NOT_ALLOWED.Error())
}