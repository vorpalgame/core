package audio

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func NewAudioPlayer() AudioPlayer {
	return &audioPlayer{ctrl: &beep.Ctrl{}}
}

const (
	once = 1
	loop = -1
)

type AudioPlayer interface {
	PlayOnce()
	PlayLooped()
	Stop()
	Pause()
	Unload()
	Load(fileName string) AudioPlayer
	IsStopped() bool
	IsPlaying() bool
}

type audioPlayer struct {
	stream beep.StreamSeeker
	format beep.Format
	file   *os.File
	ctrl   *beep.Ctrl
}

func (s *audioPlayer) Load(fileName string) AudioPlayer {
	var err error
	s.file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)

	}
	s.stream, s.format, err = mp3.Decode(s.file)

	return s
}

func (s *audioPlayer) PlayOnce() {
	s.play(once)

}

func (s *audioPlayer) PlayLooped() {
	s.play(loop)
}

func (s *audioPlayer) play(loop int) {

	s.ctrl = &beep.Ctrl{Streamer: beep.Loop(loop, s.stream), Paused: false}
	speaker.Init(s.format.SampleRate, s.format.SampleRate.N(time.Second/10))
	s.stream.Seek(0)
	s.ctrl.Paused = false
	speaker.Play(s.ctrl)
}

// TODO Distinguish betweeen pause and stopped??
func (s *audioPlayer) IsStopped() bool {

	return s.ctrl.Paused || s.stream.Position() == s.stream.Len()
}

func (s *audioPlayer) IsPlaying() bool {
	return !s.IsStopped()
}
func (s *audioPlayer) Stop() {
	//
	log.Println("Stop called")

	speaker.Lock()
	s.ctrl.Paused = true
	speaker.Unlock()

}

func (s *audioPlayer) Pause() {
	log.Println("Pause called...")
	speaker.Lock()
	s.ctrl.Paused = true
	speaker.Unlock()

}

func (s *audioPlayer) Unload() {

}
