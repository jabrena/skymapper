package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"
import "image/color"

//Using red palette from:
//http://www.colourlovers.com/palette/693254/heart_of_evil?c=1
//http://www.colourlovers.com/palette/698883/Thankyou_100_Lovers!
//http://miraxe.deviantart.com/art/Red-Violet-Mix-Palettes-297308640

func main() {
	//Testing functions
    rojo := colorful.Color{ 255.0/255.0, 0.0/255.0, 0.0/255.0 }
    rojito1 := colorful.Color{ 219.0/255.0, 64.0/255.0, 64.0/255.0 }
    rojito2 := colorful.Color{ 219.0/255.0, 23.0/255.0, 80.0/255.0 }
    rojito3 := colorful.Color{ 222.0/255.0, 80.0/255.0, 80.0/255.0 }
    rojito4 := colorful.Color{ 217.0/255.0, 0.0/255.0, 37.0/255.0 }
    blanco := colorful.Color{ 255.0/255.0, 255.0/255.0, 255.0/255.0 }
    verde := colorful.Color{ 0.0/255.0, 255.0/255.0, 0.0/255.0 }

    fmt.Printf("DistanceLuv:   rojo: %v\tand rojito1: %v\n", rojo.DistanceLuv(rojito1), rojito1.DistanceLuv(rojo))
    fmt.Printf("DistanceLuv:   rojo: %v\tand rojito2: %v\n", rojo.DistanceLuv(rojito2), rojito2.DistanceLuv(rojo))
    fmt.Printf("DistanceLuv:   rojo: %v\tand rojito3: %v\n", rojo.DistanceLuv(rojito3), rojito3.DistanceLuv(rojo))
    fmt.Printf("DistanceLuv:   rojo: %v\tand rojito4: %v\n", rojo.DistanceLuv(rojito4), rojito4.DistanceLuv(rojo))
    fmt.Printf("DistanceLuv:   rojo: %v\tand verde: %v\n", rojo.DistanceLuv(verde), verde.DistanceLuv(rojo))
    fmt.Printf("DistanceLuv:   rojo: %v\tand blanco: %v\n", rojo.DistanceLuv(blanco), blanco.DistanceLuv(rojo))

    //Testing functions for Skymapper

    currentColor := color.RGBA{0, 255, 0, 255} //GREEN
    //currentColor := color.RGBA{219, 64, 64, 255} //mini RED   
    patternRed := color.RGBA{255, 0, 0, 255}
    colorThreshold := 0.75

    result := IsColorClosed(currentColor, patternRed, colorThreshold)
    if(result) {
		fmt.Printf("OK")    	
    }else {
    	fmt.Printf("KO")  
    }
}


func IsColorClosed(current color.Color, pattern color.Color, colorThreshold float64) bool {
	distanceFlag := false
	r1,g1,b1,_ := current.RGBA()
	r2,g2,b2,_ := pattern.RGBA()

	cColor1 := colorful.Color{ float64(r1)/255.0, float64(g1)/255.0, float64(b1)/255.0 }
	cColor2 := colorful.Color{ float64(r2)/255.0, float64(g2)/255.0, float64(b2)/255.0 }

	if(cColor1.DistanceLuv(cColor2) <= colorThreshold) {
		distanceFlag := true
		return distanceFlag
	}

	return distanceFlag
}
