PORTNAME=	crowdsec-blocklist-mirror
DISTVERSIONPREFIX=	v
DISTVERSION=	0.0.5
CATEGORIES=	security

MAINTAINER=	marco@crowdsec.net
COMMENT=	CrowdSec Blocklist Mirror
WWW=		https://github.com/crowdsecurity/cs-blocklist-mirror

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
_BUILD_DATE=	$$(date -u "+%F_%T")
USE_RC_SUBR=	crowdsec_mirror
_COMMIT=	f33db07e

GO_MODULE=	github.com/crowdsecurity/cs-blocklist-mirror
GO_TARGET=	${PORTNAME}:./${PORTNAME}
GO_BUILDFLAGS=	-trimpath -tags netgo \
		-ldflags="\
		-a -s -extldflags '-static' \
		-X github.com/crowdsecurity/go-cs-lib/version.Version=${DISTVERSIONPREFIX}${PKGVERSION}-freebsd \
		-X github.com/crowdsecurity/go-cs-lib/version.BuildDate=${_BUILD_DATE} \
		-X github.com/crowdsecurity/go-cs-lib/version.Tag=${_COMMIT}"

CGO_ENABLED=	0

SUB_FILES=	pkg-deinstall \
		pkg-install \
		pkg-message

ETCDIR=		${PREFIX}/etc/crowdsec/bouncers

do-install:
	#
	# Binaries
	#

	${INSTALL_PROGRAM} ${WRKDIR}/bin/${PORTNAME} \
		${STAGEDIR}${PREFIX}/bin/${PORTNAME}

	#
	# Configuration
	#

	@${MKDIR} ${STAGEDIR}${ETCDIR}
	${INSTALL_DATA} ${WRKSRC}/config/${PORTNAME}.yaml \
		${STAGEDIR}${ETCDIR}/${PORTNAME}.yaml.sample

.include <bsd.port.mk>
