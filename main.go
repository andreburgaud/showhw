package main

import ( 
    "fmt"
    "math"
    "os"
    //"strings"

    "github.com/jaypipes/ghw"
    "github.com/jaypipes/ghw/pkg/unitutil"
)


// func product() {
//     product, err := ghw.Product()
//     if err != nil {
//         fmt.Printf("Error getting product info: %v", err)
//     }

//     fmt.Printf("%v\n", product)
// }


func baseboard() {
    baseboard, err := ghw.Baseboard()
    if err != nil {
        fmt.Printf("Error getting baseboard info: %v", err)
    }

    fmt.Printf("%v\n", baseboard)
}

func bios() {
        bios, err := ghw.BIOS()
    if err != nil {
        fmt.Printf("Error getting BIOS info: %v", err)
    }

    fmt.Printf("%v\n", bios)
}

func chassis() {
    chassis, err := ghw.Chassis()
    if err != nil {
        fmt.Printf("Error getting chassis info: %v", err)
    }

    fmt.Printf("%v\n", chassis)
}

func gpu() {
    gpu, err := ghw.GPU()
    if err != nil {
        fmt.Printf("Error getting GPU info: %v", err)
    }

    fmt.Printf("%v\n", gpu)

    for _, card := range gpu.GraphicsCards {
        fmt.Printf(" %v\n", card)
    }
}

func pci() {
        pci, err := ghw.PCI()
    if err != nil {
        fmt.Printf("Error getting PCI info: %v", err)
    }
    fmt.Printf("host PCI devices:\n")
    fmt.Println("====================================================")

    for _, device := range pci.Devices {
        vendor := device.Vendor
        vendorName := vendor.Name
        if len(vendor.Name) > 20 {
            vendorName = string([]byte(vendorName)[0:17]) + "..."
        }
        product := device.Product
        productName := product.Name
        if len(product.Name) > 40 {
            productName = string([]byte(productName)[0:37]) + "..."
        }
        fmt.Printf("%-12s\t%-20s\t%-40s\n", device.Address, vendorName, productName)
    }

    addr := "0000:00:00.0"
    if len(os.Args) == 2 {
        addr = os.Args[1]
    }
    fmt.Printf("PCI device information for %s\n", addr)
    fmt.Println("====================================================")
    deviceInfo := pci.GetDevice(addr)
    if deviceInfo == nil {
        fmt.Printf("could not retrieve PCI device information for %s\n", addr)
        return
    }

    vendor := deviceInfo.Vendor
    fmt.Printf("Vendor: %s [%s]\n", vendor.Name, vendor.ID)
    product := deviceInfo.Product
    fmt.Printf("Product: %s [%s]\n", product.Name, product.ID)
    subsystem := deviceInfo.Subsystem
    subvendor := pci.Vendors[subsystem.VendorID]
    subvendorName := "UNKNOWN"
    if subvendor != nil {
        subvendorName = subvendor.Name
    }
    fmt.Printf("Subsystem: %s [%s] (Subvendor: %s)\n", subsystem.Name, subsystem.ID, subvendorName)
    class := deviceInfo.Class
    fmt.Printf("Class: %s [%s]\n", class.Name, class.ID)
    subclass := deviceInfo.Subclass
    fmt.Printf("Subclass: %s [%s]\n", subclass.Name, subclass.ID)
    progIface := deviceInfo.ProgrammingInterface
    fmt.Printf("Programming Interface: %s [%s]\n", progIface.Name, progIface.ID)
}


func network() {
        net, err := ghw.Network()
    if err != nil {
        fmt.Printf("Error getting network info: %v", err)
    }

    fmt.Printf("%v\n", net)

    for _, nic := range net.NICs {
        fmt.Printf(" %v\n", nic)

        enabledCaps := make([]int, 0)
        for x, cap := range nic.Capabilities {
            if cap.IsEnabled {
                enabledCaps = append(enabledCaps, x)
            }
        }
        if len(enabledCaps) > 0 {
            fmt.Printf("  enabled capabilities:\n")
            for _, x := range enabledCaps {
                fmt.Printf("   - %s\n", nic.Capabilities[x].Name)
            }
        }
    }
}

func topology() {
        topology, err := ghw.Topology()
    if err != nil {
        fmt.Printf("Error getting topology info: %v", err)
    }

    fmt.Printf("%v\n", topology)

    for _, node := range topology.Nodes {
        fmt.Printf(" %v\n", node)
        for _, cache := range node.Caches {
            fmt.Printf("  %v\n", cache)
        }
    }
}

func storage() {
    block, err := ghw.Block()
    if err != nil {
        fmt.Printf("Error getting block storage info: %v", err)
    }

    fmt.Printf("%v\n", block)

    for _, disk := range block.Disks {
        fmt.Printf(" %v\n", disk)
        for _, part := range disk.Partitions {
            fmt.Printf("  %v\n", part)
        }
    }
}







/////////////////

// Product

func product() {
    title("Product")
    product, err := ghw.Product(ghw.WithDisableWarnings())
    if err != nil {
        fmt.Printf("Error getting product info: %v", err)
    }
    fmt.Printf("Name: %s\n", product.Name)
    fmt.Printf("Vendor: %s\n", product.Vendor)
}

func cpu() {
    title("CPU")
    cpu, err := ghw.CPU(ghw.WithDisableWarnings())
    if err != nil {
        fmt.Printf("Error getting CPU info: %v", err)
    }

    cpu_count := len(cpu.Processors)
    cpu_string := "processor"
    if cpu_count > 1 {
        cpu_string = cpu_string + "s"
    }

    fmt.Printf("%d %s, ", cpu_count, cpu_string)
    fmt.Printf("%d cores, ", cpu.TotalCores)
    fmt.Printf("%d threads\n", cpu.TotalThreads)

    for _, proc := range cpu.Processors {
        fmt.Printf("Vendor: %s\n", proc.Vendor)
        fmt.Printf("Model: %s\n", proc.Model)
    }
}


func memory() {
    title("Memory")
    memory, err := ghw.Memory(ghw.WithDisableWarnings())
    if err != nil {
        fmt.Printf("Error getting memory info: %v", err)
    }

    fmt.Println(memory.String())

    phys := memory.TotalPhysicalBytes

    fmt.Println("Physical Memory %d (bytes)", phys)
    usable := memory.TotalUsableBytes
    fmt.Println("Usable Memory %d (bytes)", usable)


    unit, unitStr := unitutil.AmountString(phys)
    total := int64(math.Ceil(float64(phys) / float64(unit)))
    fmt.Printf("Physical Memory: %d%s\n", total, unitStr)
}

////////////////


func title(text string) {
    fmt.Printf("\n%s\n", text)
}

func main() {

    // title("Memory")
    // memory()
    // title("Storage")
    // storage()
    // title("Topology")
    // topology()
    // title("Network")
    // network()
    // title("PCI")
    // pci()
    // title("GPU")
    // gpu()
    // title("Chassis")
    // chassis()
    // title("BIOS")
    // bios()
    // title("Baseboard")
    // baseboard()

    product()
    cpu()
    memory()
}
