package render

import (
	"image"
	"sync"
)

// Minimal synchronization limited to cache access.
var (
	imageMapMutex = sync.RWMutex{}
)

type ImageCache interface {
	GetImage(fileName string) *image.Image
	CacheImage(fileName string, img *image.Image)
}
type imageCacheData struct {
	imageCache map[string]*image.Image
}

func NewImageCache() ImageCache {
	cache := imageCacheData{}
	cache.imageCache = make(map[string]*image.Image, 100)

	return &cache
}

func (c *imageCacheData) GetImage(fileName string) *image.Image {
	var img *image.Image
	imageMapMutex.Lock()
	img = c.imageCache[fileName]
	imageMapMutex.Unlock()
	return img

}

func (c *imageCacheData) CacheImage(fileName string, img *image.Image) {
	imageMapMutex.Lock()
	c.imageCache[fileName] = img
	imageMapMutex.Unlock()
}
