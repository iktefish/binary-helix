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
  goOsStat = buildGoModule {
    src = fetchFromGitHub
      { } + "github.com/mackerelio/go-osstat";
  };
  googleUuid = buildGoModule {
    src = fetchFromGitHub
      { } + "github.com/google/uuid";
  };

in
stdenv.mkDerivation {
  name = "go";
  buildInputs = with pkgsUnstable; [
    go
    gopls
    graphviz
  ];
}
