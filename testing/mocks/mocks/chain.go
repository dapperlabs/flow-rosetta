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

import (
	"testing"

	"github.com/onflow/flow-go/model/flow"
)

type Chain struct {
	RootFunc         func() (uint64, error)
	HeaderFunc       func(height uint64) (*flow.Header, error)
	CommitFunc       func(height uint64) (flow.StateCommitment, error)
	CollectionsFunc  func(height uint64) ([]*flow.LightCollection, error)
	GuaranteesFunc   func(height uint64) ([]*flow.CollectionGuarantee, error)
	TransactionsFunc func(height uint64) ([]*flow.TransactionBody, error)
	ResultsFunc      func(height uint64) ([]*flow.TransactionResult, error)
	EventsFunc       func(height uint64) ([]flow.Event, error)
	SealsFunc        func(height uint64) ([]*flow.Seal, error)
}

func BaselineChain(t *testing.T) *Chain {
	t.Helper()

	c := Chain{
		RootFunc: func() (uint64, error) {
			return GenericHeight, nil
		},
		HeaderFunc: func(height uint64) (*flow.Header, error) {
			return GenericHeader, nil
		},
		CommitFunc: func(height uint64) (flow.StateCommitment, error) {
			return GenericCommit(0), nil
		},
		CollectionsFunc: func(height uint64) ([]*flow.LightCollection, error) {
			return GenericCollections(2), nil
		},
		GuaranteesFunc: func(height uint64) ([]*flow.CollectionGuarantee, error) {
			return GenericGuarantees(2), nil
		},
		TransactionsFunc: func(height uint64) ([]*flow.TransactionBody, error) {
			return GenericTransactions(4), nil
		},
		ResultsFunc: func(height uint64) ([]*flow.TransactionResult, error) {
			return GenericResults(4), nil
		},
		EventsFunc: func(height uint64) ([]flow.Event, error) {
			return GenericEvents(4), nil
		},
		SealsFunc: func(height uint64) ([]*flow.Seal, error) {
			return GenericSeals(4), nil
		},
	}

	return &c
}

func (c *Chain) Root() (uint64, error) {
	return c.RootFunc()
}

func (c *Chain) Header(height uint64) (*flow.Header, error) {
	return c.HeaderFunc(height)
}

func (c *Chain) Commit(height uint64) (flow.StateCommitment, error) {
	return c.CommitFunc(height)
}

func (c *Chain) Collections(height uint64) ([]*flow.LightCollection, error) {
	return c.CollectionsFunc(height)
}

func (c *Chain) Guarantees(height uint64) ([]*flow.CollectionGuarantee, error) {
	return c.GuaranteesFunc(height)
}

func (c *Chain) Transactions(height uint64) ([]*flow.TransactionBody, error) {
	return c.TransactionsFunc(height)
}

func (c *Chain) Results(height uint64) ([]*flow.TransactionResult, error) {
	return c.ResultsFunc(height)
}

func (c *Chain) Events(height uint64) ([]flow.Event, error) {
	return c.EventsFunc(height)
}

func (c *Chain) Seals(height uint64) ([]*flow.Seal, error) {
	return c.SealsFunc(height)
}
