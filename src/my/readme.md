# 1
<pre>
package main

import (
"fmt"
"strings"
)

type VideoPlayer struct {
File string
}

type AudioPlayer struct {
File string
}

func (v VideoPlayer) Play() {
fmt.Println("Playing video", v.File)
}

func (v VideoPlayer) Stop() {
fmt.Println("Stop playing video", v.File)
}

func (a AudioPlayer) Play() {
fmt.Println("Playing audio", a.File)
}

func (a AudioPlayer) Stop() {
fmt.Println("Stop playing audio", a.File)
}

// Create the `DigitalPlayer` interface that implements the `Play()` and `Stop()` methods.
type DigitalPlayer interface {
Play()
Stop()
}

// nolint: gosimple // DO NOT remove this comment!
func main() {
// DO NOT delete this code block; it takes as an input the name of the file to play or stop.
var fileName string
fmt.Scanln(&fileName)

    // Create the `player` variable of the `DigitalPlayer` type below:
    var player DigitalPlayer

    switch {
    case strings.HasSuffix(fileName, ".mp3"):
        // Make the `player` an `AudioPlayer` below:
        player = AudioPlayer{File: fileName}

    case strings.HasSuffix(fileName, ".mp4"):
        // Make the `player` a `VideoPlayer` below:
        player = VideoPlayer{File: fileName}

    default:
        fmt.Println("Unsupported file format")
        return
	}

	// Call the `Play()` and `Stop()` methods on the `player` below:
	player.Play()
    player.Stop()
}
</pre>

# 2 
<pre>
package main

import (
	"fmt"
	"math"
)

// Do not change the type declarations below!
type (
	Square struct {
		Side float64
	}

	Circle struct {
		Radius float64
	}

	Shape interface {
		Area() float64
	}
)

// DO NOT change these lines -- they create the proper output string:
func (s Square) String() string { return fmt.Sprintf("Square area: %.2f", s.Area()) }
func (c Circle) String() string { return fmt.Sprintf("Circle area: %.2f", c.Area()) }

// Implement the interface methods for the 'Square' and 'Circle' structs below:
func (s Square) Area() float64 { return s.Side * s.Side }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func main() {
	// Do NOT change the code below! This reads the input and outputs as required.
	var length float64
	fmt.Scanln(&length)

	for _, shape := range []Shape{Square{length}, Circle{length / 2}} {
		fmt.Println(shape)
	}
}
</pre>
