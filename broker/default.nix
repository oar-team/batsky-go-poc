{ pkgs ? import <nixpkgs> {} }:
with pkgs;
let
  inherit stdenv buildGoModule zeromq pkg-config;
in
  buildGoModule rec {
   pname = "broker";
   version = "0.0.1";
  
   goPackagePath = "broker";

   src = ./.;
   modSha256 = "02fvav564b0vmnsrkw7xvqg1afwhp92h0glhn7nvrgdidk8r4534";

   buildInputs = [ zeromq ];

   nativeBuildInputs = [ pkg-config ];

   subPackages = [ "." ];

   meta = with stdenv.lib; {
     description = "Broker poc";
     license = licenses.asl20;
     maintainers = with maintainers; [ augu5ste ];
   };
 }
