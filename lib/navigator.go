package lib

//TODO Standardize the Point interface and struct and create appropriate helper methods.

func NewLocation() Navigator {
	return &NavigatorData{}
}

func NewNavigator(x, y, xMove, yMove, maxXOffset, maxYOffset int32, controller ActionStageController) Navigator {
	return &NavigatorData{x, y, xMove, yMove, maxXOffset, maxYOffset, controller}
}

type Navigator interface {
	MoveTowardMouse(x, y int32)
	MoveByIncrement(x, y int32)
	CalculateMoveIncrement(moveToX, moveToY int32) (x, y int32)
	GetCurrentPoint() (x, y int32)
	ActionStageController
}

func (cl *NavigatorData) Add(x, y int32) {
	cl.X += x
	cl.Y += y
}

// Clone
func (cl *NavigatorData) GetCurrentPoint() (x, y int32) {
	return cl.X, cl.Y
}

func (cl *NavigatorData) MoveByIncrement(x, y int32) {

	if cl.isLegal(x, y) {
		cl.Add(x, y)
	}

}

func (cl *NavigatorData) isLegal(x, y int32) bool {
	var isLegal bool = true
	if cl.ActionStageController != nil {
		//Sample mechanisms for determining legal behavior...
		absoluteX, absoluteY := cl.GetCurrentPoint()
		absoluteX += x
		absoluteY += y
		color := cl.ActionStageController.CheckBehaviorColorAt(absoluteX, absoluteY)

		r, g, b, _ := color.RGBA()
		//TODO Should probably use black as legal color so we just check 0
		isLegal = 65535 == r && 65535 == g && 65535 == b
	}
	return isLegal
}

func (cl *NavigatorData) MoveTowardMouse(x, y int32) {
	cl.MoveByIncrement(cl.CalculateMoveIncrement(x, y))

}

// This is in the middle of refactoring.
func (cl *NavigatorData) CalculateMoveIncrement(moveToX, moveToY int32) (x, y int32) {

	incrementX, incrementY := cl.increment(moveToX, moveToY)
	//Offset checks could be cleaner...
	var xOffset = moveToX - cl.X
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < cl.MaxXOffset {
		incrementX = 0
	}
	yOffset := moveToY - cl.Y
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < cl.MaxYOffset {
		incrementY = 0
	}

	return incrementX, incrementY
}

func (cl *NavigatorData) increment(moveToX, moveToY int32) (incrementX, incrementY int32) {
	x, y := cl.GetCurrentPoint()
	incrementX = cl.XMove
	incrementY = cl.YMove
	if moveToX < x {
		incrementX *= -1
	}

	if moveToY < y {
		incrementY *= -1
	}
	return incrementX, incrementY
}
