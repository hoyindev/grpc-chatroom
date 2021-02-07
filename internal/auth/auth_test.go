package auth

import (
	"fmt"
	"testing"
)

func TestMakeClientHashedPW(t *testing.T) {
	tests := []struct {
		a    string
		want string
	}{
		{
			"password",
			"lgviX0dEsDNo5bi0aF93Dhey7LGPZCcHZwWODL8saqo=",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.a)
		t.Run(testname, func(t *testing.T) {

			clientH := MakeClientHashedPW(tt.a)
			ans := MakeServerHashedPW(clientH)

			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}

}
