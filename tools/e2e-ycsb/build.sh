#!/usr/bin/env bash
# This script builds the image in the host docker registry with
# the latest tag
# if a registry is specified then it pushes the image to that
# registry with the configure TAG value see below
# if --also-tag-as-latest option is passed then the image is
# also pushed to the registry with the tagged as latest
# This works for legacy test runs and also other test frameworks
# as long as we do not make breaking changes.

set -e
IMAGE="openebs/e2e-ycsb"
TAG="v1.0.0"
registry=""
tag_as_latest=""

while [ "$#" -gt 0 ] ; do
    case "$1" in
        --registry)
            shift
            registry=$1
            ;;
        --also-tag-as-latest)
            tag_as_latest="Y"
            ;;
        *)
            echo "Unknown option: $1"
            help
            exit 1
            ;;
    esac
    shift
done

if docker build -t ${IMAGE} . ; then
    if [ "${registry}" != "" ]; then
        echo "image registry: ${registry}"
        YCSB_IMAGE="${registry}/${IMAGE}"
    else
        echo "image registry: dockerhub"
        YCSB_IMAGE="${IMAGE}"
    fi
    if [ "${tag_as_latest}" == "Y" ];  then
        echo "tagging as latest and push image ${YCSB_IMAGE}"
        docker tag ${IMAGE} ${YCSB_IMAGE}
        docker push ${YCSB_IMAGE}
    fi
    if [ "${TAG}" != "" ];  then
        echo "tagging as ${TAG} and push image  ${YCSB_IMAGE}"
        docker tag ${IMAGE} ${YCSB_IMAGE}:${TAG}
        docker push ${YCSB_IMAGE}:${TAG}
    else
        echo "TAG was not defined - image not retagged and pushed"
    fi
else
    exit 1
fi

