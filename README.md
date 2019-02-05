# goscheduler [![Build Status](https://travis-ci.org/jtaczanowski/goscheduler.png?branch=master)](https://travis-ci.org/jtaczanowski/goscheduler)
goscheduler - Golang package providing function for executing task periodically. Optionaly you can add delay on task start.

Example usage (also included in example catalog)
```
// run function exampleTask() every 10s
scheduler.RunTaskAtInterval(exampleTask, time.Second*10, time.Second*0)

// run function exampleTask() every 10s, but delay on start by 3 second
scheduler.RunTaskAtInterval(exampleTask, time.Second*10, time.Second*3)

// run function exampleTaskWithArguments("example task with arguments") every 10s
scheduler.RunTaskAtInterval(
    func() { exampleTaskWithArguments("example task with arguments") },
    time.Second*10,
    time.Second*0,
)
```