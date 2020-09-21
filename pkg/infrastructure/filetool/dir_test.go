package filetool

import (
	"os"
	"testing"
)

func TestDirIsExist(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "tmp",
			args: args{
				dir: os.TempDir(),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirIsExist(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirIsExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirIsExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}
