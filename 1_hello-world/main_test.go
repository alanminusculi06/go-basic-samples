package main

import (
	"errors"
	"testing"
)

//func Test_getGreeting(t *testing.T) {
//	// Arrange
//	name := "Alan"
//
//	// Act
//	message := getGreeting(name)
//
//	// Assert
//	if !reflect.DeepEqual("Hello Alan!", message) {
//		t.Fail()
//	}
//}

//Table Driven Tests

func Test_getGreeting1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name          string
		args          args
		expected      string
		expectedError error
	}{
		{
			name: "Test 1",
			args: args{
				name: "Alan",
			},
			expected: "Hello, Alan!",
		},
		{
			name: "Test 2",
			args: args{
				name: "John",
			},
			expected: "Hello, John!",
		},
		{
			name: "Test 3",
			args: args{
				name: "",
			},
			expectedError: errors.New("name is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotError := getGreeting(tt.args.name)
			if got != tt.expected {
				t.Errorf("getGreeting() = %v, expected %v", got, tt.expected)
			}
			if gotError != tt.expectedError {
				t.Errorf("getGreeting() = %v, expected %v", got, tt.expectedError)
			}
		})
	}
}
