{ pkgs ? import <nixpkgs> {} }:
with pkgs;
let
  inherit stdenv buildGoModule lib;
  nur = import (builtins.fetchTarball "https://github.com/nix-community/NUR/archive/master.tar.gz") {inherit pkgs;};
in
  buildGoModule rec {
  pname = "poc-fooled";
  version = "0.0.1";
  
  goPackagePath = "fooled";
  
  src = ./.;
  
  modSha256 = "0sjjj9z1dhilhpc8pq4154czrb79z9cm044jvn75kxcjv6v5l2m5";
  
  buildInputs = [ go];
  
  nativeBuildInputs = [ go ];

  buildFlags = [ "-ldflags" "-linkmode=external"];
  #buildFlags = [ "-ldflags" "'-linkmode=external -L ${nur.repos.kapack.glibc-batsky}'" ];
  
  subPackages = [ "." ];
  
  #preBuild = ''export CGO_LDFLAGS="-L ${nur.repos.kapack.glibc-batsky}/lib/libc_nonshared.a"'';
  preBuild = ''export CGO_LDFLAGS="-L ${nur.repos.kapack.glibc-batsky}"'';

  dontFixup = true;
  dontStrip = true;
  
  meta = with stdenv.lib; {
    description = "poc fooled for BatKube dev.";
    license = licenses.asl20;
    maintainers = with maintainers; [ augu5ste ];
  };
}
