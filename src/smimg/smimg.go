package smimg

type Color struct { R, G, B uint8 }

func (c Color) Equals(color *Color) bool {
    return c.R == c.R && c.G == color.G && c.B == color.B
}

type Point struct { x, y int }
type Rectangle struct { min, max Point }

func min(a, b int) int {
    if (a < b) {
        return a
    }
    return b
}

func max(a, b int) int {
    if (a > b) {
        return a
    }
    return b
}


// The method GetBoxes processes an image frame from a Webcam to return color Blobs. 
// The method needs a color pattern to detect blobs.
func GetBoxes(data []byte, width int, pattern Color) []Rectangle {
    total := len(data) / 3
    //TODO: 10?
    rects := make([]Rectangle, 0)
    var color Color
    var i, y, x, pixel int
    var rect * Rectangle
    for pixel = 0; pixel < total; pixel++ {
        i = pixel * 3
        color = Color{data[i], data[i + 1], data[i + 2]}
        if pattern.Equals(&color) {
            y, x = pixel / width, pixel % width;
            if rect != nil && rect.max.y == y && rect.max.x + 1 == x {
                rect.max.x++
            } else {
                rects = append(rects, Rectangle{ Point{ x, y }, Point{ x, y }})
                rect = &rects[len(rects)-1]    
            }
        }
    }
    //return rects
    size := len(rects)
    result := make([]Rectangle, 0)
    rect: for i := 0; i < size; i++ {
        for j := i + 1; j < size; j++ {
            if rects[i].max.y + 1 >= rects[j].max.y && 
                    rects[i].min.x <= rects[j].max.x &&
                    rects[i].max.x >= rects[j].min.x {
                rects[j].min.y = min(rects[i].min.y,rects[j].min.y)
                rects[j].min.y = min(rects[i].min.x,rects[j].min.x)
                rects[j].max.y = max(rects[i].max.y,rects[j].max.y)
                rects[j].max.y = max(rects[i].max.y,rects[j].max.y)
                continue rect
            }
        } 
        result = append(result, rects[i])
    }
    return result
    
}