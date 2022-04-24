with import <nixpkgs> { };
let
  unstableTarball =
    fetchTarball
      https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz;
  pkgsUnstable = import unstableTarball { };
  mongoGoDriver = buildGoModule {
    src = fetchFromGitHub
      { } + "github.com/mongodb/mongo-go-driver";
  };

in
stdenv.mkDerivation {
  name = "go";
  buildInputs = with pkgsUnstable; [
    go
    gopls
  ];
}
