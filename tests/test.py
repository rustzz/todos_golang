import test_notes
import test_auth
import random
from pprint import pprint
from threading import Thread

threads = []

def start():
    username, password, response = test_auth.signup()
    response = test_auth.signin(username, password)
    try:
        token = response["token"]
    except KeyError:
        pprint(response)
        return

    response = test_notes.get(username, token)
    pprint(response)
    for _ in range(3):
        response = test_notes.add(username, token)
        pprint(response)
    response = test_notes.get(username, token)
    pprint(response)
    try:
        _id = random.choice([int(x) for x in response["notes"].keys()])
    except KeyError:
        return
    response = test_notes.update(username, token, {"id": _id})
    response = test_notes.get(username, token)
    pprint(response)
    try:
        _ids = random.choice([int(x) for x in response["notes"].keys()])
    except KeyError:
        return
    test_notes.delete(username, token, {"id": _id})
    response = test_notes.get(username, token)
    pprint(response)
    return

for _ in range(int(input("[COUNT OF THREADS]> "))):
    x = Thread(target=start)
    threads.append(x)
    x.start()
