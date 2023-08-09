package lib

// Most data is set once at start up so limited Setters.
func NewFrameData() FrameTracker {
	return &FrameTrackerData{1, 10, 3, 1, 0, true}
}

type FrameTracker interface {
	SetMaxFrame(maxFrame int32) FrameTracker
	SetRepeatFrame(repeatPerFrame int32) FrameTracker
	GetCurrentFrame() int32
	GetMaxFrame() int32
	SetToLoop(bool) FrameTracker
	UpdateIdleFrames(x, y int32) int32
	GetIdleFrames() int32
	IsFrameOnLoop() bool
	IncrementFrameCount() FrameTracker
	ResetFrameCount()
}
type FrameTrackerData struct {
	CurrentFrame        int32 `yaml:"CurrentFrame"`
	MaxFrame            int32 `yaml:"MaxFrame"`
	RepeatPerFrame      int32 `yaml:"RepeatPerFrame"`
	CurrentFrameRepeats int32 `yaml:"CurrentFrameRepeats"`
	IdleFrames          int32 `yaml:"IdleFrames"`
	LoopFrames          bool  `yaml:"LoopFrames"`
}

func (fd *FrameTrackerData) IncrementFrameCount() FrameTracker {
	fd.CurrentFrameRepeats++

	if fd.CurrentFrameRepeats >= fd.RepeatPerFrame {
		fd.CurrentFrameRepeats = 0
		fd.CurrentFrame++

		if fd.CurrentFrame >= fd.MaxFrame {

			if fd.IsFrameOnLoop() {
				fd.CurrentFrame = 1
			} else {
				fd.CurrentFrame = fd.MaxFrame
			}
		}
	}
	return fd
}

func (fd *FrameTrackerData) SetMaxFrame(maxFrame int32) FrameTracker {
	fd.MaxFrame = maxFrame
	return fd
}
func (fd *FrameTrackerData) SetRepeatFrame(repeatPerFrame int32) FrameTracker {
	fd.RepeatPerFrame = repeatPerFrame
	return fd
}
func (fd *FrameTrackerData) UpdateIdleFrames(x, y int32) int32 {
	if x == 0 && y == 0 {
		fd.IdleFrames++
	} else {
		fd.IdleFrames = 0
	}
	return fd.IdleFrames
}

func (fd *FrameTrackerData) GetIdleFrames() int32 {
	return fd.IdleFrames
}
func (fd *FrameTrackerData) SetToLoop(repeat bool) FrameTracker {
	fd.LoopFrames = repeat
	return fd
}

func (fd *FrameTrackerData) IsFrameOnLoop() bool {
	return fd.LoopFrames
}

func (fd *FrameTrackerData) ResetFrameCount() {
	fd.CurrentFrame = 1
	fd.IdleFrames = 0
}

func (fd *FrameTrackerData) GetCurrentFrame() int32 {
	return fd.CurrentFrame
}

func (fd *FrameTrackerData) GetMaxFrame() int32 {
	return fd.MaxFrame
}
