package logs

import "github.com/darrenvechain/thorgo/thorest"

type TransfersFilterer struct {
	*Options
	criteria []thorest.TransferCriteria
	client   *thorest.Client
}

func NewTransfersFilterer(client *thorest.Client) *TransfersFilterer {
	return &TransfersFilterer{
		Options: &Options{},
		client:  client,
	}
}

func (t *TransfersFilterer) Criteria(criteria ...thorest.TransferCriteria) *TransfersFilterer {
	if t.criteria == nil {
		t.criteria = make([]thorest.TransferCriteria, 0)
	}
	t.criteria = append(t.criteria, criteria...)
	return t
}

func (t *TransfersFilterer) Execute() ([]*thorest.TransferLog, error) {
	return t.client.FilterTransfers(&thorest.TransferFilter{
		Criteria: &t.criteria,
		Range:    t.rnge,
		Options:  t.opts,
		Order:    t.order,
	})
}
