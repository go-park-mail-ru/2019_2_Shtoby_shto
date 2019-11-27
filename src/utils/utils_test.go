package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUUID(t *testing.T) {

	uuid, err := GenerateUUID()
	assert.Nil(t, err, "can not generate uuid")

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "uuju",
			args: args{},
			want: false,
		},
		{
			name: "1312",
			args: args{},
			want: false,
		},
		{
			name: uuid.String(),
			args: args{uuid.String()},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUUID(tt.args.str); got != tt.want {
				t.Errorf("IsUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUUID(t *testing.T) {

}

func TestJoin(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.args...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
