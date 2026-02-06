# ProgSys-Inventory

ProgSys-Inventory est une application Go qui expose un **tableau de bord système** et une **API HTTP** pour inventorier rapidement l’état d’une machine Linux (CPU, mémoire, disques, réseau, processus, load average).

## 1) Fonctionnement global du projet

L’application lance un serveur HTTP sur le port `80`.

- Les endpoints `/cpu`, `/mem`, `/disk`, `/net`, `/ps`, `/load` retournent des données JSON collectées avec la librairie `gopsutil`.
- L’endpoint `/health` retourne `ok` pour les checks de disponibilité.
- Les pages web statiques (dans `src/www`) sont servies à la racine (`/`) et interrogent ces endpoints pour afficher les métriques en direct.

Flux simplifié :

1. Un navigateur charge une page HTML (`/`, `/procs.html`, `/mem.html`, etc.).
2. Le JavaScript de la page appelle périodiquement les endpoints JSON.
3. Le backend Go collecte les infos système locales via `gopsutil`.
4. Les données sont rendues côté client dans des tableaux/cartes/barres.

---

## 2) Documentation Utilisateur

## Accès à l’interface

- URL d’accueil : `http://<host>/`
- Pages principales :
  - `http://<host>/procs.html` — Processus
  - `http://<host>/network.html` — Réseau
  - `http://<host>/mem.html` — Mémoire
  - `http://<host>/disk.html` — Disques
  - `http://<host>/load.html` — Charge système + CPU par cœur

## Ce que voit l’utilisateur

- **Accueil** : présentation du projet + liens rapides vers pages et APIs.
- **Processus** : liste des processus, filtres et tri (CPU, mémoire, PID, nom).
- **Réseau** : interfaces et compteurs réseau.
- **Mémoire** : état mémoire virtuelle et swap.
- **Disques** : partitions et usage.
- **Load/CPU** : load average (1/5/15 min) et utilisation par cœur.

Les pages se rafraîchissent automatiquement (généralement toutes les 5 secondes).

## Endpoints API (pour utilisateurs avancés / intégrations)

### Santé
- `GET /health`

### CPU
- `GET /cpu`
- Retour : cœurs CPU avec informations, `% d’usage` et temps CPU.

### Processus
- `GET /ps` : tous les processus
- `GET /ps/{user}` : processus d’un utilisateur
- `GET /ps/kill/{pid}` : termine le processus correspondant au PID

### Réseau
- `GET /net` : toutes les interfaces
- `GET /net/{card}` : interface spécifique (ex: `eth0`)

### Mémoire
- `GET /mem`

### Disques
- `GET /disk`

### Charge
- `GET /load`

Exemple :

```bash
curl http://localhost/health
curl http://localhost/load
curl http://localhost/ps/root
```

---

## 3) Documentation Administrateur

## Prérequis

- Go `1.25.x` (selon `go.mod`)
- Linux recommandé (les métriques reposent sur la machine hôte)
- Droits suffisants pour écouter le port `80` (root/capabilities/reverse proxy)

## Exécution locale

Depuis la racine du repo :

```bash
make run
```

Le serveur écoute sur `:80`.

## Build binaire

```bash
make build
```

Produit : `bin/inventory`

Nettoyage :

```bash
make clean
```

## Déploiement Docker

Build image :

```bash
make image
```

Démarrage container :

```bash
make start
```

Arrêt/suppression :

```bash
make stop
```

Détails image :

- Build multi-stage (`golang:1.25` → `busybox`)
- Binaire copié en `/inventory`
- Assets statiques copiés en `/www`
- Port exposé : `80`

## Déploiement systemd

Un exemple d’unité est fourni : `inventory.service`.

Procédure type :

1. Compiler et copier le binaire (ex: `/root/inventory` selon le service actuel).
2. Copier `inventory.service` vers `/etc/systemd/system/inventory.service`.
3. Recharger systemd : `systemctl daemon-reload`.
4. Activer et démarrer :
   - `systemctl enable inventory`
   - `systemctl start inventory`
5. Vérifier :
   - `systemctl status inventory`
   - `journalctl -u inventory -f`

## Observabilité / exploitation

- Vérification liveness : `GET /health`
- Les endpoints techniques peuvent être branchés à un outil de supervision interne.
- En cas d’erreurs API, le serveur retourne `500` avec le message d’erreur.

## Sécurité et exposition

- L’application expose des informations système sensibles.
- Recommandations :
  - restreindre l’accès réseau (firewall/VPN)
  - placer derrière un reverse proxy avec authentification
  - éviter l’exposition internet directe sans contrôle d’accès

## Limites connues

- Les valeurs dépendent des permissions de l’utilisateur exécutant le service.
- Certaines partitions/FS système sont volontairement filtrées dans l’endpoint disque.
- Certaines infos processus peuvent être incomplètes si les permissions sont insuffisantes.

---

## 4) Arborescence rapide

```text
.
├── Makefile
├── Dockerfile
├── inventory.service
└── src
    ├── main.go          # Démarrage serveur HTTP
    ├── routes.go        # Déclaration des routes
    ├── handle.go        # Handlers API
    ├── cpu/             # Collecte CPU
    ├── memory/          # Collecte mémoire
    ├── disk/            # Collecte disques
    ├── netcard/         # Collecte réseau
    ├── proc/            # Collecte processus
    ├── load/            # Collecte load average
    └── www/             # Front statique (HTML/CSS/JS)
```

## 5) Améliorations possibles

- Ajouter authentification/autorisation native.
- Ajouter TLS natif ou reverse proxy documenté (Nginx/Caddy).
- Ajouter pagination côté API pour `/ps`.
- Ajouter métriques Prometheus.
- Ajouter tests unitaires et d’intégration.
