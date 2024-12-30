package lib

import "testing"

func TestCalculateInstructions(t *testing.T) {
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
			want:    161,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateInstructions(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateInstructions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateInstructions2(t *testing.T) {
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
			name:    "Example Input",
			args:    args{fileName: "testInput.txt"},
			want:    48,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateInstructions2(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateInstructions2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateInstructions2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalculateInstructions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateInstructions("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCalculateInstructions2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateInstructions2("../Input.txt")
		if err != nil {
			return
		}
	}
}
