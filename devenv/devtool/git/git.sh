git filter-branch -f --env-filter "
if [ \"\$GIT_COMMITTER_NAME\" = \"$OLD_AUTHOR_NAME\" ]
then
    export GIT_COMMITTER_NAME=\"贾一饼\"
    export GIT_COMMITTER_EMAIL=\"lby.i@qq.com\"
fi
if [ \"\$GIT_AUTHOR_NAME\" = \"$OLD_AUTHOR_NAME\" ]
then
    export GIT_AUTHOR_NAME=\"贾一饼\"
    export GIT_AUTHOR_EMAIL=\"lby.i@qq.com\"
fi
" --tag-name-filter cat -- --branches --tags