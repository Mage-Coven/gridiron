package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridironContractDetailsValidation(t *testing.T) {
	specs := map[string]struct {
		src    GridironContractDetails
		expErr bool
	}{
		"all good": {
			src: GridironContractDetailsFixture(t),
		},
		"empty callbacks": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = nil
			}),
		},
		"multiple callbacks": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{
					{Position: 1, PrivilegeType: "begin_blocker"},
					{Position: 1, PrivilegeType: "end_blocker"},
				}
			}),
		},
		"duplicate callbacks": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{
					{Position: 1, PrivilegeType: "begin_blocker"},
					{Position: 2, PrivilegeType: "begin_blocker"},
				}
			}),
			expErr: true,
		},
		"unknown callback": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: 1, PrivilegeType: "unknown"}}
			}),
			expErr: true,
		},
		"empty callback": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: 1}}
			}),
			expErr: true,
		},
		"invalid callback position": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{Position: math.MaxUint8 + 1, PrivilegeType: "begin_blocker"}}
			}),
			expErr: true,
		},
		"empty callback position": {
			src: GridironContractDetailsFixture(t, func(d *GridironContractDetails) {
				d.RegisteredPrivileges = []RegisteredPrivilege{{PrivilegeType: "begin_blocker"}}
			}),
			expErr: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			gotErr := spec.src.ValidateBasic()
			if spec.expErr {
				assert.Error(t, gotErr)
				return
			}
			assert.NoError(t, gotErr)
		})
	}
}
