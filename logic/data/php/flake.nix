{
  description = "PHP development shell";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.0.tar.gz";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ] (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            # PHP
            php
            phpactor
            phpPackages.composer

            # LSP's
            typescript-language-server
            vtsls
            vscode-langservers-extracted
            tailwindcss-language-server

            #Nix
            nixfmt-classic
            nil

            # For treesitter
            libgcc

            fd
            ripgrep
          ];
        };
      });
}
