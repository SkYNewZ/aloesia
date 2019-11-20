#!/bin/bash

gcloud functions deploy send-mail \
--region=europe-west2 \
--entry-point=send_mail \
--memory=128MB \
--runtime=python37 \
--env-vars-file=env.yaml \
--trigger-topic=send-mail \
--project=$GOOGLE_CLOUD_PROJECT
