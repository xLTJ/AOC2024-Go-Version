package lib

import "testing"

func TestCountAntinodes(t *testing.T) {
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
			want:    14,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountAntinodes(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountAntinodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountAntinodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountAntinodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountAntinodes("../Input.txt")
		if err != nil {
			return
		}
	}
}
