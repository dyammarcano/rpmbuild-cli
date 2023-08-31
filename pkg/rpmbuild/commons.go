package rpmbuild

var (
	_fsm_debug    int
	specFlags     RpmSpecFlags
	noDeps        int
	shortCircuit  int
	buildMode     byte
	buildChar     byte
	nobuildAmount RpmBuildFlags
	buildTargets  []string
	buildInPlace  int
	quiet         bool
	rpmBTArgs     RpmBuildArguments
)

type (
	RpmFlags           uint32
	RpmBuildPkgFlags   RpmFlags
	RpmBuildFlags      RpmFlags
	RpmSpecFlags       RpmFlags
	errmsg_t           string
	Header             any
	HeaderIterator     any
	rpmTagT            int32
	rpm_tagtype_t      uint32
	rpm_count_t        uint32
	rpmTagVal          rpmTagT
	rpmDbiTagVal       rpmTagT
	rpm_data_t         any
	rpm_constdata_t    any
	rpmtd              any
	rpm_color_t        uint32
	rpm_flag_t         uint32
	rpm_tid_t          uint32
	rpmFlags           uint32
	rpm_off_t          uint32
	rpm_loff_t         uint64
	rpm_time_t         uint32
	rpm_mode_t         uint16
	rpm_rdev_t         uint16
	rpm_dev_t          uint32
	rpm_ino_t          uint32
	rpmts              any
	rpmte              any
	rpmds              any
	rpmfi              any
	rpmfiles           any
	rpmdb              any
	rpmdbMatchIterator any
	rpmtsi             any
	rpmps              any
	rpmtxn             any
	rpmver             any
	rpmdbIndexIterator any
	fnpyKey            any
	rpmCallbackData    any
	rpmPubkey          any
	rpmKeyring         any
	rpmsid             uint32
	rpmstrPool         any
	rpmPlugin          any
	rpmPlugins         any
	rpmgi              any
	rpmSpec            any
	rpmRelocation      struct{}
	FdT                any
	rpmRC              int
	rpmBuildFlags      int
	rpmBuildPkgFlags   int
	rpmSpecFlags       int
	BtaT               *rpmBuildArguments
	rpmMacroItemLocal  int

	rpmBuildArguments struct {
		pkgFlags          rpmBuildPkgFlags
		buildAmount       rpmBuildFlags
		buildRootOverride string
		cookie            string
		rootdir           string
	}

	RpmBuildArguments struct {
		PkgFlags          RpmBuildPkgFlags
		BuildAmount       RpmBuildFlags
		BuildRootOverride string
		Cookie            string
		Rootdir           string
	}

	poptOption struct {
		longName   string
		shortName  byte
		argInfo    uint32
		arg        interface{}
		val        int
		descrip    string
		argDescrip string
	}

	poptAlias struct {
		longName  string
		shortName byte
		argc      int
		argv      []string
	}

	poptItem struct {
		option poptOption
		argc   int
		argv   []string
	}

	poptContext struct {
		name        string
		argc        int
		argv        []string
		options     []poptOption
		flags       uint32
		optionStack []poptOption
		arg         string
		argi        int
		leftovers   []string
		nextLeft    int
		leftarg     string
	}
)

