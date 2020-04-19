package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)


func Test_Main(t *testing.T) {
	result := soma(5, 5)
	require.Equal(t, 10, result)
}
