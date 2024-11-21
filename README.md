<h1 align="center">nixdevsh</h1>

<p align="center">
  <a href="https://github.com/mewsen/nixdevsh/blob/master/license">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

<p align="center">
  <b>A Flake based nix-shell generator</p>

![nixdevsh demo](demo.gif)

## Usage
```sh
nix run github:mewsen/nixdevsh/master
direnv allow # if you're using direnv https://github.com/nix-community/nix-direnv
nix develop -c $SHELL # if you're not using direnv
```

You can also add the input to your config and then use it as a package
```nix
{ pkgs, inputs, ... }: {
  environment.systemPackages = with pkgs; [
    inputs.nixdevsh.packages.x86_64-linux.default
  ]
}
```

## Building
```sh
nix build
```
