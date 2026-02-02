# Alfred WCAG Search

An ultra-fast, offline Alfred workflow for searching the latest version of the Web Content Accessibility Guidelines (WCAG) 2.2. Find success criteria, techniques and best practices in milliseconds without leaving your keyboard.

## Features

- **Fast**: Written in Go with embedded data. No external API calls or slow JSON parsing at runtime.
- **Intelligent search**: Uses a tiered search to find what you need, even when you aren't exact.
- **Fuzzy matching**: Matching finds results even with partial strings (e.g. "cntrst" matches "Contrast").
- **Levenshtein fallback**: If fuzzy search fails, it calculates the "edit distance" to catch typos (e.g. "cnotrast" still finds "Contrast").
- **Universal binary**: Native support for both Apple Silicon and Intel Macs.

### Installation

#### Option 1: From the package folder (Recommended)

1. Go to the [package](./package) folder.
2. Download the `.alfredworkflowfile`.
3. Double-click the file to import it into Alfred.

#### Option 2: From Source

If you want to build the workflow yourself, you'll need Go 1.25+ installed.

1. Clone the repository:
```sh
git clone https://github.com/tricinel/alfred-wcag-search.git
cd alfred-wcag-search
```

2. Build and package the workflow:
```sh
make package
```

3. Find your fresh build in the `package` directory:
```sh
open package
```

4. Double-click `alfred-wcag-search.alfredworkflow` to install.

### Troubleshooting

Since this is a compiled binary, you might run into issues with Apple Gatekeeper
complaining it cannot verify if it's malware. You can fix this with:

1. Right click and Open in Finder inside the Alfred Workflow.
2. Copy the file `wcag-search`.
3. Open Terminal and run: `xattr -d com.apple.quarantine <PASTE WHAT YOU
   COPIED>`

### License

This project is licensed under the MIT License. WCAG data is sourced from the W3C.

