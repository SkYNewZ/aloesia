import os
import logging
import json
import base64
import sys

# Check needed environment variables
MANDATORY_ENV_VARS = ["SMTP_HOST", "SMTP_PORT", "SMTP_USER", "SMTP_PASSWORD"]
for var in MANDATORY_ENV_VARS:
    if var not in os.environ:
        raise EnvironmentError(f"Failed because '{var}' is not set.")

SMTP_HOST = os.getenv("SMTP_HOST")
SMTP_PORT = os.getenv("SMTP_PORT")
SMTP_USER = os.getenv("SMTP_USER")
SMTP_PASSWORD = os.getenv("SMTP_PASSWORD")
MAIL_FROM = os.getenv("MAIL_FROM", default="Aloesïa Management <management@aloesia.fr>")


# Init the SMTP client
import smtplib
from email.message import EmailMessage

SMTP_CLIENT = smtplib.SMTP(host=SMTP_HOST, port=SMTP_PORT)
SMTP_CLIENT.login(user=SMTP_USER, password=SMTP_PASSWORD)

# Check needed pub/sub content
MANDATORY_PUBSUB_CONTENT = ["subject", "destination", "content"]


def send_mail(event, context):
    """Background Cloud Function to be triggered by Pub/Sub.
    Args:
         event (dict):  The dictionary with data specific to this type of
         event. The `data` field contains the PubsubMessage message. The
         `attributes` field will contain custom attributes if there are any.
         context (google.cloud.functions.Context): The Cloud Functions event
         metadata. The `event_id` field contains the Pub/Sub message ID. The
         `timestamp` field contains the publish time.
    """
    if "data" not in event:
        raise RuntimeError("Unable to get 'data' in pub/sub message. Aborting…")

    message = base64.b64decode(event["data"]).decode("utf-8")
    try:
        message = json.loads(message)
    except Exception as error:
        logging.error("Unable to decode input JSON message. Aborting…")
        sys.exit(1)

    # Ensure we have needed parameters
    for var in MANDATORY_PUBSUB_CONTENT:
        if var not in message:
            raise RuntimeError(
                f"Failed because '{var}' is not set into pub/sub message."
            )

    mail = EmailMessage()
    mail.set_content(message["content"])
    mail["Subject"] = message["subject"]
    mail["From"] = MAIL_FROM
    mail["To"] = message["destination"]

    logging.info(
        "Sending message to {} with Subject '{}'".format(
            message["destination"], message["subject"]
        )
    )

    SMTP_CLIENT.send_message(msg=mail)


# Just for dev purpose
if __name__ == "__main__":
    debug_pubsub_message = {
        "subject": "Hello! This is the subject from my Cloud Function",
        "destination": "quentin@lemairepro.fr",
        "content": "Hello Quentin! This the mail content!",
    }
    # encoded_dict = json.dumps(debug_pubsub_message)
    encoded_dict = "coucou"
    encoded_dict = encoded_dict.encode("utf-8")
    base64_dict = base64.b64encode(encoded_dict)
    event = {}
    event["data"] = base64_dict
    send_mail(event, None)
