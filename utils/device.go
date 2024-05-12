package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type CPUInfoProvider interface {
	GetSerialNumber() (string, error)
}

// WindowsCPUInfoProvider 用于获取Windows平台上的CPU信息
type WindowsCPUInfoProvider struct{}

func (w *WindowsCPUInfoProvider) GetSerialNumber() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorId")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[1]), nil
	}

	return "", fmt.Errorf("未找到CPU序列号")
}

// LinuxCPUInfoProvider 用于获取Linux平台上的CPU信息
type LinuxCPUInfoProvider struct{}

func (l *LinuxCPUInfoProvider) GetSerialNumber() (string, error) {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`Serial\s+:\s+([a-zA-Z0-9]+)`)
	match := re.FindStringSubmatch(string(output))
	if len(match) > 1 {
		return match[1], nil
	}

	return "", fmt.Errorf("未找到CPU序列号")
}

// MacOSCPUInfoProvider 用于获取macOS平台上的CPU信息
type MacOSCPUInfoProvider struct{}

func (m *MacOSCPUInfoProvider) GetSerialNumber() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number (system)") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("未找到CPU序列号")
}

// GetCPUInfoProvider 根据运行时环境返回相应平台的CPUInfoProvider实例
func GetCPUInfoProvider() CPUInfoProvider {
	switch runtime.GOOS {
	case "windows":
		return &WindowsCPUInfoProvider{}
	case "linux":
		return &LinuxCPUInfoProvider{}
	case "darwin":
		return &MacOSCPUInfoProvider{}
	default:
		return nil
	}
}

// GetMotherboardSerial 获取主板序列号
func GetMotherboardSerial() (string, error) {
	var cmd *exec.Cmd
	var re *regexp.Regexp

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("wmic", "baseboard", "get", "serialnumber")
	case "linux":
		cmd = exec.Command("dmidecode", "-t", "2")
		re = regexp.MustCompile(`Serial\s+Number:\s+(.*)`)
	case "darwin":
		cmd = exec.Command("ioreg", "-l")
		re = regexp.MustCompile(`"IOPlatformSerialNumber" = "(.*)"\n`)
	default:
		return "", fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		lines := strings.Split(string(output), "\n")
		if len(lines) > 1 {
			return strings.TrimSpace(lines[1]), nil
		}
	} else {
		match := re.FindStringSubmatch(string(output))
		if len(match) > 1 {
			return strings.TrimSpace(match[1]), nil
		}
	}

	return "", fmt.Errorf("could not parse serial number")
}

// GetHardDriveSerial 获取硬盘序列号
func GetHardDriveSerial() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("wmic", "diskdrive", "get", "serialnumber")
	case "linux":
		cmd = exec.Command("hdparm", "-I", "/dev/sda")
	case "darwin":
		cmd = exec.Command("diskutil", "info", "/dev/disk0")
	default:
		return "", fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	outputStr := string(output)

	switch runtime.GOOS {
	case "windows":
		fmt.Println(string(output))
		lines := strings.Split(string(output), "\n")
		if len(lines) > 1 {
			return strings.TrimSpace(lines[1]), nil
		}
	case "linux":
		re := regexp.MustCompile(`Serial\s+Number\s*=\s*(.*)`)
		match := re.FindStringSubmatch(outputStr)
		if len(match) > 1 {
			return strings.TrimSpace(match[1]), nil
		}
	case "darwin":
		re := regexp.MustCompile(`Serial Number:\s+(.*)`)
		match := re.FindStringSubmatch(outputStr)
		if len(match) > 1 {
			return strings.TrimSpace(match[1]), nil
		}
	}

	return "", fmt.Errorf("could not parse serial number")
}
