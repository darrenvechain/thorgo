package logs

import "github.com/darrenvechain/thorgo/thorest"

type TransfersQuerier struct {
	*Options

	client *thorest.Client
}

func NewTransfersQuerier(client *thorest.Client) *TransfersQuerier {
	return &TransfersQuerier{
		Options: &Options{},
		client:  client,
	}
}

func (tq *TransfersQuerier) Execute(criteria []thorest.TransferCriteria) ([]*thorest.TransferLog, error) {
	return tq.client.FilterTransfers(&thorest.TransferFilter{
		Criteria: &criteria,
		Range:    tq.rnge,
		Options:  tq.opts,
		Order:    tq.order,
	})
}
