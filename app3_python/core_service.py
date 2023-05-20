import time

from flask import Flask, request
import requests
import random


app = Flask(__name__)


def authenticate_token_on_database(token):
    print('core - authenticating token')

    time.sleep(random.uniform(0.5, 2))  # Simulate database call.

    if token == '123':
        return True
    return False


def call_bootstrap_service(user):
    print(f"core - calling bootstrapsvc endpoint with {user}")
    response = requests.post(
        "http://localhost:8081/bootstrap",
        json={'user': user}
    )


def update_database():
    time.sleep(random.uniform(0.5, 3))  # Simulate database call.


@app.route("/bootstrap-user", methods=['POST', 'GET'])
def bootstrap_user():
    print("core - Starting bootstrap operation.")
    token = request.args.get('token')
    user = request.args.get('user')
    if not authenticate_token_on_database(token):
        return "core - Wrong Token!", 401

    try:
        call_bootstrap_service(user)
    except Exception as e:
        print(e)
        return "core - Error sending req to bootstrapservice", 503

    # simulate a call to update database.
    update_database()

    return "core - user has been bootstrapped", 200


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
