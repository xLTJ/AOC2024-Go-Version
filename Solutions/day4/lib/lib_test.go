package lib

import "testing"

func TestCountXmas(t *testing.T) {
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
			name: "Example input",
			args: args{
				fileName: "testInput.txt",
			},
			want:    18,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountXmas(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountXmas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountXmas() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountMasCrosses(t *testing.T) {
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
			name: "Example input",
			args: args{
				fileName: "testInput.txt",
			},
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountMasCrosses(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountMasCrosses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountMasCrosses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountXmas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountXmas("../Input.txt")
		if err != nil {
			return
		}
	}
}

func BenchmarkCountMasCrosses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := CountMasCrosses("../Input.txt")
		if err != nil {
			return
		}
	}
}
