package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

const (
    ColorReset  = "\033[0m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorCyan   = "\033[36m"
    ColorRed    = "\033[31m"
)

const (
    header = ColorGreen + `
########################################
#                                      #
#            NATUREBYTE                #
#            FORAGER TOOL              #
#                                      #
########################################
` + ColorReset

    separator = ColorGreen + "========================================================\n" + ColorReset
)

func listSubdomains(domain string) ([]string, error) {
    fmt.Println(ColorCyan + "Foraging subdomains for: " + ColorGreen + domain + ColorReset)
    fmt.Println(separator)

    _, err := exec.LookPath("subfinder")
    if err != nil {
        return nil, fmt.Errorf("subfinder is not installed or not found in PATH")
    }

    cmd := exec.Command("subfinder", "-d", domain)
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    subdomains := strings.Split(string(output), "\n")
    return subdomains, nil
}

func runReconTools(domain string) {
    tools := []string{"theHarvester", "dnsrecon", "whois"}

    fmt.Println(ColorCyan + "Foraging some more on: " + ColorGreen + domain + ColorReset)
    fmt.Println(separator)

    for _, tool := range tools {
        var cmd *exec.Cmd
        if tool == "theHarvester" {
            cmd = exec.Command(tool, "-d", domain, "-b", "all")
        } else if tool == "dnsrecon" {
            cmd = exec.Command(tool, "-d", domain)
        } else {
            cmd = exec.Command(tool, domain)
        }

        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Printf(ColorRed+"Warning: Error running recon tool '%s': %v\n"+ColorReset, tool, err)
            fmt.Printf(ColorRed+"Command output: %s\n"+ColorReset, string(output))
            continue
        }

        result := filterOutput(tool, string(output))
        if result != "" {
            fmt.Printf(ColorGreen+"Results from foraging:\n"+ColorReset+"%s\n", result)
            fmt.Println(separator)
        }
    }
}

func filterOutput(tool, output string) string {
    var filteredOutput string
    switch tool {
    case "theHarvester":
        filteredOutput = parseTheHarvesterOutput(output)
    case "dnsrecon":
        filteredOutput = parseDnsreconOutput(output)
    case "whois":
        filteredOutput = parseWhoisOutput(output)
    }
    return filteredOutput
}

func parseTheHarvesterOutput(output string) string {
    lines := strings.Split(output, "\n")
    var results []string
    capture := false
    for _, line := range lines {
        if strings.Contains(line, "[*]") && strings.Contains(line, "found") {
            capture = true
        }
        if capture && !strings.Contains(line, "[*] Searching") && line != "" {
            results = append(results, line)
        }
    }
    return filterResultsWithZero(results)
}

func parseDnsreconOutput(output string) string {
    lines := strings.Split(output, "\n")
    var results []string
    for _, line := range lines {
        if strings.Contains(line, "[*]") || strings.Contains(line, "[+]") {
            results = append(results, line)
        }
    }
    return strings.Join(results, "\n")
}

func parseWhoisOutput(output string) string {
    lines := strings.Split(output, "\n")
    var results []string
    var section string
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" || strings.HasPrefix(line, "%") || strings.Contains(line, "Database") ||
            strings.Contains(line, "query") || strings.Contains(line, "terms of use") ||
            strings.Contains(line, ">>>") || strings.HasPrefix(line, "URL of the ICANN WHOIS") ||
            strings.Contains(line, "For more information on Whois status codes") {
            continue
        }
        parts := strings.Split(line, ":")
        if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            if key == "Domain Name" {
                if section != "" {
                    results = append(results, section)
                }
                section = ColorYellow + "Domain Information:\n" + ColorReset
            } else if key == "Registrant Name" || key == "Admin Name" || key == "Tech Name" {
                section += ColorYellow + "\n" + key + ":\n" + ColorReset
            }
            section += ColorGreen + "  " + key + ": " + ColorReset + value + "\n"
        }
    }
    if section != "" {
        results = append(results, section)
    }
    return strings.Join(results, "\n")
}

func filterResultsWithZero(results []string) string {
    var filtered []string
    for _, result := range results {
        if !strings.Contains(result, "found: 0") {
            filtered = append(filtered, result)
        }
    }
    return strings.Join(filtered, "\n")
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println(ColorYellow + "Usage: ./forager domain.com" + ColorReset)
        return
    }

    fmt.Println(header)

    domain := os.Args[1]

    subdomains, err := listSubdomains(domain)
    if err != nil {
        fmt.Println(ColorYellow + "Error listing subdomains:" + ColorReset, err)
        return
    }

    fmt.Println(ColorCyan + "Subdomains found:" + ColorReset)
    for _, subdomain := range subdomains {
        if subdomain != "" {
            fmt.Println(ColorGreen + subdomain + ColorReset)
        }
    }

    fmt.Println(separator)
    runReconTools(domain)

    fmt.Println(separator)
}
