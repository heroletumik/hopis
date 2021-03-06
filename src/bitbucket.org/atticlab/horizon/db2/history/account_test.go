package history

import (
	"testing"

	"bitbucket.org/atticlab/horizon/test"
)

func _TestAccountQueries(t *testing.T) {
	tt := test.Start(t).Scenario("base")
	defer tt.Finish()
	q := &Q{tt.HorizonRepo()}

	// Test Accounts()
	acs := []Account{}
	err := q.Accounts().Select(&acs)

	if tt.Assert.NoError(err) {
		tt.Assert.Len(acs, 4)
	}
}
