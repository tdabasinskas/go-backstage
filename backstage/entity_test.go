package backstage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestListEntityFilterString tests if list entity filter string is correctly generated.
func TestListEntityFilterString(t *testing.T) {
	tests := []struct {
		name     string
		filter   ListEntityFilter
		expected string
	}{
		{
			name:     "empty filter",
			filter:   ListEntityFilter{},
			expected: "",
		},
		{
			name: "filter with one key-value pair",
			filter: ListEntityFilter{
				"key1": "value1",
			},
			expected: "key1=value1",
		},
		{
			name: "filter with one key and no value",
			filter: ListEntityFilter{
				"key1": "",
			},
			expected: "key1",
		},
		{
			name: "filter with multiple key-value pairs",
			filter: ListEntityFilter{
				"key1": "value1",
				"key2": "value2",
				"key3": "",
			},
			expected: "key1=value1,key2=value2,key3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.filter.string()
			assert.Equal(t, test.expected, actual, "List entity filter string should match expected value")
		})
	}
}

// TestListEntityOrderString tests if list entity order string is correctly generated.
func TestListEntityOrderString(t *testing.T) {
	tests := []struct {
		name      string
		order     ListEntityOrder
		expected  string
		shouldErr bool
	}{
		{
			name: "valid order",
			order: ListEntityOrder{
				Direction: OrderAscending,
				Field:     "field1",
			},
			expected:  "asc:field1",
			shouldErr: false,
		},
		{
			name: "valid descending order",
			order: ListEntityOrder{
				Direction: OrderDescending,
				Field:     "field2",
			},
			expected:  "desc:field2",
			shouldErr: false,
		},
		{
			name: "invalid order direction",
			order: ListEntityOrder{
				Direction: "invalid",
				Field:     "field3",
			},
			expected:  "",
			shouldErr: true,
		},
		{
			name:      "empty order",
			order:     ListEntityOrder{},
			expected:  "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.order.string()
			if test.shouldErr {
				assert.Error(t, err, "Expected error but got nil")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.Equal(t, test.expected, actual, "List entity order string should match expected value")
			}
		})
	}
}
