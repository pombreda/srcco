#!/usr/bin/env bash

# Copyright (c) 2011 Ian MacLeod <ian@nevir.net>
#
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
# associated documentation files (the "Software"), to deal in the Software without restriction,
# including without limitation the rights to use, copy, modify, merge, publish, distribute,
# sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies or
# substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
# NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
# DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT
# OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

# Lovingly taken from github.com/nevir/groc.

set -e # Stop on the first failure that occurs

DOCS_PATH=.git/srcco-tmp
TARGET_BRANCH=gh-pages
[[ $1 ]] && TARGET_REMOTE=$1 || TARGET_REMOTE=origin

# Git spits out status information on $stderr, and we don't want to relay that as an error to the
# user.  So we wrap git and do error handling ourselves...
exec_git() {
  args=''
  for (( i = 1; i <= $#; i++ )); do
    eval arg=\$$i
    if [[ $arg == *\ * ]]; then
      # } We assume that double quotes will not be used as part of argument values.
      args="$args \"$arg\""
    else
      args="$args $arg"
    fi
  done

  set +e
  # } Even though we wrap the arguments in quotes, bash is splitting on whitespace within.  Why?
  result=`eval git $args 2>&1`
  status=$?
  set -e

  if [[ $status -ne 0 ]]; then
    echo "$result" >&2
    exit $status
  fi

  echo "$result"
  return 0
}

if [[ `git status -s` != "" ]]; then
  echo "Please commit or stash your changes before publishing documentation to github!" >&2
  exit 1
fi

CURRENT_BRANCH=`git branch 2>/dev/null| sed -n '/^\*/s/^\* //p'`
CURRENT_COMMIT=`git rev-parse HEAD`

[[ $2 ]] && COMMIT_MESSAGE=$2 || COMMIT_MESSAGE="Generated documentation for $CURRENT_COMMIT"

# Preserve the project's .gitignore so that we don't check in or otherwise screw up hidden files
if [[ -e .gitignore ]]; then
  cp .gitignore $DOCS_PATH/
fi

if [[ `git branch --no-color | grep " $TARGET_BRANCH"` == "" ]]; then
  # Do a fetch from the target remote to see if it was created remotely
  exec_git fetch $TARGET_REMOTE

  # Does it exist remotely?
  if [[ `git branch -a --no-color | grep " remotes/$TARGET_REMOTE/$TARGET_BRANCH"` == "" ]]; then
    echo "No '$TARGET_BRANCH' branch exists.  Creating one"
    exec_git symbolic-ref HEAD refs/heads/$TARGET_BRANCH
    rm .git/index

    # Preserve ignored files, but make sure they're actually ignored!
    if [[ -e $DOCS_PATH/.gitignore ]]; then
      cp $DOCS_PATH/.gitignore .gitignore
      exec_git add .gitignore
    fi

    exec_git clean -fdq
  else
    TARGET_REMOTE=origin
    echo "No local branch '$TARGET_BRANCH', checking out '$TARGET_REMOTE/$TARGET_BRANCH' and tracking that"
    exec_git checkout -b $TARGET_BRANCH $TARGET_REMOTE/$TARGET_BRANCH
  fi

else
  exec_git checkout $TARGET_BRANCH
fi

# We want to keep in complete sync (deleting old docs, or cruft from previous documentation output)
exec_git ls-files | xargs rm -f

# The previous solution below fails to copy .dot-files, therefore we utilize
# `find`, which in turn also made the `if`-statement obsolete.
#
# >     cp -Rf $DOCS_PATH/* .
# >     if [[ -e $DOCS_PATH/.gitignore ]]; then
# >       cp $DOCS_PATH/.gitignore .gitignore
# >     fi
#   
# Alternative solution using `tar`
#
# >     ( cd $DOCS_PATH ; tar --excude .git cf - . ) | tar xkpvvf -
#
find $DOCS_PATH -maxdepth 1 -not -path $DOCS_PATH -and -not -path $DOCS_PATH/.git -exec cp -Rf "{}" . \;

# Do nothing unless we actually have changes
if [[ `git status -s` != "" ]]; then
  exec_git add -A
  exec_git commit -m "$COMMIT_MESSAGE"
  exec_git push $TARGET_REMOTE $TARGET_BRANCH
fi

# Clean up after ourselves
rm -rf $DOCS_PATH

exec_git checkout "$CURRENT_BRANCH"
