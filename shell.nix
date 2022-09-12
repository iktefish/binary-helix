with import <nixpkgs> {}; let
  mongoGoDriver = buildGoModule {
    src =
      fetchFromGitHub
      {}
      + "github.com/mongodb/mongo-go-driver";
  };
  goOsStat = buildGoModule {
    src =
      fetchFromGitHub
      {}
      + "github.com/mackerelio/go-osstat";
  };
  googleUuid = buildGoModule {
    src =
      fetchFromGitHub
      {}
      + "github.com/google/uuid";
  };
in
  stdenv.mkDerivation {
    name = "binary-helix";
    buildInputs = with pkgs; [
      graphviz
    ];
  }
