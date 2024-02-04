package app

import "test/integration"

type SetupIntegration struct {
	TransactionIntegration integration.TransactionIntegration
}

func (i *SetupIntegration) Setup() {
	i.TransactionIntegration = integration.NewTransactionIntegration()
}
