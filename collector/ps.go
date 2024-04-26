package collector

import (
    "fmt"
    "time"

    "github.com/cfbeard/ps-exporter/config"
    "github.com/cfbeard/ps-exporter/log"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/shirou/gopsutil/v3/process"
)

type psCollector struct {
    descriptors map[string]*prometheus.Desc
}

func init() {
    addCollector(config.PsCollector.Key, NewPsCollector)
}

func NewPsCollector() prometheus.Collector {
    return &psCollector{
        descriptors: make(map[string]*prometheus.Desc),
    }
}

func (this *psCollector) Describe(ch chan<- *prometheus.Desc) {
    features := GetMonitoringTasks()[processStats]
    for _, feat := range features {
        ch <- this.describe(feat)
    }
}

func (this *psCollector) Collect(ch chan<- prometheus.Metric) {
    processes, err := process.Processes()
    if err != nil {
        log.Log.Fatalf("Failed to get host processes: %v", err)
    }

    for _, process := range processes {
        this.collect(ch, process)
    }
}

func (this *psCollector) describe(key string) *prometheus.Desc {
    descriptor, ok := this.descriptors[key]
    if !ok {
        name := "ps_" + key
        help := "PS " + key

        switch key {
        case "mem_percent":
            help = "Percent of memory used by a process"
        case "cpu_percent":
            help = "Percent of cpu used by a process"
        }

        labels := []string{"user", "process_name", "pid", "ppid", "start_time"}
        descriptor = prometheus.NewDesc(name, help, labels, nil)
    }

    return descriptor
}

func (this *psCollector) collect(ch chan<- prometheus.Metric, process *process.Process) {
    username, err := process.Username()
    if err != nil {
        log.Log.Errorf("Failed to get username for process '%d': ", process.Pid)
    }

    processName, err := process.Name()
    if err != nil {
        log.Log.Errorf("Failed to get name of process '%d': ", process.Pid)
    }

    ppid, err := process.Ppid()
    if err != nil {
        log.Log.Errorf("Failed to get parent pid of process '%d': ", process.Pid)
    }

    startTime, err := process.CreateTime()
    if err != nil {
        log.Log.Errorf("Failed to get start time of process '%d': ", process.Pid)
    }
    startTimeUnixFmt := msToTime(startTime)

    labels := []string{username, processName, fmt.Sprintf("%d", ppid), startTimeUnixFmt.String()}

    memPercent, err := process.MemoryPercent()
    if err != nil {
        log.Log.Errorf("Failed to get memory percentage of process '%d': ", process.Pid)
    }

    cpuPercent, err := process.CPUPercent()
    if err != nil {
        log.Log.Errorf("Failed to get cpu percentage of process '%d': ", process.Pid)
    }

    ch <- prometheus.MustNewConstMetric(
        this.describe("mem_percent"),
        prometheus.GaugeValue,
        float64(memPercent),
        labels...
    )

    ch <- prometheus.MustNewConstMetric(
        this.describe("cpu_percent"),
        prometheus.GaugeValue,
        cpuPercent,
        labels...
    )
}

func msToTime(ms int64) time.Time {
    return time.Unix(0, ms * int64(time.Millisecond))
}
