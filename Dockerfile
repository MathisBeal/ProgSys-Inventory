# Image golang pour la compilation
FROM golang:1.25 AS builder

# Copier les fichiers dont ont a besoin
COPY src /go/src
COPY Makefile /go

# Compilation
RUN make build

# Image finale
FROM busybox

# Installer le logiciel compilé
COPY src/www /www
COPY --from=builder /go/bin/inventory /inventory

# Partages
EXPOSE 80

# Démarrage
ENTRYPOINT [ "/inventory" ]