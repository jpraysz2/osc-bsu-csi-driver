#!/bin/sh

IMAGE_NAME=doclab.azurecr.io/osc-bsu-csi-driver
IMAGE_TAG=1.0.0
REGION=cloudgouv-eu-west-1

helm uninstall osc-bsu-csi-driver  --namespace kube-system

helm install osc-bsu-csi-driver ./osc-bsu-csi-driver \
     --namespace kube-system --set enableVolumeScheduling=true \
     --set enableVolumeResizing=true --set enableVolumeSnapshot=true \
     --set region=$REGION \
    --set image.repository=$IMAGE_NAME \
    --set image.tag=$IMAGE_TAG
