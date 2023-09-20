PORTNAME=	abseil
PORTVERSION=	20230125.3
CATEGORIES=	devel

MAINTAINER=	danfe@FreeBSD.org
COMMENT=	Abseil Common Libraries (C++)
WWW=		https://abseil.io/

LICENSE=	APACHE20

USES=		cmake compiler:c++17-lang
USE_GITHUB=	yes
GH_PROJECT=	${PORTNAME}-cpp

CMAKE_ARGS=	-DABSL_PROPAGATE_CXX_STD=ON \
		-DCMAKE_CXX_STANDARD=17 # force the same highest C++ standard because of ABI differences: see https://github.com/abseil/abseil-cpp/issues/819

_GET_SOVER=	${GREP} SOVERSION ${WRKSRC}/CMake/AbseilHelpers.cmake \
		| ${CUT} -d\" -f2
PLIST_SUB+=	LIBSUFX=${_GET_SOVER:sh}

pre-configure:
	@${MKDIR} ${WRKDIR}/.build_static ${WRKDIR}/.build_shared

do-configure:
	cd ${WRKDIR}/.build_static && ${SETENV} ${CONFIGURE_ENV} ${CMAKE_BIN} ${CMAKE_ARGS} -DBUILD_SHARED_LIBS=OFF ${WRKSRC}
	cd ${WRKDIR}/.build_shared && ${SETENV} ${CONFIGURE_ENV} ${CMAKE_BIN} ${CMAKE_ARGS} -DBUILD_SHARED_LIBS=ON ${WRKSRC}

do-build:
	cd ${WRKDIR}/.build_static && ${DO_MAKE_BUILD}
	cd ${WRKDIR}/.build_shared && ${DO_MAKE_BUILD}

do-install:
	cd ${WRKDIR}/.build_static && ${SETENV} ${MAKE_ENV} ${FAKEROOT} ${MAKE_CMD} ${MAKE_FLAGS} ${MAKEFILE} ${MAKE_ARGS} ${INSTALL_TARGET}
	cd ${WRKDIR}/.build_shared && ${SETENV} ${MAKE_ENV} ${FAKEROOT} ${MAKE_CMD} ${MAKE_FLAGS} ${MAKEFILE} ${MAKE_ARGS} ${INSTALL_TARGET}

.include <bsd.port.mk>
