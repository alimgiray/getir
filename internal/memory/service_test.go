package memory_test

import (
	"testing"

	"github.com/alimgiray/getir/internal/memory"

	"github.com/stretchr/testify/assert"
)

func TestFindNotExistingRecord(t *testing.T) {

	service := memory.NewMemoryService()
	_, err := service.FindRecord("key_does_not_exists")

	assert.NotEqual(t, nil, err, "It should return an error because record because given key does not exists")
}

func TestExistingRecord(t *testing.T) {

	service := memory.NewMemoryService()
	err := service.CreateNewRecord(&memory.Record{Key: "key", Value: "value"})

	assert.Equal(t, nil, err, "It should successfully create a record with given correct values")

	record, err := service.FindRecord("key")

	assert.Equal(t, nil, err, "It should not return an error because record because given key exists")
	assert.Equal(t, "value", record.Value, "It should successfully return correct value for given key")
}
