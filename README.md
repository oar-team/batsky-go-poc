Batsky and Go: Proof of Concept
=========================================

The context of this proof of concept (PoC) is the simulation of infrastructure as proposed by [BatSim](https://github.com/oar-team/batsim). To support native/legacy scheduler we need to substitute real time seen by scheduler by the simulated one provided by Batsim. This POC addresses this substition for **Batsim adaptor** and  **scheduler** (respectively called  **broker** and **fooled** thereafter) into Golang environment. Several ways can be considered (and have been tested), actual retained is to hijacking *time.now()* and *time.runtimeNano()* of the Golang environment, to so we need to recompile Go compiler with its runtime. 

Also we use [Batsky](https://github.com/oar-team/batsky) tools and [Nix](https://nixos.org/nix/). This latter is used to simplify the building process while guaranteeing its reproducibility.

# Dependencies:
- [Batsky](https://github.com/oar-team/batsky)
- [Nix](https://nixos.org/nix/)Glibc Batsky: It's a pached version of glibc. Patch is availabble within  Batsky's source. See below for its building.

# Installation
- Install [Batsky](https://github.com/oar-team/batsky)

```sh
git clone git@github.com:oar-team/batsky
cd batsky
pip install .
```

- Install [Nix](https://nixos.org/nix/)
As user:

```sh
curl -L https://nixos.org/nix/install | sh
```
Don't forget to *source* the indicated file to activate Nix.


- Build broker, it echoes fake time in Batsky's way through zeromq (need for Batsim), compilation use a regular go enviromment
```sh
cd broker
nix-build
```
The latter command produces : *./result/bin/broker* 

- Build Go modified version (take couple of minutes)
```sh
cd ..
nix-build  https://github.com/oar-team/nur-kapack/archive/master.tar.gz --arg pkgs 'import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/release-20.03.tar.gz") {}' -A go_1_14-batsky
```
Ask version of just built go command:
```sh
./result/bin/go version
```

- Build fooled which represents a scheduler with the new Go environment 

```sh
cd fooled
../result/bin/go build fooled.go
```

# Demonstration
- 1) Launch the broker (aka  **Batsim adaptor**)
```sh
./broker/result/bin/broker
```
- 2) Launch batsky (**note:** the creation of */tmp/basky* directory to signal *fooled* that the time hijacking will occur)
```sh
mkdir -p /tmp/batsky; batsky -d -u -c localhost
```

- 3) Launch fooled

With time hijacking

```sh
./fooled/fooled
Hello, world.
Now:  1970-01-05 00:28:11.331251 +0100 CET m=+0.500490001
now from epoch:  343691.331251
```
Note the old time displayed in the output.

To disable the hijacking, kill basky, remove the /tmp/basky and relaunch the fooled and observe the return to present
```sh
killall batsky
rm -rf /tmp/batsky
./fooled/fooled 
Hello, world.
Now:  1970-01-05 00:28:11.331251 +0100 CET m=+0.500490001
now from epoch:  343691.331251
```

# Limitations
- It does not work at the lower level (low level routine of runtime is ignored, only time module is hijacked). As example the scheduler of goroutine is not impacted. 
- Missing timers 

# As it works and the other (failed) ways

## Directories
- **broker**: 
- **fooled**:
- **uds**:

# See also
[Arion Batsky](https://github.com/oar-team/arion-batsky): Tools to do experiments with Batsky and Slurm in containes thanks to [Arion](https://github.com/hercules-ci/arion) and [Nix](https://nixos.org/nix/)


