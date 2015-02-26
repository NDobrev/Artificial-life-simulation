# Artificial-life-simulation

## Install
```
 git clone https://github.com/NDobrev/Artificial-life-simulation.git
 
 Windows: just start Build_x64.bat or Build_x86.bat and  start gui/web/index.html
 Linux/Mac: build manual main.go and start gui/web/index.html
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
