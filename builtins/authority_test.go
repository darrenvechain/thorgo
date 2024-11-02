package builtins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthority_FilterCandidate(t *testing.T) {
	authority, err := NewAuthority(thor)
	assert.NoError(t, err)

	candidates, err := authority.FilterCandidate(make([]AuthorityCandidateCriteria, 0), nil, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, candidates)

	res, err := authority.Get(candidates[0].NodeMaster)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
