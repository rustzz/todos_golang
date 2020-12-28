import test_notes
import test_auth
import random
from pprint import pprint
from threading import Thread

threads = []

def start():
    username, password, response = test_auth.signup()
    response = test_auth.signin(username, password)
    token = response["token"]

    response = test_notes.get(username, token)
    pprint(response)
    for _ in range(3):
        response = test_notes.add(username, token)
        pprint(response)
    response = test_notes.get(username, token)
    pprint(response)
    _id = random.choice([int(x) for x in response["notes"].keys()])
    response = test_notes.update(username, token, {"id": _id})
    response = test_notes.get(username, token)
    pprint(response)
    _ids = random.choice([int(x) for x in response["notes"].keys()])
    test_notes.delete(username, token, {"id": _id})
    response = test_notes.get(username, token)
    pprint(response)
    return

# start()

for _ in range(1000):
    x = Thread(target=start)
    threads.append(x)
    x.start()

print("END")