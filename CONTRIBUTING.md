# Manboster Contributing Guideline

*Rev 4; May 24, 2026*

## TL;DR

Welcome to contribute Manboster! Manboster is an AI agent focused on security and personal conditions and we are looking forward you to join us!

When contributing to this project (via `git push` or Pull Request), these guidelines apply.

If you are a first-time contributor, please look for the `good-first-issue` label in the opened issues.

## QuickStart

1. Fork and clone the repository: `git clone git@github.com:/yourname/manboster.git`
2. Run `go get`
3. Use `make run` to start Manboster!
4. Edit code and push them onto GitHub
5. Open a PR and wait for merge!

## 1. About the AI code

We don't care that whether Claude or ChatGPT helped you code or not, however, Manboster features security. So you need to explain your logic and intention in your pull request as the code is yours, AI can't take responsibility to security issues.

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

## 7. About the Contribution Roles and Ethics

We are using flat structure to manage now so it's the problem of time to merge them when you opened high-quality pull requests.

If you continuously contribute high-quality codes, I'm happy to invite you to join our develop team and be the maintainer of Manboster!

## 8. About the Issues and Discussions

8.1 Issues is the place where people discuss features and bug fixes. 

8.2 Discussions is our community space for casual chat, questions, and networking. All participants must adhere to [GitHub's Terms of Service](https://docs.github.com/en/site-policy/github-terms/github-terms-of-service).

8.3 We **DO NOT** welcome spam, including "check-ins" (such as sending "留名" "打卡" and more) or meaningless comments in Issues. Such content will be closed or deleted **WITHOUT NOTICE**. While we are more lenient in Discussions, please keep the content constructive.

8.4 We welcome all security researchers to find threats and vulnerabilities in Manboster. If you are a security researcher and have found a vulnerability in Manboster, please **DO NOT** open a public issue. Use the GitHub Security Advisory feature or send a detailed report to `security@manboster.dev`. Once verified, we will coordinate a fix as a priority.

## Changelogs

This guideline may be updated as needed. Please review it frequently to stay informed of changes.

Rev 4: Simplify Contributing Guideline in order to boost the community, Updated on 2026.5.24

Rev 3.1: Add limitation to the owner, Updated on 2026.4.23

Rev 3: Update morals and updates of this guideline, Updated on 2026.4.23

Rev 2.2: Fix typo, Updated on 2026.4.13

Rev 2.1: Fix ambiguous words may occur misunderstanding, Updated on 2026.4.13

Rev 2: Add contribution roles and issues, discussions, Updated on 2026.4.13

Rev 1.1, 1.2: Minor Typo fixes, Updated on 2026.4.13

Rev 1: CONTRIBUTING.md created, Updated on 2026.4.13