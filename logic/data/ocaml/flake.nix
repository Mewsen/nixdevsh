{
  description = "Ocaml development shell";

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
          packages = with pkgs;
            [
              ocaml
              ocamlformat

              #Nix
              nixfmt-classic
              nil

              # For treesitter
              libgcc

              fd
              ripgrep
            ] ++ (with pkgs.ocamlPackages; dune_3 odoc ocaml-lsp);
        };
      });
}

