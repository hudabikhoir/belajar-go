package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

//menjalankan testing -> go test kubus.go kubus_test.go -v
func TestHitungVolume(t *testing.T) {
	// contoh testing dengan menggunakan testify
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan salah")
}
func TestHitungLuas(t *testing.T) {
	// contoh testing manual
	t.Logf("Volume : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH! seharusnya %.2f", luasSeharusnya)
	}
}
func TestHitungKeliling(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! seharusnya %.2f", kelilingSeharusnya)
	}
}

// menjalankan testing benchmark -> go test kubus.go kubus_test.go -bench=.
func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
