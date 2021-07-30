package database

import (
	"fmt"
	"testing"
)

func TestGetConnection(t *testing.T) {
	_ = GetConnection()
	fmt.Println("Success connect to database")
}

