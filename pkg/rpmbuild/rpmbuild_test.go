package rpmbuild

//int main(int argc, char *argv[])
//{
//rpmts ts = NULL;
//enum modes bigMode = MODE_BUILD;
//BTA_t ba = &rpmBTArgs;
//
//const char *pkg = NULL;
//int ec = 0;
//
//poptContext optCon = NULL;
//
//optCon = rpmcliInit(argc, argv, optionsTable);
//
///* Args required only when building, let lone --eval etc through */
//if (ba->buildAmount && poptPeekArg(optCon) == NULL) {
//printUsage(optCon, stderr, 0);
//exit(EXIT_FAILURE);
//}
//
//switch (buildMode) {
//case 'b':	bigMode = MODE_BUILD;		break;
//case 't':	bigMode = MODE_TARBUILD;	break;
//case 'B':	bigMode = MODE_REBUILD;		break;
//case 'C':	bigMode = MODE_RECOMPILE;	break;
//}
//
//if (rpmcliRootDir && rpmcliRootDir[0] != '/') {
//argerror(_("arguments to --root (-r) must begin with a /"));
//}
//
///* rpmbuild runs in verbose mode by default */
//if (rpmlogSetMask(0) < RPMLOG_MASK(RPMLOG_INFO))
//rpmSetVerbosity(RPMLOG_INFO);
//
//if (quiet)
//rpmSetVerbosity(RPMLOG_WARNING);
//
//if (rpmcliPipeOutput && initPipe())
//exit(EXIT_FAILURE);
//
//ts = rpmtsCreate();
//(void) rpmtsSetRootDir(ts, rpmcliRootDir);
//rpmtsSetFlags(ts, rpmtsFlags(ts) | RPMTRANS_FLAG_NOPLUGINS);
//
///* Mind the fallthrough order - it's the reverse of the build process */
//switch (buildChar) {
//case 'a':
//ba->buildAmount |= RPMBUILD_PACKAGESOURCE;
///* fallthrough */
//case 'b':
//ba->buildAmount |= RPMBUILD_PACKAGEBINARY;
//ba->buildAmount |= RPMBUILD_CLEAN;
//if ((buildChar == 'b') && shortCircuit)
//break;
//ba->buildAmount |= RPMBUILD_RMBUILD;
///* fallthrough */
//case 'i':
//ba->buildAmount |= RPMBUILD_INSTALL;
//ba->buildAmount |= RPMBUILD_CHECK;
//if ((buildChar == 'i') && shortCircuit)
//break;
///* fallthrough */
//case 'c':
//ba->buildAmount |= RPMBUILD_BUILD;
//if ((buildChar == 'c') && shortCircuit)
//break;
///* fallthrough */
//case 'f':
//ba->buildAmount |= RPMBUILD_CONF;
//ba->buildAmount |= RPMBUILD_BUILDREQUIRES;
//if (!noDeps) {
//ba->buildAmount |= RPMBUILD_DUMPBUILDREQUIRES;
//ba->buildAmount |= RPMBUILD_CHECKBUILDREQUIRES;
//}
//if ((buildChar == 'f') && shortCircuit)
//break;
///* fallthrough */
//case 'p':
//ba->buildAmount |= RPMBUILD_PREP;
//if (!noDeps) {
//ba->buildAmount |= RPMBUILD_CHECKBUILDREQUIRES;
//}
//break;
//case 'l':
//ba->buildAmount |= RPMBUILD_FILECHECK;
//break;
//case 'r':
///* fallthrough */
//case 'd':
//if (!shortCircuit)
//ba->buildAmount |= RPMBUILD_PREP;
//ba->buildAmount |= RPMBUILD_BUILDREQUIRES;
//ba->buildAmount |= RPMBUILD_DUMPBUILDREQUIRES;
//if (!noDeps)
//ba->buildAmount |= RPMBUILD_CHECKBUILDREQUIRES;
//if (buildChar == 'd')
//break;
///* fallthrough */
//case 's':
//ba->buildAmount |= RPMBUILD_PACKAGESOURCE;
//break;
//}
//ba->buildAmount &= ~(nobuildAmount);
//
//switch (bigMode) {
//case MODE_REBUILD:
//case MODE_RECOMPILE:
//if (bigMode == MODE_REBUILD &&
//buildChar != 'p' &&
//buildChar != 'f' &&
//buildChar != 'c' &&
//buildChar != 'i' &&
//buildChar != 'l') {
//ba->buildAmount |= RPMBUILD_RMSOURCE;
//ba->buildAmount |= RPMBUILD_RMSPEC;
//ba->buildAmount |= RPMBUILD_RMBUILD;
//}
//ba->buildAmount &= ~(nobuildAmount);
//
//while ((pkg = poptGetArg(optCon))) {
//char * specFile = NULL;
//
//ba->cookie = NULL;
//ec = rpmInstallSource(ts, pkg, &specFile, &ba->cookie);
//if (ec == 0) {
//ba->rootdir = rpmcliRootDir;
//ec = build(ts, specFile, ba, rpmcliRcfile);
//}
//ba->cookie = _free(ba->cookie);
//specFile = _free(specFile);
//
//if (ec)
//break;
//}
//break;
//case MODE_BUILD:
//case MODE_TARBUILD:
//
//while ((pkg = poptGetArg(optCon))) {
//ba->rootdir = rpmcliRootDir;
//ba->cookie = NULL;
//ec = build(ts, pkg, ba, rpmcliRcfile);
//if (ec)
//break;
//rpmFreeMacros(NULL);
//(void) rpmReadConfigFiles(rpmcliRcfile, NULL);
//}
//break;
//}
//
//rpmtsFree(ts);
//if (finishPipe())
//ec = EXIT_FAILURE;
//free(ba->buildRootOverride);
//argvFree(build_targets);
//
//rpmcliFini(optCon);
//
//return RETVAL(ec);
//}
