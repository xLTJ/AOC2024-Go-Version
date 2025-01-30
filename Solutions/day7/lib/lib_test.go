package lib

import "testing"

func TestCalculateTotalCalibration(t *testing.T) {
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
			want:    3749,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTotalCalibration(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTotalCalibration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTotalCalibration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateTotalCalibration2(t *testing.T) {
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
			want:    11387,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTotalCalibration2(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTotalCalibration2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateTotalCalibration2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalculateTotalCalibration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateTotalCalibration("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCalculateTotalCalibration2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateTotalCalibration2("../Input.txt")
		if err != nil {
			return
		}
	}
}
