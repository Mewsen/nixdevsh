{
  description = "LaTeX development shell";

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
            texlab
            texlive.combined.scheme-full
            tectonic

            # LSP
            ltex-ls

            # Nix
            nixfmt-classic
            nil

            # For treesitter
            libgcc

            fd
            ripgrep
          ];
        };
        packages.default = pkgs.stdenv.mkDerivation rec {
          name = "latex-document";
          version = "0.0.1";
          pwd = ./.;
          src = ./.;
          buildInputs = [ pkgs.texlive.combined.scheme-full ];

          buildPhase = ''
            export HOME=.
            latexmk -jobname=${name} -pdf -interaction=nonstopmode -output-directory=$out document.tex
          '';
        };
      });

}
