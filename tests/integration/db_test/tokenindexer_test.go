// +build integration

package db_test

import (
	"context"
	assert "github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"sort"
	"testing"
)

func Test_AddNewAssets(t *testing.T) {
	type testsStruct struct {
		Name         string
		Assets       []models.Asset
		AssetsIDs    []string
		WantedErr    error
		WantedAssets []models.Asset
	}
	tests := []testsStruct{
		{
			Name: "Normal case",
			Assets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 18,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 18,
					Name:     "B",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
			},
			AssetsIDs: []string{"c714_a", "c714_b"},
			WantedErr: nil,
			WantedAssets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 18,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 18,
					Name:     "B",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
			},
		},
		{
			Name: "Case with new tokens and old tokens",
			Assets: []models.Asset{
				{
					Asset:    "c714_c",
					Decimals: 18,
					Name:     "C",
					Symbol:   "FFF",
					Type:     "ERC20",
				},
				{
					Asset:    "c714_d",
					Decimals: 18,
					Name:     "D",
					Symbol:   "RRR",
					Type:     "TRC20",
				},
			},
			AssetsIDs: []string{"c714_a", "c714_b", "c714_c", "c714_d"},
			WantedErr: nil,
			WantedAssets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 18,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 18,
					Name:     "B",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_c",
					Decimals: 18,
					Name:     "C",
					Symbol:   "FFF",
					Type:     "ERC20",
				},
				{
					Asset:    "c714_d",
					Decimals: 18,
					Name:     "D",
					Symbol:   "RRR",
					Type:     "TRC20",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := database.AddNewAssets(tt.Assets, context.Background())
			assert.Equal(t, tt.WantedErr, err)
			assets, err := database.GetAssetsByIDs(tt.AssetsIDs, context.Background())
			assert.Nil(t, err)
			sort.Slice(tt.WantedAssets, func(i, j int) bool {
				return tt.WantedAssets[i].Asset > tt.WantedAssets[j].Asset
			})
			sort.Slice(assets, func(i, j int) bool {
				return assets[i].Asset > assets[j].Asset
			})
			for i, a := range assets {
				assert.Equal(t, tt.WantedAssets[i].Asset, a.Asset)
				assert.Equal(t, tt.WantedAssets[i].Name, a.Name)
				assert.Equal(t, tt.WantedAssets[i].Symbol, a.Symbol)
				assert.Equal(t, tt.WantedAssets[i].Type, a.Type)
				assert.Equal(t, tt.WantedAssets[i].Decimals, a.Decimals)
			}
		})
	}
}
