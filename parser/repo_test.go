package parser

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tnagatomi/gh-mrlabel/option"
	"testing"
)

func TestRepo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []option.Repo
		wantErr bool
	}{
		{
			name: "Single repo",
			args: args{
				input: "tnagatomi/repo1",
			},
			want: []option.Repo{
				{
					Owner: "tnagatomi",
					Repo:  "repo1",
				},
			},
			wantErr: false,
		},
		{
			name: "Multi repos",
			args: args{
				input: "tnagatomi/repo1,tnagatomi/repo2",
			},
			want: []option.Repo{
				{
					Owner: "tnagatomi",
					Repo:  "repo1",
				},
				{
					Owner: "tnagatomi",
					Repo:  "repo2",
				},
			},
		},
		{
			name: "Invalid format",
			args: args{
				input: "tnagatomi/repo1,repo2",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Repo(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Label() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
