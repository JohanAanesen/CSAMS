package util_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input string
		hash  string
	}{
		{
			input: "abc",
			hash:  "$2y$14$8z8cxB6nNs2rKF7bzCSFUupnjQlWeVxvWVjxCpnhorIXqSb7/lqk2",
		},
		{
			input: "123",
			hash:  "$2y$14$06c1yJiqV4TUL.DaR4inUuGq8pI2LxGuRTL4hfGwZalJz1gTbNuSK",
		},
		{
			input: "password",
			hash:  "$2y$14$ojBn3OiufwDYwwMyqf3vreGOB2ed3/RHG5.GlHnE/7KNRwGerYUtW",
		},
	}

	for _, test := range tests {
		_, err := util.GenerateFromPassword(test.input)
		if err != nil {
			t.Fail()
		}

		err = util.CompareHashAndPassword(test.input, test.hash)
		if err != nil {
			t.Fail()
		}
	}
}
