#!/bin/bash
# Set RHOST



##!/bin/bash

# this script pushes local changes to the vm - this is similar
# to rbuild, but only pushes and does not build
# see the look-aside wiki page for details
#
#set -e
#
#LOCAL_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
#
#source $LOCAL_DIR/rgit_utilities.sh

set -e

#
# rgit_push:
#
# $1: IN - The text to display before issuing the ssh command
# $2: IN - The command to execute on the VM, via ssh
#
function rgit_push() {
    # change Mac directory to the root of the repo, in case we're in a sub-directory
    pushd `git rev-parse --show-toplevel`

    # get the current branch name
    BRANCHNAME=`git rev-parse --abbrev-ref HEAD`

    # see if there are any changes we need to save
    CHANGES=`git status --short | egrep "^([ A-Za-z][ A-Za-z]|\?\?)" | wc -l`

    if [ $CHANGES -gt 0 ]; then
        git add .

        # see if there's a special rbuild commit we can amend
        if [ `git log -n 1 --oneline | grep -c "dummy rbuild commit for $BRANCHNAME"` -eq 0 ]; then
            # no previous rbuild commit
            echo "Creating commit 'dummy rbuild commit for $BRANCHNAME'"
            git commit -m "dummy rbuild commit for $BRANCHNAME" --no-edit --allow-empty --no-verify
        else
            echo "Amending commit 'dummy rbuild commit'"
            git commit --amend --no-edit --allow-empty --no-verify
        fi
    fi

    # save the current/dummy SHA on the MAC, to verify it's
    # the same as the one used when building on the VM
    COMMIT_SHA=`git rev-parse HEAD`

    # what the following does on the VM, in plain English:
    # 0) Pushes the branch to the build system designated by the remote
    # 1) changes to the same directory as the Mac is in
    # 2) checkout the branch (which might be the same one the VM is already in)
    # 3) deletes any untracked files/directories
    # 4) refreshes the submodules in case they changed
    # 5) verifies the current log's top SHA matches the one we got on the Mac
    # 6) executes the second argument passed to this rgit_push() function

    # NOTE: this requires setting the VM's git config for receive.denycurrentbranch=ignore,
    # so the rgpull script does this on the VM:
    #    git config --local receive.denyCurrentBranch ignore

    echo $1
    if [[ -z $RHOST ]]; then
        RHOST=vm1
    fi

    git push -f ssh://$RHOST:`pwd` +$BRANCHNAME

    ssh -tt $RHOST "pushd `pwd` && \
                    git clean -df && \
                    git checkout $BRANCHNAME && \
                    greset --hard HEAD && \
                    gmodule update && \
                    git rev-parse HEAD | grep -q $COMMIT_SHA && \
                    $2"

    # change back to the original Mac directory we were in when we entered this script
    popd
}

