package percbar

import "testing"

// sectors must sort descending by count (Less: sect[i].count > sect[j].count).
func TestSectorsSortDescending(t *testing.T) {
	tests := []struct {
		name   string
		counts []float64
		want   []float64
	}{
		{"already descending", []float64{300, 200, 100}, []float64{300, 200, 100}},
		{"ascending gets reversed", []float64{1, 2, 3}, []float64{3, 2, 1}},
		{"unordered", []float64{5, 50, 1, 25}, []float64{50, 25, 5, 1}},
		{"ties keep relative stability not required, only order", []float64{10, 10, 5}, []float64{10, 10, 5}},
		{"negative and positive", []float64{-5, 10, -1}, []float64{10, -1, -5}},
		{"single sector", []float64{7}, []float64{7}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			values := map[string]float64{}
			for i, c := range tc.counts {
				values[string(rune('a'+i))] = c
			}
			bar := New(values)

			if len(bar.sectors) != len(tc.want) {
				t.Fatalf("expected %d sectors, got %d", len(tc.want), len(bar.sectors))
			}

			for i, s := range bar.sectors {
				if s.count != tc.want[i] {
					t.Errorf("position %d: expected count %v, got %v", i, tc.want[i], s.count)
				}
			}
		})
	}
}
