package native_engine

import (
	"github.com/vorpalgame/core/bus"
	"github.com/vorpalgame/core/lib"
	"github.com/vorpalgame/core/media/audio"
	"github.com/vorpalgame/core/media/render"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"image"
	"image/color"
	"log"
)

type engine struct {
	bus.VorpalBus
	//bus.StandardMediaPeerController
	audioCache                              audio.AudioCache
	imageCache                              render.ImageCache
	currentRenderImage, currentDisplayImage image.Image
	window                                  screen.Window
	screen                                  screen.Screen

	textEventChannel    chan bus.TextEvent
	controlEventChannel chan bus.ControlEvent
}

///////////////////////////////////////////////////////////////////////////////////////
//// Listener methods to take messages from bus channel and them to local channel
//// This permits local implementation and processing on channel to be blocking, non-blocking
//// and split on multiple pipeline channels whichever makes sense for the current engine
//// implementation. The bus conveys the message but the internal channels are set up
//// to reflect the processing and concurrency requirements.
/////////////////////////////////////////////////////////////////////////////////////////

func (e *engine) OnControlEvent(inputChannel <-chan bus.ControlEvent) {
	for evt := range inputChannel {
		e.controlEventChannel <- evt
	}
}

func (e *engine) OnTextEvent(inputChannel <-chan bus.TextEvent) {
	for evt := range inputChannel {
		e.textEventChannel <- evt
	}
}

func NewEngine() lib.Engine {

	log.Println("New native engine...")
	e := engine{imageCache: render.NewImageCache(), audioCache: audio.NewAudioCache()}

	e.textEventChannel = make(chan bus.TextEvent, 1)
	e.controlEventChannel = make(chan bus.ControlEvent, 1)

	e.VorpalBus = bus.GetVorpalBus()
	e.AddTextEventListener(&e)
	e.AddControlEventListener(&e)
	go initWindow(&e)
	go runRenderPipeline(&e)
	go audio.NewAudioPipeline(&e.audioCache)
	return &e
}

func runRenderPipeline(e *engine) {
	fromRenderPipeline := make(chan *image.RGBA, 10)
	blitChannel := make(chan *image.RGBA, 10)
	go render.NewRenderPipeline(fromRenderPipeline)
	go blitImage(e, blitChannel)
	for img := range fromRenderPipeline {
		blitChannel <- img
	}

}

var blitImage = func(e *engine, channel <-chan *image.RGBA) {

	for buffer := range channel {

		b, _ := e.screen.NewBuffer(buffer.Bounds().Max)
		draw.Draw(b.RGBA(), b.Bounds(), buffer, *getPoint(0, 0), draw.Over)
		e.window.Upload(image.Point{0, 0}, b, buffer.Bounds())
		e.window.Publish()
		b.Release()
		buffer = nil
	}
}

var initWindow = func(e *engine) {
	driver.Main(func(s screen.Screen) {
		e.screen = s
		w, err := s.NewWindow(nil)
		e.window = w
		if err != nil {
			panic(err)
			return
		}
		defer w.Release()
		mouseChannel := make(chan mouse.Event, 1)
		go processMouse(e, mouseChannel)
		lifecycleChannel := make(chan lifecycle.Event, 1)
		go processLifeCycle(e, lifecycleChannel)
		keyChannel := make(chan key.Event, 1)
		go processKey(e, keyChannel)
		for {
			switch evt := w.NextEvent().(type) {
			case lifecycle.Event:
				lifecycleChannel <- evt
			case mouse.Event:
				mouseChannel <- evt
			case key.Event:
				keyChannel <- evt
			}
		}
	})
}

// Note these are being slowly refactored away.
func (e *engine) Start() {
	log.Println("Start Native Engine...")
	go textPipeline(e, e.textEventChannel)
	go controlPipeline(e, e.controlEventChannel)
	for { //Loop until exit is called...
	}
	//Revisit the looping as we get more util components in place.
}

var controlPipeline = func(e *engine, inputChannel chan bus.ControlEvent) {
	for event := range inputChannel {
		_ = event
	}
}

var textPipeline = func(e *engine, inputChannel <-chan bus.TextEvent) {
	for event := range inputChannel {
		switch evt := event.(type) {
		case bus.MultilineTextEvent:
			_ = evt
		}
		//log.Println("Text RenderEvent...")
		//log.Println(evt)
	}
}

func pt(p fixed.Point26_6) image.Point {
	return image.Point{
		X: int(p.X+32) >> 6,
		Y: int(p.Y+32) >> 6,
	}
}

func drawText(dst *image.RGBA) {
	//draw.Draw(dst, dst.Bounds(), image.Black, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: basicfont.Face7x13,
	}
	ss := []string{
		"The quick brown fox jumps over the lazy dog.",
		"Hello, 世界.",
		"U+FFFD is \ufffd.",
	}
	for i, s := range ss {
		d.Dot = fixed.P(20, 100*i+80)
		dot0 := pt(d.Dot)
		d.DrawString(s)
		dot1 := pt(d.Dot)
		dst.SetRGBA(dot0.X, dot0.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
		dst.SetRGBA(dot1.X, dot1.Y, color.RGBA{0x00, 0x00, 0xff, 0xff})
	}
}

// TODO Rewrite the key even to use the key.Event internally.
var processKey = func(e *engine, channel <-chan key.Event) {
	for event := range channel {
		log.Println(event)
		e.SendKeyEvent(bus.NewKeyEvent(event))
	}

}

var processMouse = func(e *engine, channel <-chan mouse.Event) {
	for event := range channel {
		e.SendMouseEvent(bus.NewMouseEvent(event))
	}

}

var processLifeCycle = func(e *engine, channel <-chan lifecycle.Event) {
	for event := range channel {
		log.Println("processLifecycle")
		log.Println(event)

		if event.To == lifecycle.StageDead {
		} else if event.To == lifecycle.StageVisible {
		}
	}

}

// ////////////////////////////////////////////////////////////////////////////
// TODO We need these on ImageMetadta interfaces...
func getRect(imgData *lib.ImageMetadata) image.Rectangle {
	return image.Rectangle{*getPoint(imgData.X, imgData.Y), *getPoint(imgData.X+imgData.Width, imgData.Y+imgData.Height)}
}

func getPoint(x, y int32) *image.Point {
	return &image.Point{int(x), int(y)}
}
