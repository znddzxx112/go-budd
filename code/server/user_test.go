package server

import "testing"

func TestDefaultServer_UserLogin(t *testing.T) {
	userLoginlocalBase := localBase.Post("/user/login")

	tests := []struct {
		name string
		args interface{}
		want UserLoginVOResult
	}{
		{
			name: "case101100",
			args: UserLoginArgs{
				Mobile: "",
			},
			want: UserLoginVOResult{
				Code:   1,
				Result: UserLoginVO{},
			},
		},
		{
			name: "case101101",
			args: UserLoginArgs{
				Mobile: "18800011",
			},
			want: UserLoginVOResult{
				Code:   4,
				Result: UserLoginVO{},
			},
		},
		{
			name: "case101102",
			args: UserLoginArgs{
				Mobile: "18800011121",
			},
			want: UserLoginVOResult{
				Code:   1,
				Result: UserLoginVO{},
			},
		},
		{
			name: "case101103",
			args: UserLoginArgs{
				Mobile: "18800011122",
			},
			want: UserLoginVOResult{
				Code:   0,
				Result: UserLoginVO{Token: "10001"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &UserLoginVOResult{}
			_, err := userLoginlocalBase.New().BodyJSON(test.args).ReceiveSuccess(res)
			if err != nil {
				t.Error(err)
			}
			if res.Code != test.want.Code {
				t.Error(test.name, res.Code, test.want.Code)
			}
			if res.Result != test.want.Result {
				t.Error(test.name, res.Result, test.want.Result)
			}
		})
	}
}
