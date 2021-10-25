// Copyright 2021 Optakt Labs OÜ
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package mocks

import "testing"

type Generator struct {
	GetBalanceFunc       func(symbol string) ([]byte, error)
	GetStakedBalanceFunc func(symbol string) ([]byte, error)
	TokensDepositedFunc  func(symbol string) (string, error)
	TokensWithdrawnFunc  func(symbol string) (string, error)
	TransferTokensFunc   func(symbol string) ([]byte, error)
}

func BaselineGenerator(t *testing.T) *Generator {
	t.Helper()

	g := Generator{
		GetBalanceFunc: func(string) ([]byte, error) {
			return []byte(GenericAmount(0).String()), nil
		},
		GetStakedBalanceFunc: func(string) ([]byte, error) {
			return []byte(GenericAmount(0).String()), nil
		},
		TokensDepositedFunc: func(string) (string, error) {
			return string(GenericEventType(0)), nil
		},
		TokensWithdrawnFunc: func(string) (string, error) {
			return string(GenericEventType(1)), nil
		},
		TransferTokensFunc: func(string) ([]byte, error) {
			return GenericBytes, nil
		},
	}

	return &g
}

func (g *Generator) GetBalance(symbol string) ([]byte, error) {
	return g.GetBalanceFunc(symbol)
}

func (g *Generator) GetStakedBalance(symbol string) ([]byte, error) {
	return g.GetBalanceFunc(symbol)
}

func (g *Generator) TokensDeposited(symbol string) (string, error) {
	return g.TokensDepositedFunc(symbol)
}

func (g *Generator) TokensWithdrawn(symbol string) (string, error) {
	return g.TokensWithdrawnFunc(symbol)
}

func (g *Generator) TransferTokens(symbol string) ([]byte, error) {
	return g.TransferTokensFunc(symbol)
}
