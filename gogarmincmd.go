package main

import (
  "flag"
  "os"
  "fmt"
  "github.com/samonzeweb/gogarmincmd/tcx"
)

type CmdArguments struct {
    idActivity  string
    fcMax       int
}

func main() {

  arguments := getCommandLineArguments()

  tcxData, err  := tcx.GetTCX(arguments.idActivity)
  if err != nil {
    panic(err)
  }

  displayActivities(tcxData, arguments)

}

func getCommandLineArguments() *CmdArguments {
  fcmax := flag.Int("fcmax", 0, "Fréquence cardiaque maximale")
  flag.Usage = printUsageAndQuit
  flag.Parse()
  return &CmdArguments{ idActivity: getIdActivity(),
                        fcMax: *fcmax }
}

func getIdActivity() string {
  if len(flag.Args()) != 1 {
    printUsageAndQuit()
  }
  return flag.Arg(0)
}

func printUsageAndQuit() {
  // TODO utiliser basename du programme (pas tout le chemin)
  fmt.Fprintf(os.Stderr, "usage : %v [options] id_activity\n", os.Args[0])
  flag.PrintDefaults()
  os.Exit(1)
}
