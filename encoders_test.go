package uuid

import "testing"

type encTC struct {
	src  string
	b64u string
	b64s string
}

var (
	// for the purpose of these tests UUIDs don't have to be v1
	encTCs = []encTC{
		{"f254df4a-184c-1019-80a4-c61cd00a6899", "8lTfShhMEBmApMYc0ApomQ", "8lTfShhMEBmApMYc0ApomQ"},
		{"86ef2c67-ccae-4241-8543-622e8589c62a", "hu8sZ8yuQkGFQ2IuhYnGKg", "hu8sZ8yuQkGFQ2IuhYnGKg"},
		{"04e37eeb-6881-45db-976b-ec2efbb0e475", "BON-62iBRduXa-wu-7DkdQ", "BON+62iBRduXa+wu+7DkdQ"},
		{"db1aa9d6-9497-485d-a9aa-be6609e270a7", "2xqp1pSXSF2pqr5mCeJwpw", "2xqp1pSXSF2pqr5mCeJwpw"},
		{"63ccdba7-b775-4348-b6d1-1694fae1a729", "Y8zbp7d1Q0i20RaU-uGnKQ", "Y8zbp7d1Q0i20RaU+uGnKQ"},
		{"43c590f3-a400-4a7e-84cf-fe64a99841ed", "Q8WQ86QASn6Ez_5kqZhB7Q", "Q8WQ86QASn6Ez/5kqZhB7Q"},
		{"4b9ab787-b0d0-47c4-9971-6dfc5a6d8db3", "S5q3h7DQR8SZcW38Wm2Nsw", "S5q3h7DQR8SZcW38Wm2Nsw"},
	}
)

func TestEncoders(t *testing.T) {
	for i, tc := range encTCs {
		uuid, err := NewFromString(tc.src)
		if err != nil {
			t.Errorf("TestEncoders[%d]: %s", i, err)
		}

		act := string(uuid.Encode(Base64URLEncoder))
		if act != tc.b64u {
			t.Errorf("TestEncoders[%d]: Base64URLEncoder got %s want %s", i, act, tc.b64u)
		}

		act = uuid.EncodeToString(Base64StdEncoder)
		if act != tc.b64s {
			t.Errorf("TestEncoders[%d]: Base64StdEncoder got %s want %s", i, act, tc.b64s)
		}
	}
}
