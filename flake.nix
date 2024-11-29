{
  description = "nixdevsh";

  inputs = {
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.0.tar.gz";
    flake-utils.url = "github:numtide/flake-utils";

    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gomod2nix }:
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ] (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.go_1_23;
              buildGoModule = prev.buildGo123Module;
            })
            (self: super: {
              buildGoApplication = super.callPackage ./builder { };
            })
            gomod2nix.overlays.default
          ];
        };

        p = pkgs.buildGoApplication {
          pname = "nixdevsh";
          version = "0.0.5";
          src = ./.;
          pwd = ./.;
          modules = ./gomod2nix.toml;

          buildInputs = with pkgs; [ go ];
        };

      in {
        packages = {
          default = p;
          installPhase = ''
            mkdir -p $out/bin
            cp $src/bin/$pname $out/bin/$path'';

        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
            gomod2nix.packages.${system}.default
          ];
        };
      });

}

