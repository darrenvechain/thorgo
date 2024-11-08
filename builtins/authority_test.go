package builtins

import (
	"testing"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/stretchr/testify/assert"
)

func TestAuthority(t *testing.T) {
	authority, err := NewAuthority(thor)
	assert.NoError(t, err)

	candidates, err := authority.FilterCandidate([]AuthorityCandidateCriteria{}, &thorest.LogFilters{})
	assert.NoError(t, err)
	assert.NotEmpty(t, candidates)

	res, err := authority.Get(candidates[0].NodeMaster)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
