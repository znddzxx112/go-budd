package server

import (
	"testing"
)

func TestPong(t *testing.T) {

	pongLocalBase := localBase.Post("health")

	tests := []struct {
		name string
		args PongArgs
		want PongVOResult
	}{
		{
			name: "case100100",
			args: PongArgs{},
			want: PongVOResult{
				Code:   1,
				Result: PongVO{},
			},
		},
		{
			name: "case100101",
			args: PongArgs{
				Say: "1",
			},
			want: PongVOResult{
				Code:   4,
				Result: PongVO{},
			},
		},
		{
			name: "case100102",
			args: PongArgs{
				Say: "hello",
			},
			want: PongVOResult{
				Code:   0,
				Result: PongVO{Res: "hello"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &PongVOResult{}
			_, err := pongLocalBase.New().BodyJSON(test.args).ReceiveSuccess(res)
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
