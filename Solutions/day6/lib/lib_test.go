package lib

import "testing"

func TestCountCellsPassed(t *testing.T) {
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
			want:    41,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountCellsPassed(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountCellsPassed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountCellsPassed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLoopCreatingObstacles(t *testing.T) {
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
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLoopCreatingObstacles(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountLoopCreatingObstacles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountLoopCreatingObstacles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountCellsPassed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountCellsPassed("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCountLoopCreatingObstacles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountLoopCreatingObstacles("../Input.txt")
		if err != nil {
			return
		}
	}
}
