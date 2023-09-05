# TODO

### Ideas

- [ ] Add a `--verbose` flag to the `run` command to print the output of the command being run.
- [ ] Add a `--dry-run` flag to the `run` command to print the command that would be run.
- [ ] Add a `--quiet` flag to the `run` command to suppress all output.
- [ ] Add a `serve` flag to run in gui mode to fill the spec file.
- [ ] Add a `--version` flag to print the version of the program.

use git info to append to spec file

- [x] create init command to create a new project
- [x] create clean command to remove all files created by the program
- [ ] create web command to run the program in gui mode
- [ ] create analyze command to run the program in analyze mode

### Ingsights

- [ ] Add a `prepare` command to use the folder structure and criate a spec file with the information about the project and the respective macros.
- [ ] Add a `--verbose` flag to the `run` command to print the output of the command being run.
- [ ] Add a `--dry-run` flag to the `run` command to print the command that would be run.

```
rpmbuild-cli prepare --java --use jdk-11.0.2 --use maven-3.6.0 --use tomcat-9.0.16 --use postgresql-11.2
rpmbuild-cli prepare --native --use gcc-8.2.1 --use make-4.2.1 --use cmake-3.13.4 --use postgresql-11.2
rpmbuild-cli prepare --native --use golang-1.11.5 --use postgresql-11.2
```

### Demo how to use the program