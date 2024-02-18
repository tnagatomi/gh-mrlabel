package parser

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tnagatomi/gh-mrlabel/option"
	"testing"
)

func TestLabel(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []option.Label
		wantErr bool
	}{
		{
			name: "Single label",
			args: args{
				input: "bug:ff0000:This is a bug",
			},
			want: []option.Label{
				{
					Name:        "bug",
					Color:       "ff0000",
					Description: "This is a bug",
				},
			},
			wantErr: false,
		},
		{
			name: "Single label without description",
			args: args{
				input: "bug:ff0000",
			},
			want: []option.Label{
				{
					Name:        "bug",
					Color:       "ff0000",
					Description: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Multi labels",
			args: args{
				input: "bug:ff0000:This is a bug,enhancement:00ff00:This is an enhancement,question:0000ff:This is a question",
			},
			want: []option.Label{
				{
					Name:        "bug",
					Color:       "ff0000",
					Description: "This is a bug",
				},
				{
					Name:        "enhancement",
					Color:       "00ff00",
					Description: "This is an enhancement",
				},
				{
					Name:        "question",
					Color:       "0000ff",
					Description: "This is a question",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid label format",
			args: args{
				input: "bug:ff0000:This is a bug:invalid,enhancement:00ff00:This is an enhancement",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid color format",
			args: args{
				input: "bug:ff000:This is a bug,enhancement:00ff00:This is an enhancement",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Label(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Label() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Label() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
