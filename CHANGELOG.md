# Changelog

All notable changes to this project will be documented in this file.

## [1.0.1] - 2025-07-18
### Fixed
- Corrected the `basePath` generation to use proper pluralization via the `utils.Pluralize` function instead of manually appending `"s"`.  
  This ensures correct plural forms for entity names (e.g., "category" → "categories").
