package sensor

import (
	"testing"
	"time"
)

func TestLookup(t *testing.T) {
	type args struct {
		addr    string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		count   int
		want1   time.Duration
		wantErr bool
	}{
		{
			name:    "google.com",
			args:    args{"google.com", 1000 * time.Millisecond},
			count:   1,
			want1:   200 * time.Millisecond,
			wantErr: false,
		},
		{
			name:    "google187.com",
			args:    args{"google187.com", 1000 * time.Millisecond},
			count:   0,
			want1:   200 * time.Millisecond,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Lookup(tt.args.addr, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.count {
				t.Errorf("Lookup() addresses returned got %v, wanted <= %v", len(got), tt.count)
			}
			if got1 >= tt.want1 {
				t.Errorf("Lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPrtgLookup(t *testing.T) {
	type args struct {
		a       []string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"pass", args{[]string{"www.google.com","www.facebook.com"}, time.Second}, false},
		{"fail", args{[]string{"www.google.com","www.facebook.com","www.abcdbcda.com"}, time.Second}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrtgLookup(tt.args.a, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("PrtgLookup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
