package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Rename after integration...

var testMetadataFile = "testdata/testMetadataFile.yaml"
var testLayerFile = "testdata/testLayerFile.yaml"
var testNavigatorFile = "testdata/testNavigatorFile.yaml"
var testSceneFile = "testdata/testSceneFile.yaml"

func TestRoundtripImageMetadata(t *testing.T) {
	out := createImageData()
	MarshalImageMetadata(testMetadataFile, out)
	in := UnmarshalImageMetadata(testMetadataFile)
	assert.Equal(t, out, in)
}
func TestRoundTripImageLayer(t *testing.T) {
	out := createImageLayer()
	MarshalImageLayer(testLayerFile, out)
	in := UnmarshalImageLayer(testLayerFile)
	assert.Equal(t, out, in)
}

func TestRoundTripNavigatorData(t *testing.T) {
	out := createNavigatorData()
	MarshalNavigator(testNavigatorFile, out)
	in := UnmarshalNavigator(testNavigatorFile)
	assert.Equal(t, out, in)
}

func TestRoundTripScene(t *testing.T) {
	out := createScene()
	MarshalScene(testSceneFile, out)
	in := UnmarshalScene(testSceneFile)
	assert.Equal(t, out, in)
}

// ////////////// Helper functions ////////////////////////////////////////////
func createImageLayer() *ImageLayerData {
	layer := ImageLayerData{}
	layer.LayerMetadata = append(layer.LayerMetadata, createImageData())
	return &layer

}
func createImageData() *ImageMetadata {
	return &ImageMetadata{"samples/resources/zombiecide/background.png", 0, 0, 1920, 1080, false}
}

func createNavigatorData() *NavigatorData {
	return &NavigatorData{180, 200, 5, 5, 5, 5, nil}

}

func createScene() *Scene {
	return &Scene{"Window Title", 1920, 1080, []rune{'e', 'E', 'R', 'r', 'g', 'G', 'h', 'H'}, createImageLayer(), createImageLayer(), createImageLayer(), []string{"./samples/etc/henry.yaml"}}
}
