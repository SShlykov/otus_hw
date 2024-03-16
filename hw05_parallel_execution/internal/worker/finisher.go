package worker

func Finisher(context *Context, maxTasks int, doneChan <-chan struct{}) {
	var doneCount int
	for {
		select {
		case <-context.Ctx.Done():
			return
		case <-doneChan:
			doneCount++
			if doneCount >= maxTasks {
				context.Cancel()
			}
		}
	}
}
