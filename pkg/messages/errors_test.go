package messages

import (
	"testing"
)

func TestPlanetError_Error(t *testing.T) {
	tests := []struct {
		name string
		p    *PlanetError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Error(); got != tt.want {
				t.Errorf("PlanetError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilmError_Error(t *testing.T) {
	tests := []struct {
		name string
		f    *FilmError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Error(); got != tt.want {
				t.Errorf("FilmError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
