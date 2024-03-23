package config

import (
	"errors"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	type args struct {
		configPath string
		config     any
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "empty config path",
			args: args{
				configPath: "",
				config:     struct{}{},
			},
			wantErr: ErrConfigPathIsEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadFromFile(tt.args.configPath, tt.args.config); !errors.Is(err, tt.wantErr) {
				t.Errorf("LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
