package lib

import "testing"

type test struct {
	name     string
	fileName string
	want     int
	wantErr  bool
}

func TestCalculateDistance(t *testing.T) {
	tests := []test{
		{
			name:     "Example Input",
			fileName: "testInput.txt",
			want:     11,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateDistance(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSimilarity(t *testing.T) {
	tests := []test{
		{
			name:     "Example Input",
			fileName: "testInput.txt",
			want:     31,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateSimilarity(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSimilarity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateSimilarity() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalculateDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateDistance("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCalculateSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CalculateSimilarity("../Input.txt")
		if err != nil {
			return
		}
	}
}
