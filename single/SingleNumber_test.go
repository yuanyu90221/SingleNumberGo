package single

import "testing"

func TestSingleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,2,1]",
			args: args{nums: []int{2, 2, 1}},
			want: 1,
		},
		{
			name: "[4,1,2,1,2]",
			args: args{nums: []int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "[1]",
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.args.nums); got != tt.want {
				t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
