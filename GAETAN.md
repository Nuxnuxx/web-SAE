# DONE

choix des tech
Structure du projet
Git setup repo
Sonarqube Check GH (github actions)
Snyk Check GH (github actions)
ecriture du backend en node puis en go
ecriture des fonctionnalite frontend
    - authentification
    - recherche
    - pagination
Pourquoi :
    Permet d'utiliser le module de concurrence built in go et une codebase puis succinte (node devenait boilerplate)
Dockerization du projet
    - Creation docker compose et dockerfile pour chaque service
Repository prive avec ghrc (github packages)
Build image automatic sur packages github GH (github actions)
Securite server UPJV
    - changement du port ssh par default
    - ajout du non root user au group docker
    - seulement afficher le port du frontend
    - credential store avec docker et pass
    - desactivation de la connexion avec mot de passe (cle ssh seulement)
docker swarm et stack deployement sur serveur UPJV

# TODO
- [ ] ssh run deploy CI CD if build as succeed
