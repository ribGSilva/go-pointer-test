package pointer

import (
	"runtime"
	"testing"
)

const max = 1_000_000
const gc = true

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			Run()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunRef()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunP()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunNew()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunN()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunNRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunNRef()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunNP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunNP()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunNNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunNNew()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunM()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunMRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunMRef()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunMP()
		}
		if gc {
			runtime.GC()
		}
	}
}

func BenchmarkRunMNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for o := 0; o < max; o++ {
			RunMNew()
		}
		if gc {
			runtime.GC()
		}
	}
}
