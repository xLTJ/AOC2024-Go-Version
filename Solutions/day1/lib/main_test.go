package lib

import "testing"

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    int
		wantErr bool
	}{
		{
			name: "Example Input",
			arg:  "testInput.txt",
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateDistance(tt.arg)
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
