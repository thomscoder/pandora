# Pandora 🏺✨

Pandora is toy browser engine written in Golang and compiled in WASI.
Because why not?

(I just wanted to know how browser render stuff, so I tried to build a browser engine that receives `html` and `css` and outputs a `png`).

Okay. There's no fancy stuff like positioning, z-indexes, flexbox etc... it's very basic.
It was really complex (for my level) and forgive me for errors in advance. Any contribution is welcomed!!

<hr/>


```html
<html lang="en">
<head>
    <title>Pandora</title>
</head>
<body class="something">
    <div id="container">
        <p class="paragraph">Hello</p>
        <p class="another-text">World</p>
    </div>
</body>
</html>
```
```css
html {
    background-color: rgb(255, 224, 193);
    width: 500px;
    height: 500px;
}

body {
    display: block;
    background-color: rgb(99, 35, 0);
    width: 480px;
    height: 480px;
    top: 10px;
    left: 10px;
}

#container {
    background-color: rgb(0, 126, 164);
    width: 280px;
    height: 280px;
    top: 100px;
    left: 100px;
}

.paragraph {
    background-color: rgb(228, 139, 255);
    width: 200px;
    height: 100px;
    top: 10px;
    left: 10px;
}

.another-text {
    background-color: rgb(0, 58, 103);
    width: 140px;
    height: 80px;
    top: 120px;
    left: 10px;
}
```

<img src="image.png" alt="Pandora" border="0">

Anyway you can play around with the `html` structure and the `css` rules (top, left, background-color)

# Why ❓

Why not?

Joke... I've built Pandora because I wanted to learn how browser renders web pages (and it is really complicated). 
<strong>So Pandora was built for LEARNERS.</strong>

> I tried to document each section of the code as much as I could (I'm still doing it) with link to resources I've studied while building this, but if you want to improve it, feel free to open issues and pull requests.

# How it works ❓

Reading this amazing articles about building a <a href="https://limpet.net/mbrubeck/2014/08/08/toy-layout-engine-1.html" target="_blank">Browser engine</a> I decided to try to build one in Go (as I do not know anything about Rust).

- Pandora takes a `.html` and `.css`files
- Builds the `DOM tree` and a very basic `CSSOM`
- Builds a `Render Tree` from the two
- Builds a `Layout Tree` from the Render Tree
- Creates a `Display List`
- Renders, creating a `.png` image

# How to use ❓

To render the html and css in `example/*`

Pandora can be compiled and used as a normal Go program

```bash
pandora
```

```bash
pandora --html example/example.html --css example/example
```

Pandora can also be used as a WASI

With `wasmtime` (specify the directory for the `--dir` flag in which start looking for the files)

```bash
wasmtime --dir . main.wasm -- --html example/example.html --css example/example.css
```

Pandora supports `background-color` `top` `left` `margin` `margin-top` etc...
Follow the Css sample.
Obviously you can contribute whenever you want to make Pandora support more stuff !!

# Requirements ✋

- Go
- TinyGo
- Wasmtime / wasmer or any WASI runtime

Build

Go
```bash
go build
```

WASI

```bash
tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go
```

# Roadmap

1. <h3>Fonts</h3>
2. <h3>Display block / inline / inline - block</h3>

# Contributing

To contribute simply 
- create a branch with the feature or the bug fix
- open pull requests

Pandora is by no means finished there are a lot of things that can be implemented and A LOT of things that can be improved. Any suggestions, pull requests, issues or feedback is greatly welcomed!!!