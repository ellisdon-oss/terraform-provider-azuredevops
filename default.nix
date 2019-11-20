with import <nixpkgs> {};

buildGoModule rec {
  name = "terraform-provider-azuredevops-unstable-${version}";
  version = "2019-02-20";

  goPackagePath = "github.com/ellisdon-oss/terraform-provider-azuredevops";

  src = ./.;

  modSha256 = "1ns5ld61jkjk2xr3c7n32zlyf08k9agf4gca0ravwaaqlqzwa72j";
  
}
