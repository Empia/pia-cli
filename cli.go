package main

import "os"
import "github.com/codegangsta/cli"

func main() {
  app := cli.NewApp()
  app.Name = "greet"
  app.Usage = "fight the loneliness!"
  app.Flags = []cli.Flag {
    cli.StringFlag{"lang, l", "english", "language for the greeting"},
  }
  app.Action = func(c *cli.Context) {
    name := "someone"
  if len(c.Args()) > 0 {
    name = c.Args()[0]
  }
  if c.String("lang") == "spanish" {
    println("Hola", name)
  } else {
    println("Hello", name)
  }
  }
app.Commands = []cli.Command{
  {
    Name:      "new",
    ShortName: "n",
    Usage:     "add a task to the list",
    Action: func(c *cli.Context) {
      println("added task: ", c.Args()[0])
    },
  },
  {
    Name:      "update",
    ShortName: "u",
    Usage:     "complete a task on the list",
    Action: func(c *cli.Context) {
      println("completed task: ", c.Args()[0])
    },
  },
  {
    Name:      "keys",
    ShortName: "k",
    Usage:     "complete a task on the list",
    Action: func(c *cli.Context) {
      println("completed task: ", c.Args()[0])
    },
  },
  {
    Name:      "search",
    ShortName: "s",
    Usage:     "complete a task on the list",
    Action: func(c *cli.Context) {
      println("completed task: ", c.Args()[0])
    },
  },
  {
    Name:      "login",
    ShortName: "l",
    Usage:     "complete a task on the list",
    Action: func(c *cli.Context) {
      println("completed task: ", c.Args()[0])
    },
  },
  {
    Name:      "status",
    ShortName: "st",
    Usage:     "complete a task on the list",
    Action: func(c *cli.Context) {
      println("completed task: ", c.Args()[0])
    },
  },
}


  app.Run(os.Args)
}
