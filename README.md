# 🐱🦞Manboster: Your Personal Manbo Lobster!

[简体中文](README.zh.md)

> Tips: Manboster = Manbo + Lobster

Inspired by IronClaw and OpenClaw, we've built a lobster more securely!

## Meet Your Personal Lobster

At the beginning, it is a personal AI assistant which is able to chat with you with ease. However, equipped with wasm-based plugins, you can use it to do anything you want with security guaranteed!

Since this project is work in progress and only shows as a MVP now, there are a few chat options(Telegram and Feishu/Lark as a second supported provider in next versions) and LLM options(openrouter, kimi, DeepSeek and openai-compatible APIs)

We will sincerely appreciate if you contribute your own codes!

BTW, Manboster's code is mainly written by human, with AIs written in helper functions and boilerplate.

## Why Manboster?

1. Out of the box, only an executable file, based on Golang.
2. Low memory occupation, fast, multithreaded, non-blocking in chats thanks to Golang.
3. Introducing Hachimi, a guard model running normally locally, will evaluate LLM's tool call if you need. If Hachimi thinks some action is unsafe, it will stop and let you decide. 
4. Default zero-trust design, gatekeeper system and TTL settings to handle all tool calls. 
5. Pluggable built-in tools, you can enable or disable it using `manboster config`.
6. A built-in web search tool with headless browsers using `go-rod`. Use it to search for the Internet or give your search API's keys
7. [Experimental] Compatibility with your old OpenClaw skills, just type `manboster skills install SKILLS.md` or `manboster skills install SKILLS.zip` to install!
8. [Experimental] A built-in vault tool helps you store your sensitive data using industry best practices while balancing your experience. LLM NEVER has access to you credentials.
9. [Experimental] A built-in script runner sandbox tool enables JavaScript or Python scripts can run in wasm sandbox. 
10. [Work in Progress] Pluggable RAG memory system and mem0 theory adaption.
11. [Work in Progress] MCP(Model Context Protocol) Compatibility, just add them and use them as native tool calls in Manboster.
12. [Work in Progress] Simulate UI/Input interactions, screenshot, there are a plenty of Built-in SDKs to use for building your plugins.
13. [Work in Progress] Plugins based on wasm & extism which is lightweight and prevents malicious plugins from breaking your machine.
14. [Work in Progress] MamboHub, a distribution center enables you to use and download skills and plugins with ease. Also, we will compat ClawHub and more skill hubs.
15. [Work in Progress] Install plugins using `.manboplugin` files, you can even develop and write your plugins using `manbodev` helper!

## QuickStart

1. Download binary files built in releases and open it via double-click or terminal shell `./manboster`.
2. You're required to configure your manboster when it runs in the first time, just configure it.
3. Enjoy yourselves!

If you downloaded the Go programming language development kit? Just type `go install github.com/manboster/manboster/cmd/manboster@latest` to install!

For more details, please run `manboster help` in your terminal or [read this documentation](https://manboster.dev/docs/quickstart.html)!

**Notes for the Daemon**: Simply double-click would not start the daemon, if you want to create and start the daemon, please run `manboster start`.
## Make a contribution

Manboster is now open to the community and we are looking forward to your contributions! Read [CONTRIBUTING.md](./CONTRIBUTING.md) to start!

## Troubleshoot

If you are facing troubles, read [our documentation](https://manboster.dev/docs/) first.

Create a new issue after trying ways that documentation says or what you've faced is not recorded in documentation.

Before creating a new issue, search for the problem. If there is none, create it after reading [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html).

## License

This application's license is Apache 2.0.

## Thanking

Please read [THANKING.md](./THANKING.md)
