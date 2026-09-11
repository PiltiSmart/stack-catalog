# Installing PiltiSmart 'ps' CLI on Linux & macOS

This guide provides instructions to install the **PiltiSmart Enterprise CLI (`ps`)** globally on **Linux** (Ubuntu, Debian, RHEL, CentOS, Arch, Proxmox/LXC) and **macOS** (Intel & Apple Silicon M1/M2/M3/M4).

---

## ⚡ Method 1: Automatic 1-Line Global Installation (Recommended)

Run the universal installer script in your terminal:

```bash
curl -fsSL https://raw.githubusercontent.com/PiltiSmart/stack-catalog/main/install.sh | bash
```

### What This Script Does:
1. Automatically detects your Operating System (**Linux** or **macOS / Darwin**).
2. Detects your CPU architecture (**`x86_64` / `amd64`** or **`arm64` / Apple Silicon**).
3. Downloads the matching pre-compiled static binary.
4. Places the executable into `/usr/local/bin/ps` and sets executable permissions (`chmod +x`).
5. Verifies installation by running `ps version`.

---

## 📦 Method 2: Manual Installation via Pre-compiled Binaries

If you prefer downloading the binary manually:

### 1. Choose Your Architecture:

| Operating System | Architecture | Binary Download Link |
|---|---|---|
| **Linux** | `x86_64` / `amd64` | `https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-linux-amd64` |
| **Linux** | `ARM64` / `aarch64` | `https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-linux-arm64` |
| **macOS (Apple Silicon)** | `M1 / M2 / M3 / M4` (`arm64`) | `https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-darwin-arm64` |
| **macOS (Intel)** | `x86_64` | `https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-darwin-amd64` |
| **Windows** | `x86_64` | `https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-windows-amd64.exe` |

### 2. Download and Move to `/usr/local/bin`:

#### On Linux (amd64 / x86_64):
```bash
sudo curl -fsSL https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-linux-amd64 -o /usr/local/bin/ps
sudo chmod +x /usr/local/bin/ps
```

#### On macOS (Apple Silicon M1/M2/M3/M4):
```bash
sudo curl -fsSL https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-darwin-arm64 -o /usr/local/bin/ps
sudo chmod +x /usr/local/bin/ps
```

#### On macOS (Intel):
```bash
sudo curl -fsSL https://github.com/PiltiSmart/stack-catalog/releases/latest/download/ps-darwin-amd64 -o /usr/local/bin/ps
sudo chmod +x /usr/local/bin/ps
```

---

## 🔨 Method 3: Build from Source (Requires Go 1.21+)

If you have Go installed on your machine:

```bash
# Clone the repository
git clone https://github.com/PiltiSmart/stack-catalog.git
cd stack-catalog

# Compile stripped production binary
go mod tidy
go build -ldflags="-s -w" -o bin/ps main.go

# Install globally
sudo cp bin/ps /usr/local/bin/ps
sudo chmod +x /usr/local/bin/ps
```

---

## 🔍 Verifying the Installation

Check that `ps` is recognized in your terminal:

```bash
# Verify version
ps version

# Run pre-flight health diagnostics
ps doctor

# Discover stacks from GitHub catalog
ps catalog

# Deploy ThingsBoard Stack
ps install tb-stack
```

---

## 💡 Troubleshooting & Notes

### Linux Process Status Collision
On standard Linux distributions, `/bin/ps` or `/usr/bin/ps` is the standard Linux process reporting tool.
- By installing the PiltiSmart CLI into `/usr/local/bin/ps`, Linux `$PATH` will prioritize `/usr/local/bin/ps` for interactive CLI commands.
- If you ever need the native Linux process table, you can invoke it directly via `/bin/ps` or `/usr/bin/ps aux`.

### macOS Security Prompt (Gatekeeper)
If macOS blocks the binary on first execution because it was downloaded via the web, clear the quarantine attribute:
```bash
xattr -d com.apple.quarantine /usr/local/bin/ps
```

---

## 🗑️ Uninstallation

To remove `ps` from your system at any time:
```bash
sudo rm -f /usr/local/bin/ps
```
