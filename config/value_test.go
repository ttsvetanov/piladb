package config

import (
	"testing"
	"time"

	"github.com/fern4lvarez/piladb/config/vars"
)

func TestMaxStackSize(t *testing.T) {
	c := NewConfig()

	inputOutput := []struct {
		input  interface{}
		output int
	}{
		{8, 8},
		{23.7, 23},
		{"3", 3},
		{-1, vars.MaxStackSizeDefault},
		{"foo", vars.MaxStackSizeDefault},
		{-35, vars.MaxStackSizeDefault},
		{[]byte("foo"), vars.MaxStackSizeDefault},
	}

	for _, io := range inputOutput {
		c.Set(vars.MaxStackSize, io.input)

		if s := c.MaxStackSize(); s != io.output {
			t.Errorf("MaxStackSize is %d, expected %d", s, io.output)
		}
	}
}

func TestReadTimeout(t *testing.T) {
	c := NewConfig()

	inputOutput := []struct {
		input  interface{}
		output time.Duration
	}{
		{8, 8},
		{23.7, 23},
		{"3", 3},
		{-1, vars.ReadTimeoutDefault},
		{"foo", vars.ReadTimeoutDefault},
		{[]byte("foo"), vars.ReadTimeoutDefault},
	}

	for _, io := range inputOutput {
		c.Set(vars.ReadTimeout, io.input)

		if s := c.ReadTimeout(); s != io.output {
			t.Errorf("ReadTimeout is %d, expected %d", s, io.output)
		}
	}
}

func TestWriteTimeout(t *testing.T) {
	c := NewConfig()

	inputOutput := []struct {
		input  interface{}
		output time.Duration
	}{
		{8, 8},
		{23.7, 23},
		{"3", 3},
		{-1, vars.WriteTimeoutDefault},
		{"foo", vars.WriteTimeoutDefault},
		{[]byte("foo"), vars.WriteTimeoutDefault},
	}

	for _, io := range inputOutput {
		c.Set(vars.WriteTimeout, io.input)

		if s := c.WriteTimeout(); s != io.output {
			t.Errorf("WriteTimeout is %d, expected %d", s, io.output)
		}
	}
}

func TestPort(t *testing.T) {
	c := NewConfig()

	inputOutput := []struct {
		input  interface{}
		output int
	}{
		{8090, 8090},
		{23676.7, 23676},
		{"3756", 3756},
		{-1, vars.PortDefault},
		{"foo", vars.PortDefault},
		{[]byte("foo"), vars.PortDefault},
		{6736373635, vars.PortDefault},
	}

	for _, io := range inputOutput {
		c.Set(vars.Port, io.input)

		if s := c.Port(); s != io.output {
			t.Errorf("Port is %d, expected %d", s, io.output)
		}
	}
}
