# Alfred WCAG Search

An ultra-fast, offline Alfred workflow for searching the latest version of the Web Content Accessibility Guidelines (WCAG) 2.2. Find success criteria, techniques and best practices in milliseconds without leaving your keyboard.

## Features

- **Fast**: Written in Go with embedded data. No external API calls or slow JSON parsing at runtime.
- **Intelligent search**: Uses a tiered search to find what you need, even when you aren't exact.
- **Fuzzy matching**: Matching finds results even with partial strings (e.g. "cntrst" matches "Contrast").
- **Levenshtein fallback**: If fuzzy search fails, it calculates the "edit distance" to catch typos (e.g. "cnotrast" still finds "Contrast").
- **Universal binary**: Native support for both Apple Silicon and Intel Macs.


### License

This project is licensed under the MIT License. WCAG data is sourced from the W3C.
