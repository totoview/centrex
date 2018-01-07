package cmd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	"github.com/spf13/cobra"
	"github.com/totoview/centrex/log"
)

var (
	loggingLevel string
	verbose      bool
	maxProcs     int
	cpuProfile   string
	traceFile    string
	memProfile   string
	logWithTs    bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&loggingLevel, "logging", "l", "info", "Set logging `level` to one of debug, info, warn or error")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output. Same as setting logging level to debug")
	RootCmd.PersistentFlags().BoolVarP(&logWithTs, "ts", "T", false, "Display timestamps in logging")
	RootCmd.PersistentFlags().IntVarP(&maxProcs, "maxprocs", "M", runtime.GOMAXPROCS(0), "Set the maximum number of CPUs to `n`")
	RootCmd.PersistentFlags().StringVarP(&cpuProfile, "prof", "p", "", "save cpu profile to `file`")
	RootCmd.PersistentFlags().StringVarP(&traceFile, "trace", "t", "", "save trace to `file`")
	RootCmd.PersistentFlags().StringVarP(&memProfile, "memprof", "o", "", "save object allocation (heap) profile to `file`")
	RootCmd.AddCommand(centrexCmd)
}

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "centrex",
	Short: "centrex is a high-performance data exchange",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		log.Verbose = verbose
		log.LoggingLevel = loggingLevel
		log.LogWithTs = logWithTs

		runtime.GOMAXPROCS(maxProcs)
		if cpuProfile != "" {
			f, err := os.Create(cpuProfile)
			if err != nil {
				cpuProfile = ""
				fmt.Fprintf(os.Stderr, "*** Failed to create cpu profile file %v\n", err)
			} else {
				pprof.StartCPUProfile(f)
			}
		}
		if traceFile != "" {
			f, err := os.Create(traceFile)
			if err != nil {
				traceFile = ""
				fmt.Fprintf(os.Stderr, "*** Failed to create trace file %v\n", err)
			} else {
				trace.Start(f)
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if cpuProfile != "" {
			pprof.StopCPUProfile()
		}
		if traceFile != "" {
			trace.Stop()
		}
		if memProfile != "" {
			f, err := os.Create(memProfile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "*** Failed to create heap profile file %v\n", err)
			} else {
				runtime.GC()
				if err := pprof.WriteHeapProfile(f); err != nil {
					fmt.Fprintf(os.Stderr, "*** Failed to write heap profile file %v\n", err)
				} else {
					f.Close()
				}
			}
		}
	},
}
