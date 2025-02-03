package lib

import "testing"

func TestDay9Part1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Example input",
			args:    args{fileName: "testInput.txt"},
			want:    1928,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Day9Part1(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day9Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day9Part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDay9Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Day9Part1("../Input.txt")
		if err != nil {
			return
		}
	}
}
