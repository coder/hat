package hat

import "testing"

func Test_urlJoin(t *testing.T) {
	type args struct {
		elems []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no input slashes", args{[]string{"http://google.com", "maps", "austin"}}, "http://google.com/maps/austin"},
		{"input slashes", args{[]string{"http://google.com/", "/maps/", "/austin/"}}, "http://google.com/maps/austin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlJoin(tt.args.elems...); got != tt.want {
				t.Errorf("urlJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
