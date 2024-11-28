{
  description = "Go package";

  inputs = {
    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.0.tar.gz";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };

  outputs = { self, nixpkgs, flake-utils, gomod2nix }:
    (flake-utils.lib.eachDefaultSystem (system:
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

        goEnv = pkgs.mkGoEnv { pwd = ./.; };

        p = pkgs.buildGoApplication {
          pname = "test";
          version = "0.1";
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
          packages = [
            (with pkgs; [ go gopls gotools libgcc fd ripgrep ])
            goEnv
            gomod2nix.packages.${system}.default
          ];
        };
      }));
}
