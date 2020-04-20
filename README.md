POC Batsky GO
==================

The context of this proof of concept (POC) if the simulation of infrastructure as proposed by [BatSim](https://github.com/oar-team/batsim). To support native/legacy scheduler we need to substitute real time seen by scheduler by the simulated one provided by Batsim. This POC addresses is substition for **Batsim adaptor** and  **scheduler** (respectively called  **broker** and **fooled** thereafter) into Golang environment. Also we use [Batsky](https://github.com/oar-team/batsky) tools and [Nix](https://nixos.org/nix/).

Note: 
One way do this is to patch the venerable Glibc library


# Dependencies:
- [Batsky](https://github.com/oar-team/batsky)
- Glibc Batsky: It's a pached version of glibc. Patch is availabble within  Batsky's source. See below for its building.

# Installation
- Install [Batsky](https://github.com/oar-team/batsky)

```sh
git clone git@github.com:oar-team/batsky
cd batsky
pip install .
```

- Build Glibc-Batsky
cd fooled
nix-build https://github.com/oar-team/nur-kapack/archive/master.tar.gz -A glibc-batsky -o glibc-batsky

go build --ldflags '-linkmode external -L ./glibc-batsky



go build --ldflags '-linkmode external -L /path/to/another_glibc/'
To stop containers
```sh
cd fooled
go build --ldflags '-linkmode external -L ./glibc-batsky
```

# Test
## Broker only

## Broker, Bastky and Fooled
# See also
[Arion Batsky](https://github.com/oar-team/arion-batsky): Tools to do experiments with Batsky and Slurm in containes thanks to [Arion](https://github.com/hercules-ci/arion) and [Nix](https://nixos.org/nix/)


