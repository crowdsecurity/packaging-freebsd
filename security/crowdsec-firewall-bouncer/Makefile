PORTNAME=	crowdsec-firewall-bouncer
DISTVERSIONPREFIX=	v
DISTVERSION=	0.0.33
CATEGORIES=	security

MAINTAINER=	marco@crowdsec.net
COMMENT=	CrowdSec bouncer written in golang for firewalls
WWW=		https://github.com/crowdsecurity/cs-firewall-bouncer

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
_BUILD_DATE=	$$(date -u "+%F_%T")
USE_RC_SUBR=	crowdsec_firewall
_COMMIT=	cb8b3e3c

GO_MODULE=	github.com/crowdsecurity/cs-firewall-bouncer
GO_TARGET=	${PORTNAME}:./${PORTNAME}
GO_BUILDFLAGS=	-trimpath -tags netgo \
		-ldflags="\
		-a -s \
		-X github.com/crowdsecurity/go-cs-lib/version.Version=${DISTVERSIONPREFIX}${PKGVERSION}-freebsd \
		-X github.com/crowdsecurity/go-cs-lib/version.BuildDate=${_BUILD_DATE} \
		-X github.com/crowdsecurity/go-cs-lib/version.Tag=${_COMMIT}"

# to allow WITH_PIE with a non PIC libc
CGO_ENABLED=	1

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