const (
	POPT_ARG_NONE        = 0
	POPT_ARG_STRING      = 1
	POPT_ARG_INT         = 2
	POPT_ARG_LONG        = 3
	PoptArgIncludeTable  = 4
	POPT_ARG_CALLBACK    = 5
	POPT_ARG_INTL_DOMAIN = 6
	PoptArgVal           = 7
	POPT_ARG_FLOAT       = 8
	POPT_ARG_DOUBLE      = 9
	POPT_ARG_LONGLONG    = 10

	POPT_ARG_MAINCALL = 16 + 11
	POPT_ARG_ARGV     = 12
	POPT_ARG_SHORT    = 13
	POPT_ARG_BITSET   = 16 + 14

	POPT_ARGFLAG_ONEDASH    = 0x80000000
	PoptArgflagDocHidden    = 0x40000000
	POPT_ARGFLAG_STRIP      = 0x20000000
	POPT_ARGFLAG_OPTIONAL   = 0x10000000
	PoptArgflagOr           = 0x08000000
	POPT_ARGFLAG_NOR        = 0x09000000
	PoptArgflagAnd          = 0x04000000
	PoptArgflagNand         = 0x05000000
	PoptArgflagXor          = 0x02000000
	POPT_ARGFLAG_NOT        = 0x01000000
	POPT_ARGFLAG_LOGICALOPS = PoptArgflagOr | PoptArgflagAnd | PoptArgflagXor

	POPT_BIT_SET = PoptArgVal | PoptArgflagOr
	POPT_BIT_CLR = PoptArgVal | PoptArgflagNand

	POPT_ARGFLAG_SHOW_DEFAULT = 0x00800000
	POPT_ARGFLAG_RANDOM       = 0x00400000
	POPT_ARGFLAG_TOGGLE       = 0x00200000

	POPT_CBFLAG_PRE        = 0x80000000
	POPT_CBFLAG_POST       = 0x40000000
	POPT_CBFLAG_INC_DATA   = 0x20000000
	POPT_CBFLAG_SKIPOPTION = 0x10000000
	POPT_CBFLAG_CONTINUE   = 0x08000000

	POPT_ERROR_NOARG        = -10
	POPT_ERROR_BADOPT       = -11
	POPT_ERROR_UNWANTEDARG  = -12
	POPT_ERROR_OPTSTOODEEP  = -13
	POPT_ERROR_BADQUOTE     = -15
	POPT_ERROR_ERRNO        = -16
	POPT_ERROR_BADNUMBER    = -17
	POPT_ERROR_OVERFLOW     = -18
	POPT_ERROR_BADOPERATION = -19
	POPT_ERROR_NULLARG      = -20
	POPT_ERROR_MALLOC       = -21
	POPT_ERROR_BADCONFIG    = -22

	MODE_BUILD     = 1 << 4
	MODE_REBUILD   = 1 << 5
	MODE_RECOMPILE = 1 << 8
	MODE_TARBUILD  = 1 << 11

	RPMRC_OK         rpmRC = 0
	RPMRC_NOTFOUND   rpmRC = 1
	RPMRC_FAIL       rpmRC = 2
	RPMRC_NOTTRUSTED rpmRC = 3
	RPMRC_NOKEY      rpmRC = 4

	RPMBUILD_NONE               rpmBuildFlags = 0
	RpmbuildPrep                rpmBuildFlags = 1 << 0
	RpmbuildBuild               rpmBuildFlags = 1 << 1
	RpmbuildInstall             rpmBuildFlags = 1 << 2
	RPMBUILD_CHECK              rpmBuildFlags = 1 << 3
	RPMBUILD_CLEAN              rpmBuildFlags = 1 << 4
	RPMBUILD_FILECHECK          rpmBuildFlags = 1 << 5
	RPMBUILD_PACKAGESOURCE      rpmBuildFlags = 1 << 6
	RpmbuildPackagebinary       rpmBuildFlags = 1 << 7
	RpmbuildRmsource            rpmBuildFlags = 1 << 8
	RPMBUILD_RMBUILD            rpmBuildFlags = 1 << 9
	RPMBUILD_STRINGBUF          rpmBuildFlags = 1 << 10
	RpmbuildRmspec              rpmBuildFlags = 1 << 11
	RPMBUILD_FILE_FILE          rpmBuildFlags = 1 << 16
	RPMBUILD_FILE_LIST          rpmBuildFlags = 1 << 17
	RPMBUILD_POLICY             rpmBuildFlags = 1 << 18
	RPMBUILD_CHECKBUILDREQUIRES rpmBuildFlags = 1 << 19
	RPMBUILD_BUILDREQUIRES      rpmBuildFlags = 1 << 20
	RPMBUILD_DUMPBUILDREQUIRES  rpmBuildFlags = 1 << 21
	RpmbuildNobuild             rpmBuildFlags = 1 << 31

	RPMBUILD_PKG_NONE      rpmBuildPkgFlags  = 0
	RpmbuildPkgNodirtokens rpmBuildPkgFlags  = 1 << 0
	BUFSIZ                                   = 8192
	RmilTarball            rpmMacroItemLocal = 1

	PoptNolang         = -1012
	PoptRmsource       = -1013
	PoptRmbuild        = -1014
	PoptBuildroot      = -1015
	PoptTargetplatform = -1016
	PoptNobuild        = -1017
	PoptRmspec         = -1019
	PoptNodirtokens    = -1020
	PoptBuildinplace   = -1021

	PoptRebuild   = 0x4262 /* Bb */
	PoptRecompile = 0x4369 /* Ci */
	PoptBa        = 0x6261
	PoptBb        = 0x6262
	PoptBc        = 0x6263
	POPT_BD       = 0x6264
	POPT_BF       = 0x6266
	POPT_BI       = 0x6269
	POPT_BL       = 0x626c
	POPT_BP       = 0x6270
	POPT_BS       = 0x6273
	POPT_BR       = 0x6272
	POPT_RA       = 0x4261
	POPT_RB       = 0x4262
	POPT_RC       = 0x4263
	POPT_RD       = 0x4264
	POPT_RF       = 0x4266
	POPT_RI       = 0x4269
	POPT_RL       = 0x426c
	POPT_RP       = 0x4270
	POPT_RS       = 0x4273
	POPT_RR       = 0x4272
	POPT_TA       = 0x7461
	POPT_TB       = 0x7462
	POPT_TC       = 0x7463
	POPT_TD       = 0x7464
	POPT_TF       = 0x7466
	POPT_TI       = 0x7469
	POPT_TL       = 0x746c
	POPT_TP       = 0x7470
	POPT_TS       = 0x7473
	POPT_TR       = 0x7472
)
