package main

import (
  "flag"
  "fmt"
  "math/rand"
  "os"
  "time"
  "github.com/jedib0t/go-pretty/v6/table"
)


func main() {
  flag.Parse()
  people := flag.Args()
  if len(people) < 3 {
    panic("At least 3 people are required")
  }

  rand.Seed(time.Now().UnixNano())
  rand.Shuffle(len(people), func(i, j int) {
    people[i], people[j] = people[j], people[i]
  })

  teamATable := table.NewWriter()
  teamATable.SetOutputMirror(os.Stdout)
  teamATable.AppendHeader(table.Row{"", "Spymaster", "Operative"})
  teamBTable := table.NewWriter()
  teamBTable.SetOutputMirror(os.Stdout)
  teamBTable.AppendHeader(table.Row{"", "Spymaster", "Operative"})

  for i := 0; i < len(people); i += 1 {
    teamIndex := i % 2
    name := people[i]

    if teamIndex == 0 {
      if i == 0 { // Red Team
        teamATable.AppendRow(table.Row{name, "✓", ""}) // Spymaster
      } else {
        teamATable.AppendRow(table.Row{name,  "", "✓"}) // Operative
      }
    } else {
      if i - 1 == 0 { // Blue Team
        teamBTable.AppendRow(table.Row{name, "✓", ""}) // Spymaster
      } else {
        teamBTable.AppendRow(table.Row{name,  "", "✓"}) // Operative
      }
    }
  }

  fmt.Printf("Red Team\n")
  teamATable.Render()
  fmt.Printf("\nBlue Team\n")
  teamBTable.Render()
}
