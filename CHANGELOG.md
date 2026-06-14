# v0.1.0 - Released on May 17, 2026

## Features
1. Hachimi local security LLM: GGUF model loading via CGO, HITL evaluation, configurable context length (low/medium/high), 10-15 min auto-unload, defensive overflow protection, auto-download manager, YZMA support
2. Cron scheduler: engine runner push channel, cron expressions (persistent) & delay format +5m/+3h/+7d (in-memory), cron tool with get/set/list/delete, MessageFromCron / MessageFromCronIgnore bitmap flags for dual-track gatekeeper
3. Interactive config wizard (TUI): first-run onboard, add/edit/delete providers, database session management, config editing via $EDITOR or system default
4. New system tools: system info, shell execution, web browser, memory (KV & Markdown), file (read-only with ACL), datetime
5. Telegram provider enhancements: configurable collapse length, reaction notify modes (enabled/disabled/clean), HTML converter, callback selection, delete/edit messages, forward parsing
6. Model updates: DeepSeek V4 Pro/Flash, Kimi K2.6, Xiaomi MiMo v2.5 Pro/v2.5/v2 Omni
7. Database cost tracking in chat_data, Cron & Soul tables
8. New Makefile with cross-platform builds (linux/mac/windows, amd64/arm64/riscv64), version type system (stable/rc/beta/alpha/canary)
9. schema: add validation support to the schema/args conversion layer

## Refactors
1. CLI split into manboster / manbodev entrypoints
2. Chat & LLM providers moved to spec/ packages
3. Repository split from monolith into smaller repos
4. Command handler extracted to standalone command package
5. System prompts moved to prompt package
6. Session compaction supports before/after hooks
7. HandleMessage now returns error for proper propagation
8.  Hachimi params cache to skip redundant model calls

## Fixes
1. Data races in hachimi and session management
2. CGO SIGABRT crash on context overflow
3. Session compaction running twice
4. Gatekeeper message overflow
5. Handler overwriting tool calls
6. nil pointer access in notifier and daemon
7. Missing error handling in op command
8. Config comma separator parsing
9. Markdown parser issues
10. Telegram HTML escape replacement
11. engine: fix nil pointer dereference crash during message handling 
12. engine: fix pair command leaking into group-side conversations 
13. engine: fix onboard flow not performing as expected 
14. config: add missing selected options in configuration forms

## Chores
1. update descriptions across all tool providers for clarity

## v0.0.1 - Released on Apr 21, 2026

## Added
1. Core Engine
2. Initial project structure and engine architecture
3. Command system with session and provider management 
4. Access control system (ACL) with op/deop commands 
5. Group chat support with message handling 
6. Onboard pairing system for user authentication 
7. Chat context history and persistent storage 
8. Timeout context for request handling
9. Status command and statistics tracking
10. Chat & Providers
11. Brand new Telegram provider with message reactions and notifications
12. Message selection feature in Telegram
13. Support for sending long messages as text files
14. Auto-chat creation for admin users
15. Temporary system prompt support
16. Compact chat system
17. LLM Integration
18. OpenAI-compatible provider support
19. Dynamic provider factory and configuration
20. Model experience enhancement features
21. Token counting and event types
22. Support for OpenRouter, Kimi, Baishan, and other OpenAI-compatible APIs
23. Configuration & CLI
24. Interactive configuration wizard
25. Daemon mode support with manboster start
26. Log viewer functionality
27. manbodev toolchain for plugin development
28. Default home directory configuration read/write
29. Security & Safeguard
30. Hachimi safeguard system for command evaluation
31. Repository interface for data management
32. SQLite database integration with GORM
33. Documentation
34. Added CONTRIBUTING.md with detailed guidelines (Rev 2.2)
35. Added SECURITY.md policy document (Rev 1.1)
36. Added README.md and README.zh.md
37. Apache 2.0 License

## Changed
1. Refactoring
2. Migrated from github.com/chi-net/manboster to github.com/manboster/manboster
3. Optimized file structure and folder organization
4. Refactored LLM provider and model split
5. Updated services instance management
6. Changed JSON to mapstructure for configuration parsing
7. Replaced config factory with chat/provider factory
8. License
9. Changed license from MIT to Apache 2.0

## Fixed
1. Fixed user ID not matching chat ID issue
2. Fixed onboard pair key handling
3. Fixed race condition bugs
4. Fixed permission errors and handler issues
5. Fixed Telegram provider pointer nil issues
6. Fixed provider goroutine and retry mechanisms
7. Fixed background execution errors
8. Fixed configuration folder write permissions