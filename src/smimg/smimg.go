package smimg

type Point struct { x, y int }
type Rectangle struct { min, max Point }

func (* Rectangle) Intersect (rec * Rectangle) {
    
}

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

func GetBoxes(data []byte, width int) []Rectangle {
    total := len(data) / 3
    rects := make([]Rectangle, 0, 10)
    var r, g, b byte
    var i, y, x, size, index, pixel int
    var rect * Rectangle
    pixel: for pixel = 0; pixel < total; pixel++ {
        i, r, g, b = pixel * 3, data[i], data[i + 1], data[i + 2]
        if r == 255 && g == 0 && b == 0 {
            y, x = pixel / width, pixel % width;
            size = len(rects)
            for index = size - 1; index >= 0; index-- {
                rect = &rects[index]
                //Horizontal
                if rect.max.y == y && rect.max.x + 1 == x {
                    rect.max.x++
                    continue pixel
                //Vertical
                } else if rect.max.y + 1 >= y &&
                        rect.min.x - 1 <= x &&
                        rect.max.x + 1 >= x {
                    rect.max.y = y
                    continue pixel
                }
            }
            rects = append(
                rects,
                Rectangle{ Point{ x, y },
                Point{ x, y }})
        }
    }
    size = len(rects)
    result := make([]Rectangle, 0, 4)
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