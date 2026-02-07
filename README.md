# 📦 Porigo

Porigo is a simple and lightweight port scanner built with Golang, designed to quickly check whether a port is open or closed.

Unlike tools like Nmap, which are extremely powerful but feature-heavy, Porigo focuses on:

⚡ Speed

🪶 Simplicity

🧠 Minimal overhead

This makes Porigo ideal for quick port checks and automation workflows.

### visit porigo web page https://porigo.ekasukra.my.id
## ✨ Features
⚡ Fast Scanning

Scan multiple ports quickly and efficiently.

🪶 Lightweight

Minimal resource usage for optimal performance.

🧩 Easy to Use

Simple commands for quick setup and usage.

🌍 Open Source

Completely free for transparency and collaboration.

## 📥 Installation

Step 1 — Download Binary

**Download Porigo from the releases page:**
```bash
https://github.com/synckaa/porigo/releases
```

**Step 2 — Give Execute Permission**
```bash
sudo chmod +x porigo
```

**Step 3 — Move Binary to PATH**
```bash
sudo mv porigo /usr/local/bin/
```

**Step 4 — Verify Installation**
```bash
porigo --version
```

## 🚀 Usage

**Scan Target with Default Ports (1–1024)**
```bash
porigo scan -t example.com
```
**Scan Target with Custom Port Range**
```bash
porigo scan -p 1-8080 -t example.com
```

**Scan Target with Single Port**
```bash
porigo scan -p 80 -t example.com
```

**Scan Target and Export Results to JSON**
```bash
porigo scan -t example.com -d
```

## 🗑 Uninstall

**Step 1 — Check Binary Location**
```bash
which porigo
```

**Step 2 — Remove Binary**
```bash
sudo rm /usr/local/bin/porigo
```

**Step 3 — Verify Removal**
```bash
porigo --version
```

If uninstalled correctly, command will not be found.

## 👨‍💻 Author

Eka Sukra | synckaa

GitHub: https://github.com/synckaa
