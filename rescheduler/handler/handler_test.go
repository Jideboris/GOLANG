package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"umotif.com/rescheduler/handler/mocks/repomocks"
)

func TestHander_When_GenerateTable_ReturnsNoError(t *testing.T) {
	t.Parallel()
	repo := &repomocks.ComputeHandler{}
	repo.On("GenerateTable", mock.Anything).
		Return(true, nil).
		Once()
		
	done, err := repo.GenerateTable()

	assert.Nil(t, err)
	assert.Equal(t, true, done)
}
