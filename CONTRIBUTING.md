# Manboster Contributing Guideline

*Rev 1.1; Apr 13, 2026*

## TL;DR

This guideline explains how to contribute to Manboster. You **MUST** read this before writing or refactoring any code.

When contributing to this project (via `git push` or Pull Request), these guidelines apply, and you are held responsible for the code you commit.

If you are a first-time contributor, please look for the `good-first-issue` label in the opened issues.

## 1. About the AI code

1.1 We **DO NOT** accept **ANY** AI Agent's contribution, unless the pull request is opened by the owner and the owner is responsible for the code.

1.2 We accept AI generated codes. But the creator of the pull request is **FULLY RESPONSIBLE** for the code AI generated.

1.3 AI generated code **SHOULD NOT** be used to fix or refactor `good-first-issue`, as these are for human beginner developers.

## 2. About the Versioning

2.1 The Versioning follows the principle of [SemVer](https://semver.org/).

2.2 Version `0.0.1` represents minor bug fixes.

2.3 Version `0.1.0` represents new features.

2.4 Before a `0.1.0` release, Beta and Release Candidate (RC) versions should be released, tagged as `0.x.0-beta` and `0.x.0-rc`.

2.5 It's **NOT ALLOWED** to add new features in RC versions.

2.6 Core contributors may release RC, Stable, and Beta versions for `0.x.0`.

2.7 Version `1.0.0` represents a significant leap in features and architecture. The configuration version shall be upgraded accordingly.

2.8 Before a `1.0.0` release, Alpha, Beta, and Release Candidate (RC) versions should be released, tagged as `x.0.0-alpha`, `x.0.0-beta`, and `x.0.0-rc`.

2.9 Core contributors may release RC and Beta versions for `1.0.0`. Only the owner can release stable `1.0.0` versions.

2.10 Core contributors shall create new git branches for 0.1.0 and 1.0.0 versions (e.g., `dev-v0.x`, `dev-v1.x`) and merge bug fixes from older versions.
## 3. About the commit 

3.1 Commit messages must follow [the Conventional Commits](https://www.conventionalcommits.org/) standard. 

## 4. About the features

4.1 Features should be discussed and documented as [issues](https://github.com/manboster/manboster/issues).

4.2 Please open an issue of feature before you open your pull request.

4.3 New features that introduce breaking changes or modify the core architecture **MUST BE** approved by the owner.

4.4 Other features, such as new chat or LLM providers, should be approved by at least one core contributor.

## 5. About the style

5.1 You should run `go fmt` before you commit your code.

5.2 Please follow standard Go naming conventions: use `MixedCaps` (PascalCase) for exported identifiers and `camelCase` for unexported variables.

5.3 Please use human-friendly function, variable, type names. Meaningless or ambiguous words **ARE NOT** allowed in any identifiers.

5.4 Comments in code **MUST BE** English, not other languages.

## 6. About the testing

6.1 If you contribute to the `engine` package, you **MUST** provide corresponding unit tests for your changes.

6.2 Before you commit, you should run `go test ./...` to ensure all tests are passed.

## Changelogs

This guideline may be updated as needed. Please review it frequently to stay informed of changes.

Rev 1.1, 1.2: Minor Typo fixes, Updated on 2026.4.13

Rev 1: CONTRIBUTING.md created, Updated on 2026.4.13