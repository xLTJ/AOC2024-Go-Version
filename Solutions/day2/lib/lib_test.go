package lib

import (
	"testing"
)

func TestCountSafeReports(t *testing.T) {
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
			name: "Example",
			args: args{
				fileName: "testInput.txt",
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountSafeReports(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountSafeReports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountSafeReports() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSafeReports2(t *testing.T) {
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
			name: "Example",
			args: args{
				fileName: "testInput.txt",
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountSafeReports2(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountSafeReports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountSafeReports() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountSafeReports(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountSafeReports("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCountSafeReports2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountSafeReports2("../Input.txt")
		if err != nil {
			return
		}
	}
}
