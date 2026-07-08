PORTNAME= 	crowdsec-haproxy-spoa-bouncer
DISTVERSION=    0.3.0
DISTVERSIONPREFIX= v
CATEGORIES= 	security

MAINTAINER= 	manuel@crowdsec.net
COMMENT=       	crowdsec haproxy spoa bouncer

LICENSE= 	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
_BUILD_DATE=	$$(date -u "+%F_%T")
USE_RC_SUBR=	crowdsec_spoa_bouncer
_COMMIT=	39bfe92

GO_MODULE=	github.com/crowdsecurity/cs-haproxy-spoa-bouncer
GO_TARGET=	.:crowdsec-spoa-bouncer
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
	${INSTALL_PROGRAM} ${WRKDIR}/bin/crowdsec-spoa-bouncer \
		${STAGEDIR}${PREFIX}/bin/crowdsec-spoa-bouncer

	@${MKDIR} ${STAGEDIR}${ETCDIR}
	${INSTALL_DATA} ${WRKSRC}/config/crowdsec-spoa-bouncer.yaml \
		${STAGEDIR}${ETCDIR}/crowdsec-spoa-bouncer.yaml.sample

	@${MKDIR} ${STAGEDIR}${DATADIR}
	${INSTALL_SCRIPT} ${WRKSRC}/scripts/_bouncer.sh \
		${STAGEDIR}${DATADIR}/

	@${MKDIR} ${STAGEDIR}${DATADIR}/lua
	${INSTALL_DATA} ${WRKSRC}/lua/*.lua \
		${STAGEDIR}${DATADIR}/lua/

	@${MKDIR} ${STAGEDIR}${DATADIR}/html
	${INSTALL_DATA} ${WRKSRC}/templates/*.html \
		${STAGEDIR}${DATADIR}/html/

	@${MKDIR} ${STAGEDIR}${DOCSDIR}/examples
	${INSTALL_DATA} ${WRKSRC}/config/crowdsec.cfg \
		${STAGEDIR}${DOCSDIR}/examples/
	${INSTALL_DATA} ${WRKSRC}/config/haproxy*.cfg \
		${STAGEDIR}${DOCSDIR}/examples/

	@${MKDIR} ${STAGEDIR}/var/log/crowdsec-spoa

.include <bsd.port.mk>
