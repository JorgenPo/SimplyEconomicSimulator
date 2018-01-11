package sound

import (
	"github.com/faiface/beep/mp3"
	"os"
	"github.com/faiface/beep/speaker"
	"time"
	"github.com/faiface/beep"
	"log"
)

const (
	MainThemeSound string = "sound/main_theme.mp3"
)

func PlayMainTheme(loop bool) {
	f, err := os.Open(MainThemeSound)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	s, format, err := mp3.Decode(f)
	if err != nil {
		// TODO: error logging
		return
	}

	// Playing audio at sample rate of file and set
	// buffer size to 1/10 of second to stable output
	speaker.Init(format.SampleRate * 2, format.SampleRate.N(time.Second / 10))

	for {
		s.Seek(0)

		done := make(chan struct{})

		speaker.Play(beep.Seq(s, beep.Callback(func() {
			// End of audio stream!
			close(done)
		})))

		// Wait for end of stream
		<-done

		if !loop {
			return
		}
	}
}
