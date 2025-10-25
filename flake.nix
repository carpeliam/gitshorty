{
  description = "A Shortcut Command-Line tool that works within git branches";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages = {
          default = self.packages.${system}.gitshorty;

          gitshorty = pkgs.buildGoModule {
            pname = "gitshorty";
            version = "0.0.4";

            src = ./.;

            vendorHash = "sha256-/QFFgQinroq0e6dBmRwkGmPP1p8CVtOgQXx9PJXooiM=";

            # The binary is named "sc"
            ldflags = [ "-s" "-w" ];

            # Rename the binary from gitshorty to sc
            postInstall = ''
              mv $out/bin/gitshorty $out/bin/sc
            '';

            meta = with pkgs.lib; {
              description = "A Shortcut Command-Line tool that works within git branches";
              homepage = "https://github.com/carpeliam/gitshorty";
              license = licenses.mit;
              maintainers = [ ];
              mainProgram = "sc";
            };
          };
        };

        # Development shell
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gotools
          ];
        };
      }
    );
}
