with import <nixpkgs> {};

buildGoPackage rec {
  name = "terraform-provider-azuredevops-unstable-${version}";
  version = "2019-02-20";

  goPackagePath = "github.com/ellisdon/terraform-provider-azuredevops";

  src = ./.;

  goDeps = ./deps.nix;
}
