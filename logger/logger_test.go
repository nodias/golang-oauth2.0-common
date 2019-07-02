package logger

import (
	"context"
)

func Example_Log() {
	Log.Println("just log.Println")
	ctx := context.WithValue(context.Background(), "go", "Go")
	mylog := NewMyLogger(ctx)
	mylog.Println("mylog.Println")
	mylog.Println(ctx.Value("go"))
	//Output:
}
