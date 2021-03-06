package api

import (
	"context"
	"github.com/mochahub/coinprice-scraper/config"
	"github.com/mochahub/coinprice-scraper/scraper/service/api/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"testing"
	"time"
)

func InjectDependencies(
	test interface{},
) error {
	fxApp := fx.New(
		GetAPIProviders(),
		fx.Provide(config.GetSecrets),
		fx.Invoke(
			test,
		),
	)
	if err := fxApp.Start(context.Background()); err != nil {
		return err
	}
	return fxApp.Stop(context.Background())
}

func TestExchangeClients(t *testing.T) {
	config.LoadEnv()
	err := InjectDependencies(func(clients ExchangeClients) {
		for i := range clients.Clients {
			exchangeClient := clients.Clients[i]
			pass := true
			pass = t.Run("TestGetAllOHLCMarketData", func(t *testing.T) {
				expectedLength := 1000 * time.Minute
				startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				endTime := startTime.Add(expectedLength - time.Minute)
				candleStickData, err := exchangeClient.GetAllOHLCMarketData(
					"BTC",
					"USDT",
					common.Minute,
					startTime,
					endTime,
				)
				assert.NoError(t, err)
				assert.NotEmpty(t, candleStickData)
				assert.Equal(t, int(expectedLength.Minutes()), len(candleStickData))
			}) && pass
			pass = t.Run("TestGetSupportedPairs", func(t *testing.T) {
				symbols, err := exchangeClient.GetSupportedPairs()
				assert.Nil(t, err)
				assert.NotEmpty(t, symbols)
			}) && pass
			pass = t.Run("TestGetExchangeIdentifier", func(t *testing.T) {
				identifier := exchangeClient.GetExchangeIdentifier()
				assert.NotEmpty(t, identifier)
			}) && pass
			assert.Equal(t, true, pass)
		}
	})
	assert.NoError(t, err)
}
