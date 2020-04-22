{ pkgs ? import <nixpkgs> {} }:
with pkgs;
let
  inherit stdenv buildGoModule;
in
  buildGoModule rec {
   pname = "fooled";
   version = "0.0.1";
  
   goPackagePath = "fooled";

   src = ./.;
   modSha256 = "0qk0kk79xabnjm6121g3brvqvzz76ka7v4xrpfbffvbv7hcg8j8a";

   #buildInputs = [  ];

   #nativeBuildInputs = [ pkg-config ];

   subPackages = [ "." ];

   meta = with stdenv.lib; {
     description = "fooled batsky-go";
     license = licenses.asl20;
     maintainers = with maintainers; [ augu5ste ];
   };
 }
