package ui

import "testing"

func TestSelectOptions(t *testing.T) {
	type args struct {
		options []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{name: "Test case 1",
			args: args{
				options: []string{"Option 1", "Option 2", "Option 3"},
			},
			want:    "Option 1",
			wantErr: false,
		},
		{name: "Test case 2",
			args: args{
				options: []string{"Option 1", "Option 2", "Option 3"},
			},
			want:    "Option 2",
			wantErr: false,
		},
		{name: "Test case 3",
			args: args{
				options: []string{"Option 1", "Option 2", "Option 3"},
			},
			want:    "Option 3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := SelectOptions(tt.args.options)

			if (err != nil) != tt.wantErr {
				t.Errorf("SelectOptions() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("SelectOptions() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
