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

	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"
)

type VirtualMachine struct {
	GetAccountFunc func(ctx fvm.Context, address flow.Address, v state.View, programs *programs.Programs) (*flow.Account, error)
	RunFunc        func(ctx fvm.Context, proc fvm.Procedure, v state.View, programs *programs.Programs) error
}

func BaselineVirtualMachine(t *testing.T) *VirtualMachine {
	t.Helper()

	vm := VirtualMachine{
		GetAccountFunc: func(ctx fvm.Context, address flow.Address, v state.View, programs *programs.Programs) (*flow.Account, error) {
			return &GenericAccount, nil
		},
		RunFunc: func(ctx fvm.Context, proc fvm.Procedure, v state.View, programs *programs.Programs) error {
			return nil
		},
	}

	return &vm
}

func (v *VirtualMachine) GetAccount(ctx fvm.Context, address flow.Address, view state.View, programs *programs.Programs) (*flow.Account, error) {
	return v.GetAccountFunc(ctx, address, view, programs)
}

func (v *VirtualMachine) Run(ctx fvm.Context, proc fvm.Procedure, view state.View, programs *programs.Programs) error {
	return v.RunFunc(ctx, proc, view, programs)
}
