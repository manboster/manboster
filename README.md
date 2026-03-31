# 🐱🦞Manboster: Your Personal Manbo Lobster!

[简体中文](README.md)

> Tips: Manboster = Manbo + Lobster

Inspired by IronClaw and OpenClaw, we've built a lobster more secure than others!

At the beginning, it is a personal AI assistant which is able to chat with you with ease. However, equipped with wasm-based plugins, you can use it to do anything you want with security guaranteed!

## Features

1. Out of the box, only an executable file, based on Golang.
2. Fast, multithreaded, non-blocking in chats.
3. When the LLM(either using Markdown skills or wasm plugins) wants to do anything on your machine, small LLM(hachimi) in your machine will evaluate and score it first. If the score is high, it will send you a message, let you decide.
4. While maintaining compatability with OpenClaw markdown-based skills, we introduced plugins based on wasm & extism, which is lightweight and prevents malicious plugins from breaking your machine.
5. Mock touching, screenshot, web search, execute commands and more, there are a plenty of extism-based Built-in SDKs to use for building your plugins.
6. When it comes to web search, you can choose API Key or headless browsers, which the latter will save you from expenses!
7. MamboHub enables you to use and download skills and plugins with ease, install skills using `.manboskill` files or plugins using `.manboplugin` files, and you can even develop and write your skills and plugins!

## QuickStart

1. Download binary files built in releases and open it.
2. You're required to configure your manboster when it runs in the first time, just configure it.
3. Enjoy yourselves!

For more details, you can run `manboster help` in your terminal!

**Notes for the Daemon**: Simply double-click would not start the daemon, if you want to create and start the daemon, please run `manboster start`.

## License

This application's license is MIT.
