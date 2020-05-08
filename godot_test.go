package godot

import (
	"strings"
	"testing"
)

func TestDot(t *testing.T) {
	tests := []struct {
		name     string
		src      []byte
		wantDist []byte
		wantErr  bool
	}{
		{
			name: "simple test",
			src: []byte(`
			digraph {
				a->b
			}
			`),
			wantDist: []byte(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"
 "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`),
			wantErr: false,
		},
		{
			name: "error",
			src: []byte(`
			digraph {
				a->
			}
			`),
			wantDist: []byte(""),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDist, err := Dot(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dot() error = %s, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasPrefix(string(gotDist), string(tt.wantDist)) {
				t.Errorf("Dot() = %s, want %s", gotDist, tt.wantDist)
			}
		})
	}
}
