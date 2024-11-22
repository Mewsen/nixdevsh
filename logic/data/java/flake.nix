{
  description = "Java development shell";

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
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ self.overlays.default ];
        };
      in {
        overlays.default = final: prev: rec {
          jdk = prev.jdk23;
          maven = prev.maven.override { jdk_headless = jdk; };
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            jdk
            maven
            jdt-langauge-server
            lombok

            #Nix
            nixfmt-classic
            nil

            # For treesitter
            libgcc

            fd
            ripgrep
          ];
          env = {
            JDTLS_JVM_ARGS = "-javaagent:${pkgs.lombok}/share/java/lombok.jar";
          };
        };
      });
}
