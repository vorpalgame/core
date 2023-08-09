package render

///No x11 packages here...
import (
	"github.com/vorpalgame/core/bus"
	"github.com/vorpalgame/core/lib"
	"golang.org/x/image/draw"
	"image"
)

// TODO separate the load, resize and cache stages for better performance and
// to avoid cache misses.
type pipelineData struct {
	drawEventChannel chan bus.DrawEvent
	outputChannel    chan *image.RGBA
	ImageCache
}

func (data *pipelineData) OnDrawEvent(inputChannel <-chan bus.DrawEvent) {
	for evt := range inputChannel {
		data.drawEventChannel <- evt
	}
}

func NewRenderPipeline(outputChannel chan *image.RGBA) {
	data := pipelineData{make(chan bus.DrawEvent, 10), outputChannel, NewImageCache()}
	bus.GetVorpalBus().AddDrawEventListener(&data)
	go renderPipelineFunc(&data, data.drawEventChannel)
}

var renderPipelineFunc = func(data *pipelineData, inputChannel <-chan bus.DrawEvent) {
	cache := data.ImageCache
	loadCacheResizeChan := make(chan bus.DrawLayersEvent, 1)
	NewLoadResizeCachePipeline(&cache, loadCacheResizeChan)
	renderChan := make(chan bus.DrawLayersEvent, 1)
	go renderImageLayersFunc(data, renderChan)
	for evt := range inputChannel {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			loadCacheResizeChan <- evt
			renderChan <- evt
		}

	}
}

var renderImageLayersFunc = func(data *pipelineData, inputChannel <-chan bus.DrawLayersEvent) {

	for evt := range inputChannel {
		var buffer *image.RGBA
		var currentImg *image.Image
		for _, layer := range evt.GetImageLayers() {
			for _, imgData := range layer.LayerMetadata {
				//TODO We need to loop on this after refactor...
				currentImg = nil
				for currentImg == nil {
					//log.Println("Wait for image...")
					currentImg = data.GetImage(imgData.ImageFileName)
				}
				//Should be first layer. Probably make this more explict.
				if buffer == nil {
					buffer = image.NewRGBA((*currentImg).Bounds())
				}
				if imgData.HorizontalFlip {
					currentImg = flip(currentImg)
				}
				draw.Draw(buffer, getRect(imgData), *currentImg, *getPoint(0, 0), draw.Over)
			}

		}
		data.outputChannel <- buffer

	}
}

func flip(img *image.Image) *image.Image {
	size := (*img).Bounds().Size()
	flipImg := image.NewNRGBA((*img).Bounds())
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			xp := size.X - x - 1
			flipImg.Set(x, y, (*img).At(xp, y))
		}
	}
	returnImg := image.Image(flipImg)
	return &returnImg
}

// ////////////////////////////////////////////////////////////////////////////
// TODO We need these on ImageMetadta interfaces...
func getRect(imgData *lib.ImageMetadata) image.Rectangle {
	return image.Rectangle{*getPoint(imgData.X, imgData.Y), *getPoint(imgData.X+imgData.Width, imgData.Y+imgData.Height)}
}

func getPoint(x, y int32) *image.Point {
	return &image.Point{int(x), int(y)}
}
