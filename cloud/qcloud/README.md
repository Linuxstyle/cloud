echo "# cloud" >> README.md
git init
git add README.md
git commit -m "first commit"
git remote add cloud git@github.com:Linuxstyle/cloud.git
git push -u cloud  master
