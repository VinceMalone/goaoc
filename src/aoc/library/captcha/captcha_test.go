package captcha

import "testing"

var cases = []struct {
	in   string
	fn   Lookahead
	want int
}{
	{"1122", WrapAroundAdd, 3},
	{"1111", WrapAroundAdd, 4},
	{"1234", WrapAroundAdd, 0},
	{"91212129", WrapAroundAdd, 9},
	{"1212", HalfwayAroundAdd, 6},
	{"1221", HalfwayAroundAdd, 0},
	{"123425", HalfwayAroundAdd, 4},
	{"123123", HalfwayAroundAdd, 12},
	{"12131415", HalfwayAroundAdd, 4},
}

func TestInverseCaptcha(t *testing.T) {
	for _, c := range cases {
		if got := InverseCaptcha(c.in, c.fn); got != c.want {
			t.Errorf("test failed for %v; wanted:%v but got:%v", c.in, c.want, got)
		}
	}
}

func BenchmarkInverseCaptchaWithWrapAround(b *testing.B) {
	b.StopTimer()
	for _, c := range cases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			InverseCaptcha(c.in, WrapAroundAdd)
		}
		b.StopTimer()
	}
}

func BenchmarkInverseCaptchaWithHalfwayAround(b *testing.B) {
	b.StopTimer()
	for _, c := range cases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			InverseCaptcha(c.in, HalfwayAroundAdd)
		}
		b.StopTimer()
	}
}
