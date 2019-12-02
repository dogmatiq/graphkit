package main

import (
	"fmt"

	"github.com/dogmatiq/configkit"
	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/dogma/fixtures"
	"github.com/dogmatiq/graphkit"
)

type SignUp struct{}
type CustomerSignedUp struct{}
type CreditAccountWithBonus struct{}
type AccountCreditedForBonus struct{}
type AccountCreditedForDeposit struct{}
type AccountDebitedForWithdrawal struct{}
type CheckBonusEligibility struct{}
type EligibleForBonus struct{}
type NotEligibleForBonus struct{}
type AnnualBonusDue struct{}

func main() {
	customer := &fixtures.Application{
		ConfigureFunc: func(c dogma.ApplicationConfigurer) {
			c.Identity("customer", "1ad3d548-f1f9-4a35-aaf8-e319476c4799")

			c.RegisterAggregate(&fixtures.AggregateMessageHandler{
				ConfigureFunc: func(c dogma.AggregateConfigurer) {
					c.Identity("customer", "bb29157b-a50b-486e-9e4e-fc1fb0931f6f")
					c.ConsumesCommandType(SignUp{})
					c.ProducesEventType(CustomerSignedUp{})
				},
			})

			c.RegisterProjection(&fixtures.ProjectionMessageHandler{
				ConfigureFunc: func(c dogma.ProjectionConfigurer) {
					c.Identity("customers", "5ce7b4c6-d047-4ec1-8e36-fb3124cf1e91")
					c.ConsumesEventType(CustomerSignedUp{})
				},
			})
		},
	}

	account := &fixtures.Application{
		ConfigureFunc: func(c dogma.ApplicationConfigurer) {
			c.Identity("account", "22310d8c-227c-4442-8b81-b662b72e8a21")

			c.RegisterAggregate(&fixtures.AggregateMessageHandler{
				ConfigureFunc: func(c dogma.AggregateConfigurer) {
					c.Identity("account", "2e9b8f26-aed7-4395-b6e6-41f5713ee8ba")
					c.ConsumesCommandType(CreditAccountWithBonus{})
					c.ProducesEventType(AccountCreditedForBonus{})
				},
			})

			c.RegisterProjection(&fixtures.ProjectionMessageHandler{
				ConfigureFunc: func(c dogma.ProjectionConfigurer) {
					c.Identity("balance", "cba96593-c718-4955-94e6-4554767b8e3f")
					c.ConsumesEventType(CustomerSignedUp{})
					c.ConsumesEventType(AccountCreditedForBonus{})
					c.ConsumesEventType(AccountCreditedForDeposit{})
					c.ConsumesEventType(AccountDebitedForWithdrawal{})
				},
			})

			c.RegisterProjection(&fixtures.ProjectionMessageHandler{
				ConfigureFunc: func(c dogma.ProjectionConfigurer) {
					c.Identity("ledger", "ff7e610d-2b0f-46f0-bde7-44dcc4590350")
					c.ConsumesEventType(AccountCreditedForBonus{})
					c.ConsumesEventType(AccountCreditedForDeposit{})
					c.ConsumesEventType(AccountDebitedForWithdrawal{})
				},
			})
		},
	}

	promo := &fixtures.Application{
		ConfigureFunc: func(c dogma.ApplicationConfigurer) {
			c.Identity("promo", "1c34ab15-85f7-4b4a-af62-f45170ed10a9")

			c.RegisterProcess(&fixtures.ProcessMessageHandler{
				ConfigureFunc: func(c dogma.ProcessConfigurer) {
					c.Identity("signup-bonus", "a53e314b-28c3-4cfb-9a3e-b3cc6f454bb4")
					c.ConsumesEventType(CustomerSignedUp{})
					c.ConsumesEventType(EligibleForBonus{})
					c.ConsumesEventType(NotEligibleForBonus{})
					c.ProducesCommandType(CreditAccountWithBonus{})
					c.ProducesCommandType(CheckBonusEligibility{})
				},
			})

			c.RegisterProcess(&fixtures.ProcessMessageHandler{
				ConfigureFunc: func(c dogma.ProcessConfigurer) {
					c.Identity("anniversary-bonus", "51394485-a6b9-41bb-9490-c9140097f2ea")
					c.ConsumesEventType(CustomerSignedUp{})
					c.ProducesCommandType(CreditAccountWithBonus{})
					c.SchedulesTimeoutType(AnnualBonusDue{})
				},
			})

			c.RegisterIntegration(&fixtures.IntegrationMessageHandler{
				ConfigureFunc: func(c dogma.IntegrationConfigurer) {
					c.Identity("rewards-program", "cd83a782-6a7f-4e10-90df-f672e2afb8cf")
					c.ConsumesCommandType(CheckBonusEligibility{})
					c.ProducesEventType(EligibleForBonus{})
					c.ProducesEventType(NotEligibleForBonus{})
				},
			})
		},
	}

	g, err := graphkit.Generate(
		configkit.FromApplication(promo),
		configkit.FromApplication(account),
		configkit.FromApplication(customer),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(g)
}
