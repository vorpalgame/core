package lib

/////////////////////////////////////////////////////
//// Font
/////////////////////////////////////////////////////

func NewFont(font string, size int32) Font {
	return &FontData{font, size}
}

type Font interface {
	GetFont() string
	GetFontSize() int32
	SetFont(font string) Font
	SetFontSize(int32) Font
}

func (f *FontData) SetFont(font string) Font {
	f.Font = font
	return f
}

func (f *FontData) SetFontSize(size int32) Font {
	f.FontSize = size
	return f
}

func (f *FontData) GetFont() string {
	return f.Font
}

func (f *FontData) GetFontSize() int32 {
	return f.FontSize
}

/////////////////////////////////////////////////////
////// Text Line
////////////////////////////////////////////////////

type TextLine interface {
	Font
	GetText() string
}

// //// TextLine

func NewTextLine(text string, font Font) TextLine {
	return &TextLineData{text, font}
}

func (p *TextLineData) GetText() string {
	return p.Test
}
