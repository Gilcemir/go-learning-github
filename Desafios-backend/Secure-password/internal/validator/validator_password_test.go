package validator

import (
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectedValid bool
		expectedError string
	}{
		{
			name:          "Valid password",
			password:      "Abcdef1@",
			expectedValid: true,
			expectedError: "",
		},
		{
			name:          "Password too short",
			password:      "Ab1@",
			expectedValid: false,
			expectedError: MinLenErr + "\n",
		},
		{
			name:          "Missing uppercase",
			password:      "abcdef1@",
			expectedValid: false,
			expectedError: MaiusErr + "\n",
		},
		{
			name:          "Missing lowercase",
			password:      "ABCDEF1@",
			expectedValid: false,
			expectedError: MinsErr + "\n",
		},
		{
			name:          "Missing number",
			password:      "Abcdefgh@",
			expectedValid: false,
			expectedError: NumErr + "\n",
		},
		{
			name:          "Missing special character",
			password:      "Abcdef12",
			expectedValid: false,
			expectedError: SpecErr + "\n",
		},
		{
			name:          "Missing lowercase and number",
			password:      "ABCDEFG@",
			expectedValid: false,
			expectedError: MinsErr + "\n" + NumErr + "\n",
		},
		{
			name:          "Too short and missing special character",
			password:      "Ab1",
			expectedValid: false,
			expectedError: MinLenErr + "\n" + SpecErr + "\n",
		},
		{
			name:          "Missing uppercase, lowercase, and special character",
			password:      "12345678",
			expectedValid: false,
			expectedError: MaiusErr + "\n" + MinsErr + "\n" + SpecErr + "\n",
		},
		{
			name:          "Too short, missing number and special character",
			password:      "Abc",
			expectedValid: false,
			expectedError: MinLenErr + "\n" + NumErr + "\n" + SpecErr + "\n",
		},
		{
			name:          "Missing everything",
			password:      "",
			expectedValid: false,
			expectedError: MinLenErr + "\n" + MaiusErr + "\n" + MinsErr + "\n" + NumErr + "\n" + SpecErr + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := IsValidPassword(&tt.password)

			if valid != tt.expectedValid {
				t.Errorf("expected validity: %v, got: %v", tt.expectedValid, valid)
			}

			if err == nil && tt.expectedError != "" {
				t.Errorf("expected error: %v, got: nil", tt.expectedError)
			} else if err != nil && tt.expectedError == "" {
				t.Errorf("expected error: nil, got: %v", err.Error())
			} else if err != nil && tt.expectedError != "" {
				if err.Error() != tt.expectedError {
					t.Errorf("expected error message: %v, got: %v", tt.expectedError, err.Error())
				}
			}
		})
	}
}
