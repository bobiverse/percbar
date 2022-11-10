package percbar

import (
	"strings"
	"testing"
)

type testCase struct {
	header              string
	data                map[string]float64
	expectedSectorCount int
	expectedSum         float64
	expectedHeader      string
}

func TestBar(t *testing.T) {
	testcases := []testCase{

		testCase{
			data: map[string]float64{
				"A": 1,
			},
			expectedSectorCount: 1,
			expectedSum:         1,
			expectedHeader:      "1 in total",
		},

		testCase{
			data: map[string]float64{
				"A": 500,
				"B": 499,
				"C": 1,
			},
			expectedSectorCount: 3,
			expectedSum:         1000,
			expectedHeader:      "1000 in total",
		},

		testCase{
			data: map[string]float64{
				"A": 3451.1,
				"B": 894.2,
				"C": 21.3,
				"D": 1.1,
				"E": 0.02,
			},
			header:              "{SUM} is {SUM}!",
			expectedSectorCount: 5,
			expectedSum:         4367.72,
			expectedHeader:      "4367.72 is 4367.72!",
		},
	}

	for _, tc := range testcases {
		bar := New(tc.data)
		if tc.header != "" {
			bar.options.Header = tc.header
		}
		out := bar.String()

		if int(bar.sum*100) != int(tc.expectedSum*100) {
			t.Fatalf("Expected sum should be %.2f. Found: %.2f", tc.expectedSum, bar.sum)
		}

		if len(bar.sectors) != tc.expectedSectorCount {
			t.Fatalf("Expected sector count should be %d. Found: %d", tc.expectedSectorCount, len(bar.sectors))
		}

		if !strings.Contains(out, tc.expectedHeader) {
			t.Fatalf("Expected header `%s`. Found: `%s`", tc.expectedHeader, strings.Split(out, "\n")[0])
		}
	}

}
