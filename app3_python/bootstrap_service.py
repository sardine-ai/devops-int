from flask import Flask
from flask import request
import requests
import time
import random

app = Flask(__name__)


def get_user_information_from_external_vendor_api():
    print("bootstrapsvc - extvendorapi")
    requests.get('http://www.google.com')


def get_user_information_from_external_vendor_api_404_error():
    print("bootstrapsvc - 404")
    requests.get('https://httpstat.us/404')


def do_some_redis_operations():
    print("bootstrapsvc - redis op")
    time.sleep(random.uniform(0.01, 0.05))


@app.route("/bootstrap", methods=['POST', 'GET'])
def boot_strap_user():
    print(f'Received data with {request.data}')

    get_user_information_from_external_vendor_api()
    get_user_information_from_external_vendor_api_404_error()
    do_some_redis_operations()

    return "bootstrapsvc - bootstrapped", 200


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8081)
