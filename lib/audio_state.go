package lib

//Having to return the struct instead of the interface is an unfortunate
//side effect of the YAML marshaling.

func NewAudioState(fileName string, loop bool) AudioState {
	return &AudioStateData{fileName, loop, 0}
}

type AudioState interface {
	SetAudioFile(fileName string) AudioState
	SetAudioLoop(shouldLoop bool) AudioState
	IsAudioOnLoop() bool
	GetAudioFile() string
	IncrementAudio() int32
	ResetAudioCount() AudioState
}

type AudioStateData struct {
	AudioFile string `yaml:"AudioFile"`
	LoopAudio bool   `yaml:"LoopAudio"`
	count     int32  `yaml:"-"`
}

func (s *AudioStateData) IncrementAudio() int32 {
	s.count++
	return s.count
}

func (s *AudioStateData) ResetAudioCount() AudioState {
	s.count = 0
	return s
}

func (s *AudioStateData) SetAudioLoop(shouldLoop bool) AudioState {
	s.LoopAudio = shouldLoop
	return s
}

func (s *AudioStateData) IsAudioOnLoop() bool {
	return s.LoopAudio
}

func (s *AudioStateData) GetAudioFile() string {
	return s.AudioFile
}
func (s *AudioStateData) SetAudioFile(fileName string) AudioState {
	s.AudioFile = fileName
	return s
}
