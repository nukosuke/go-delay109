# go-delay109

Delay information of Tokyu lines

## What's this?

This library is simply parsing Tokyu delay certificate page and check if the train is delayed. It supports all lines of Toyku railway as bellow.

- TY: Toyoko Line
- MG: Meguro Line
- DT: Den-en-toshi Line
- OM: Oimachi Line
- IK: Ikegami Line
- TM: Tokyu Tamagawa Line
- SG: Setagaya Line
- KD: Kodomonokuni Line

See [reference](https://godoc.org/github.com/nukosuke/go-delay109) for details.

## Example

### Code

```sh
package main

import (
	"fmt"

	delay109 "github.com/nukosuke/go-delay109"
)

func main() {
	client := delay109.New(nil)
	status, _ := client.Search(&delay109.Query{
		Line:       delay109.DT,
		Year:       2018,
		Month:      1,
		Day:        1,
		Direction:  delay109.DirectionUp,
		TimePeriod: delay109.Before10am,
	})

	fmt.Println(status.LineNameJa())
	fmt.Println(status.DirectionText(), "方面")
	fmt.Println(status.DelayDuration().Minutes(), "分遅延")
}
```

### Output

```sh
# This is sample, NOT actual information.
$ ./example
田園都市線
上り 方面
10 分遅延
```

## Notice

Redistributing Tokyu official delay information is prohibited.  
You can use this library only for your own purpose.

## License

MIT
