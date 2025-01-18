package lib

import "testing"

func TestGetMiddleSum(t *testing.T) {
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
			want:    143,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMiddleSum(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMiddleSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMiddleSum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortAndGetMiddleSum(t *testing.T) {
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
			want:    123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortAndGetMiddleSum(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMiddleSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMiddleSum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetMiddleSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetMiddleSum("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkSortAndGetMiddleSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := SortAndGetMiddleSum("../Input.txt")
		if err != nil {
			return
		}
	}
}
