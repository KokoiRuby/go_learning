package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// a mock obj that mock Printer interface
type PrinterMock struct {
	mock.Mock
}

func (mk *PrinterMock) Print(paperSize string, content string) error {
	// stub it out instead of doing something useful
	args := mk.Called(paperSize, content)
	fmt.Println("Printing from mock objct")
	return args.Error(0)
}

func TestPrintA4(t *testing.T) {
	// printer mock
	printer := &PrinterMock{}
	// set expectation that the mocked function will receive
	printer.On("Print", "A4", "Hello World").Return(nil)

	computer := Computer{Pr: printer}
	computer.PrintA4("Hello World")

	printer.AssertExpectations(t)

}
