# windows
create Heroku account https://signup.heroku.com/dc
installer golang https://golang.org/doc/install && heroku CLI https://cli-assets.heroku.com/heroku-x64.exe
Création d'un compte Github

# sur votre IDE
crée un folder du nom de votre projet
dans la console (powershell) integrer a VSCODE
go mod init (le nom de votre fichier projet)
git init

# création de votre vos fichier
clique droit dans l'explorateur et crée vaut fichier comme vous le voulez

# sauvegarde et mise en place sur Heroku
    # Sur GitHub
    Création d'un Repositerie du nom de votre projet

heroku login

    # Sauvegarde sur git
        # Pour la premiere fois
        git remote add origin https://github.com/votrespeudogithub/nomdurepositerie
        git config user.email "votre@mail utilisé pour GitHub"
        git config user.name "Pseudo GitHub"

    git add * (pour add tout votre projet)
    git commit -m "nom du commit"
    git push origin master
 
heroku create nom-de-votre-projet
git add *
git commit -m "nom d'un commit"
git push heroku master

# voir le resultat
vous rendre sur votre espace heroku
cliquer sur votre projet 
open app