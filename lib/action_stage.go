package lib

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

type ActionStageController interface {
	CheckBehaviorColorAt(x, y int32) color.Color
	SetControlMap(image *image.Image, width, height int) ActionStageController
	Load(data *ImageLayerData) ActionStageController
}

type ActionStageControllerData struct {
	controlMap *image.Image
}

func (a *ActionStageControllerData) CheckBehaviorColorAt(x, y int32) color.Color {
	if a.controlMap != nil {
		img := *a.controlMap
		return img.At(int(x), int(y))
	}
	return color.White
}

func (a *ActionStageControllerData) SetControlMap(i *image.Image, width, height int) ActionStageController {

	// Resize:
	dest := resize.Resize(uint(width), uint(height), *i, resize.InterpolationFunction(resize.NearestNeighbor))
	a.controlMap = &dest

	return a
}

func writeTestFile(file string, dest image.Image) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, dest); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

// TODO This shoudl probably be the ImageMetadata and not layer. Not sure if
// a behavior map would have more than one layer but it is possible.
// In that case we'll likely need different behavior map types.
func (a *ActionStageControllerData) Load(data *ImageLayerData) ActionStageController {
	layer := data.LayerMetadata[0]
	f, err := os.Open(layer.ImageFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	a.SetControlMap(&img, int(layer.Width), int(layer.Height))
	return a
}
