package tilda_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		d       DateTime
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "success",
			args: args{b: []byte(`"2021-09-01 15:04:05"`)},
			want: time.Date(2021, 9, 1, 15, 4, 5, 0, time.UTC),
		}, {
			name: "null",
			args: args{b: []byte(`"null"`)},
			want: time.Time{},
		}, {
			name:    "invalid",
			args:    args{b: []byte(`"2021-09-01 15:04:05:00"`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dt DateTime
			err := dt.UnmarshalJSON(tt.args.b)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, time.Time(dt))
		})
	}
}
