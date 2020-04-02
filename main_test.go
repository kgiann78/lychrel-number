package lychrel_number

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_facts(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		converges int
	}{
		{"converges", 1, 0},
		{"converges", 2, 0},
		{"converges", 9, 0},
		{"converges", 10, 1},
		{"converges", 11, 0},
		{"converges", 56, 1},
		{"converges", 57, 2},
		{"converges", 59, 3},
		{"converges", 89, 24},
		{"converges", 187, 23},
		{"converges", 1297, 21},
		{"converges", 10911, 55},
		{"converges", 1005499526, 109},
		{"converges", 1186060307891929990, 261},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := converges(tt.n)
			if ans != tt.converges {
				t.Errorf("got %d, wanted %d", ans, tt.converges)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"isPalindrome", args{n: big.NewInt(int64(1111))}, true},
		{"isPalindrome", args{n: big.NewInt(int64(2212))}, false},
		{"isPalindrome", args{n: big.NewInt(int64(1000))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 bool
	}{
		{"reverse", args{n: big.NewInt(int64(1234))}, big.NewInt(int64(4321)), true},
		{"reverse", args{n: big.NewInt(int64(12345677890))}, big.NewInt(int64(9877654321)), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := reverse(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("reverse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
