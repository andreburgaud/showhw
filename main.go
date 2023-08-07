package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
	"github.com/jaypipes/ghw"
)

const banner = ` __ _                          __    __ 
/ _\ |__   _____      __/\  /\/ / /\ \ \
\ \| '_ \ / _ \ \ /\ / / /_/ /\ \/  \/ /
_\ \ | | | (_) \ V  V / __  /  \  /\  /
\__/_| |_|\___/ \_/\_/\/ /_/    \/  \/

`

const padding = 10

const (
	KB = iota + 1
	MB
	GB
	TB
	PB
)

func printInfo(text string) {
	green := color.New(color.FgHiGreen, color.Bold)
	green.Print(text)
}

func paddWithSpaces(text string, count int) string {
	return fmt.Sprintf("%-*s", count, text)
}

func showBanner() {
	printInfo(banner)
}

func storageUnit(unit float64) float64 {
	return math.Pow(math.Pow(2, 10), unit)
}

func ceilUnit(s float64, unit float64) int64 {
	return int64(math.Ceil(float64(s) / storageUnit(unit)))
}

func formatSize(size int64) string {
	s := float64(size / 1000.0)
	if s < storageUnit(KB) {
		return fmt.Sprintf("%dKB", int64(math.Ceil(s)))
	}
	if s < storageUnit(MB) {
		return fmt.Sprintf("%dMB", ceilUnit(s, KB))
	}
	if s < storageUnit(GB) {
		return fmt.Sprintf("%dGB", ceilUnit(s, MB))
	}
	if s < storageUnit(TB) {
		return fmt.Sprintf("%dTB", ceilUnit(s, GB))
	}
	return fmt.Sprintf("%dUNKOWN", size)
}

func gpu() {
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	}

	for idx, card := range gpu.GraphicsCards {
		printInfo(paddWithSpaces(fmt.Sprintf("GPU%d", idx+1), padding))
		fmt.Printf(": %s %s\n", card.DeviceInfo.Product.Name, card.DeviceInfo.Vendor.Name)
	}
}

func storage() {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	idx := 1
	for _, disk := range block.Disks {
		if disk.Model == "unknown" {
			continue
		}
		printInfo(paddWithSpaces(fmt.Sprintf("Disc%d", idx), padding))
		fmt.Printf(": %s %s %s\n", formatSize(int64(disk.SizeBytes)), disk.DriveType, disk.Model)
		idx += 1
	}
}

func product() {
	product, err := ghw.Product(ghw.WithDisableWarnings())
	if err != nil {
		fmt.Printf("Error getting product info: %v", err)
	}
	printInfo(paddWithSpaces("Product", padding))
	fmt.Printf(": %s %s\n", product.Name, product.Vendor)
}

func cpu() {
	cpu, err := ghw.CPU(ghw.WithDisableWarnings())
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}
	for idx, proc := range cpu.Processors {
		printInfo(paddWithSpaces(fmt.Sprintf("CPU%d", idx+1), padding))
		fmt.Printf(": %s (%d cores, %d threads)\n", proc.Model, cpu.TotalCores, cpu.TotalThreads)
	}
}

func memory() {
	memory, err := ghw.Memory(ghw.WithDisableWarnings())
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}
	printInfo(paddWithSpaces("Memory", padding))
	fmt.Printf(": %s\n", formatSize(memory.TotalUsableBytes))
}

func main() {
	showBanner()
	product()
	memory()
	cpu()
	storage()
	gpu()
}
