package entity_test

import (
	"testing"

	"github.com/rknruben56/feedback-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTemplate(t *testing.T) {
	temp, err := entity.NewTemplate("Class123", "[Student] is doing well")
	assert.Nil(t, err)
	assert.Equal(t, temp.Class, "Class123")
	assert.NotNil(t, temp.ID)
}

func TestTemplateValidate(t *testing.T) {
	type test struct {
		class    string
		content  string
		expected error
	}

	tests := []test{
		{
			class:    "Class123",
			content:  "[Student] is doing well",
			expected: nil,
		},
		{
			class:    "",
			content:  "[Student] is doing well",
			expected: entity.ErrInvalidEntity,
		},
		{
			class:    "Class123",
			content:  "",
			expected: entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewTemplate(tc.class, tc.content)
		assert.Equal(t, err, tc.expected)
	}
}
