# Artificial-life-simulation

## Install
```
  go get "https://github.com/NDobrev/Artificial-life-simulation.git"
 ```
  
## Usage

```go

package main
import (
	"Artificial-life-simulation/core"
	"Artificial-life-simulation/gui"
	//"fmt"
)

func main() {
	a := core.NewField(50)
	pla := core.NewPhytoPlankton()
	var point core.FieldPoint
	point.SetPoint(1, 1)
	a.AddObject(point, pla)
	for i := 0; i < 150; i++ {
		if i == 10 {

			a.AddObject(point, core.NewZooPlankton())
		}
		a.OnTick()

	}
	gui.WinGuiMain()

}
```
