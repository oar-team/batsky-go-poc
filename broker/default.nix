{ pkgs ? import <nixpkgs> {} }:
with pkgs;
let
  inherit stdenv buildGoModule zeromq pkg-config lib;
in
  buildGoModule rec {
   pname = "poc-broker";
   version = "0.0.1";
  
   goPackagePath = "poc-broker";

   src = ./.;
   modSha256 = "02fvav564b0vmnsrkw7xvqg1afwhp92h0glhn7nvrgdidk8r4534";

   #buildFlags = [ "-tags" "extended" ];

   buildInputs = [ zeromq ];

   nativeBuildInputs = [ pkg-config ];

   subPackages = [ "." ];

   meta = with stdenv.lib; {
     description = "poc for batKube.";
     license = licenses.asl20;
     maintainers = with maintainers; [ augu5ste ];
   };
 }
