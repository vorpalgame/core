package lib

import (
	"gopkg.in/yaml.v3"
	"os"
)

func CreateTestNavigator() Navigator {
	maxIncrement := int32(2)
	maxOffset := int32(5)
	return &NavigatorData{500, 600, maxIncrement, maxIncrement, maxOffset, maxOffset, nil}
}
func CreateTestAudioData() AudioStateData {
	data := AudioStateData{}
	data.LoopAudio = true
	data.AudioFile = "AmazingPolkas.mp3"
	data.count = 5

	return data

}

func CreateTestFrameTrackerData() FrameTrackerData {
	data := FrameTrackerData{}
	data.IdleFrames = 1
	data.CurrentFrameRepeats = 2
	data.RepeatPerFrame = 3
	data.MaxFrame = 10
	data.LoopFrames = true
	return data
}

func WriteFile(fileName string, data interface{}) {
	file, _ := yaml.Marshal(data)

	e := os.WriteFile("../etc/"+fileName, file, 0644)
	if e != nil {
		panic(e)
	}
}
