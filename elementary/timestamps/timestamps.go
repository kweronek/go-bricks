package main
// Demo timestamps

import (
	"fmt"
	"time"
)

const longTimeStampForm = "02.01.2006 at 15:04:05.12345689 (MESZ)"

func main() {
	fmt.Println("Demo for timestamps")
	t := time.Now()
	fmt.Println(t.Format(longTimeStampForm))
	fmt.Println(t.Format("02.01.2006"))
	fmt.Println(t.Format("15:04:05.12345689 (MESZ)"))
	fmt.Println(t.Format("15:04:05.12345689"))
}

