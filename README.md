# Jeu du Pendu (Hangman) en Go

Bienvenue dans le jeu du Pendu (Hangman) réalisé en Go. Ce projet implémente le célèbre jeu où vous devez deviner un mot caché en proposant des lettres une par une avant d'être "pendu"!

## Fonctionnalités

- Joueur unique
- Choix aléatoire d'un mot à deviner parmi une liste prédéfinie
- Affichage des lettres trouvées et des lettres manquantes
- Comptabilisation des essais ratés
- Gestion des victoires et des défaites
- Interface en ligne de commande

## Prérequis

- **Golang** version 1.16 ou plus récente
- Un terminal

## Installation

Clonez le dépôt dans votre environnement local :

```bash
git clone https://github.com/votre-utilisateur/hangman-go.git
cd hangman-go
cd src
```

## Utilisation

Lancez simplement le jeu dans votre terminal :

```bash
go run .\main\
```

Les règles du jeu sont les suivantes :

- Un mot mystère sera sélectionné aléatoirement parmi une liste.
- Vous devez deviner le mot en proposant une lettre à chaque tour.
- Si la lettre fait partie du mot, elle sera révélée dans les emplacements correspondants.
- Si la lettre est incorrecte, vous perdrez une tentative.
- Vous gagnez si vous devinez toutes les lettres du mot avant d'épuiser vos tentatives.
- Vous perdez si vous atteignez le nombre maximal de tentatives incorrectes (6).
