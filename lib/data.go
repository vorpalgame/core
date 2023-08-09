package lib

// Common struct used by all that are also yaml annotated for marshaling.
// Marshaling creates issues when using interfaces.

type ImageLayerData struct {
	LayerMetadata []*ImageMetadata `yaml:"LayerMetadata"`
}

func UnmarshalImageLayer(fileName string) *ImageLayerData {
	value := ImageLayerData{}
	UnmarshalFile(fileName, &value)
	return &value
}

func MarshalImageLayer(fileName string, value *ImageLayerData) {
	Marshal(fileName, value)
}

/////////////////////////////////////////////////////////////////////////

type ImageMetadata struct {
	ImageFileName  string `yaml:"ImageFileName"`
	X              int32  `yaml:"X"`
	Y              int32  `yaml:"Y"`
	Width          int32  `yaml:"Width"`
	Height         int32  `yaml:"Height"`
	HorizontalFlip bool   `yaml:"HorizontalFlip"`
}

func UnmarshalImageMetadata(fileName string) *ImageMetadata {
	value := ImageMetadata{}
	UnmarshalFile(fileName, &value)
	return &value
}

func MarshalImageMetadata(fileName string, imgLayer *ImageMetadata) {
	Marshal(fileName, imgLayer)
}

// /////////////////////////////////////////////////////////////////////////////
type NavigatorData struct {
	X                     int32 `yaml:"CurrentX"`
	Y                     int32 `yaml:"CurrentY"`
	XMove                 int32 `yaml:"XMove"`
	YMove                 int32 `yaml:"YMove"`
	MaxXOffset            int32 `yaml:"MaxXOffset"`
	MaxYOffset            int32 `yaml:"MaxYOffset"`
	ActionStageController `yaml:"-"`
}

func UnmarshalNavigator(fileName string) *NavigatorData {
	value := NavigatorData{}
	UnmarshalFile(fileName, &value)
	return &value
}

func MarshalNavigator(fileName string, value *NavigatorData) {
	Marshal(fileName, value)
}

// ///////////////////////////////////////////////////////////////////////////////
type Scene struct {
	WindowTitle  string          `yaml:"WindowTitle"`
	WindowWidth  int             `yaml:"WindowWidth"`
	WindowHeight int             `yaml:"WindowHeight"`
	RegisterKeys []rune          `yaml:"RegisterKeys"`
	Background   *ImageLayerData `yaml:"Background"`
	BehaviorMap  *ImageLayerData `yaml:"BehaviorMap"`
	Foreground   *ImageLayerData `yaml:"Foreground"`
	Actors       []string        `yaml:"Actors"`
}

func UnmarshalScene(fileName string) *Scene {
	value := Scene{}
	UnmarshalFile(fileName, &value)
	return &value
}

func MarshalScene(fileName string, value *Scene) {
	Marshal(fileName, value)
}

// ///////////////////////////////////////////////////////////////////////////////
type FontData struct {
	Font     string `yaml:"Font"`
	FontSize int32  `yaml:"FontSize"`
}

// ///////////////////////////////////////////////////////////////////////////////
type TextLineData struct {
	Test string `yaml:"Text"`
	Font
}

/////////////////////////////////////////////////////////////////////////////////
