# 🐱🦞Manboster: Your Personal Manbo Lobster!

[简体中文](README.zh.md)

> Tips: Manboster = Manbo + Lobster

Inspired by IronClaw and OpenClaw, we've built a lobster more secure than others!

At the beginning, it is a personal AI assistant which is able to chat with you with ease. However, equipped with wasm-based plugins, you can use it to do anything you want with security guaranteed!

Since this project is working in progress and only shows as a MVP now, there is a few chat options(Telegram) and LLM options(openrouter, kimi, baishan and openai-compatible APIs.) We will sincerely appreciate if you contribute your own codes!

## Why Manboster?

1. Out of the box, only an executable file, based on Golang.
2. Low memory occupation, fast, multithreaded, non-blocking in chats thanks to Golang.
3. Introducing Hachimi, a guard model running normally locally, will evaluate LLM's tool call if you need.
4. Default zero-trust design, gatekeeper system and TTL settings to handle all tool calls. 
5. Pluggable build-in tools, you can enable or disable it using `manboster config`.
6. A build-in web search tool with headless browsers using `go-rod`. Use it to search for the Internet or give your search API's keys
7. [Experimental] Compatability with your old OpenClaw skills, just type `manboster skills install SKILLS.md` or `manboster skills install SKILLS.zip` to install!
8. [Experimental] A build-in vault tool helps you store your sensitive data using industry best practices while balancing your experience.
9. [Experimental] A build-in script runner sandbox tool enables JavaScript or Python scripts can run in wasm sandbox. 
10. [Work in Progress] Pluggable RAG memory system and mem0 theory adaption.
11. [Work in Progress] Simulate UI/Input interactions, screenshot, there are a plenty of Built-in SDKs to use for building your plugins.
12. [Work in Progress] Plugins based on wasm & extism which is lightweight and prevents malicious plugins from breaking your machine.
13. [Work in Progress] MamboHub, a distribute center enables you to use and download skills and plugins with ease
14. [Work in Progress] Install skills using `.manboskill` files or plugins using `.manboplugin` files, you can even develop and write your skills and plugins using `manbodev` helper!

## QuickStart

1. Download binary files built in releases and open it via double-click or terminal shell `./manboster`.
2. You're required to configure your manboster when it runs in the first time, just configure it.
3. Enjoy yourselves!

For more details, please run `manboster help` in your terminal!

**Notes for the Daemon**: Simply double-click would not start the daemon, if you want to create and start the daemon, please run `manboster start`.

## Make a contribution

We are looking forward you to contribute this repository! But Before you start, please carefully read [CONTRIBUTING.md](./CONTRIBUTING.md) then write code, open PR. 

## Troubleshoot

If you are facing troubles, please read [our documentation](https://manboster.dev/docs/) first.

You can create a new issue after trying ways that documentation says or what you've faced is not recorded in documentation.

Before you create a new issue, you can search for the problem, if there is none, you can create it after reading [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html).

## License

This application's license is Apache 2.0.
