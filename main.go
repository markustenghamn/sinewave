package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"math"
	"time"
)

func sineWave(sr beep.SampleRate, freq float64) beep.Streamer {
	t := 0.0
	sineFn := sine(1, freq, 0)
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		sampleLength := len(samples)
		for i := range samples {
			y := sineFn(t)
			samples[i][0] = y
			samples[i][1] = y
			t += sr.D(1).Seconds()
		}
		return sampleLength, true
	})
}

func sine(amplitude float64, frequency float64, phaseShift float64) func(t float64) float64 {
	return func(t float64) float64 {
		return amplitude * math.Sin(
			2*math.Pi*frequency*t+phaseShift,
		)
	}
}

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(sineWave(sr, 450))
	select {}
}
